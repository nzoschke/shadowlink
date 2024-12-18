package bot

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nzoschke/shadowlink/db"
	"github.com/nzoschke/shadowlink/extract"
	"golang.org/x/xerrors"
)

func Open(token string, db db.DB) (*discordgo.Session, func() error, error) {
	ctx := context.Background()

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, nil, xerrors.Errorf(": %w", err)
	}

	dg.AddHandler(use(ctx, db, messageCreate))
	dg.AddHandler(use(ctx, db, messageReactionAdd))

	dg.Identify.Intents |= discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentsGuildMessageReactions

	if err = dg.Open(); err != nil {
		return nil, nil, xerrors.Errorf(": %w", err)
	}

	return dg, dg.Close, nil
}

func use[T any](ctx context.Context, d db.DB, f func(context.Context, db.DB, *discordgo.Session, T) error) func(*discordgo.Session, T) {
	return func(s *discordgo.Session, m T) {
		if err := f(ctx, d, s, m); err != nil {
			log.Printf("ERROR: %+v", err)
		}
	}
}

func messageReactionAdd(ctx context.Context, d db.DB, s *discordgo.Session, r *discordgo.MessageReactionAdd) error {
	if r.Emoji.Name != "❌" {
		return nil
	}

	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	infos, err := extract.MediaInfos(m.Content)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	for _, info := range infos {
		log.Printf("Deleting %s\n", info.URL)

		if err := d.ItemDelete(ctx, db.ItemKey{
			ServiceID: r.GuildID,
			URL:       info.URL,
		}); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	if len(infos) > 0 {
		if err := s.MessageReactionRemove(m.ChannelID, m.ID, "🔗", s.State.User.ID); err != nil {
			return xerrors.Errorf(": %w", err)
		}

		if err := s.MessageReactionAdd(m.ChannelID, m.ID, "⛓️‍💥"); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	return nil
}

func messageCreate(ctx context.Context, d db.DB, s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.ID == s.State.User.ID {
		return nil
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	infos, err := extract.MediaInfos(m.Content)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	for _, info := range infos {
		log.Printf("Inserting %s\n", info.URL)

		if err := d.ItemUpsert(ctx, db.ItemUpsert{
			Author: db.Author{
				Channel: channel.Name,
				Name:    m.Author.Username,
				Service: guild.Name,
			},
			Meta:      info,
			ServiceID: m.GuildID,
			URL:       info.URL,
		}); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	if len(infos) > 0 {
		if err := s.MessageReactionAdd(m.ChannelID, m.ID, "🔗"); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	return nil
}
