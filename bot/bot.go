package bot

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nzoschke/shadowlink/db"
	"github.com/nzoschke/shadowlink/extract"
	"golang.org/x/xerrors"
)

type Bot struct {
	db db.DB
}

func Open(token string, db db.DB) (*discordgo.Session, func() error, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, nil, xerrors.Errorf(": %w", err)
	}

	b := &Bot{db: db}
	dg.AddHandler(b.messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err = dg.Open(); err != nil {
		return nil, nil, xerrors.Errorf(": %w", err)
	}

	return dg, dg.Close, nil
}

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	n, err := b.extractReact(s, m)
	if err != nil {
		log.Printf("ERROR: %+v", err)
	}

	if n > 0 {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🔗")
	}
}

func (b *Bot) extractReact(s *discordgo.Session, m *discordgo.MessageCreate) (int, error) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return 0, xerrors.Errorf(": %w", err)
	}

	n := 0
	links := extract.Extract(m.Content)
	for _, link := range links {
		info, err := extract.Info(link)
		if err != nil {
			return 0, xerrors.Errorf(": %w", err)
		}

		if err := b.db.ItemCreate(context.Background(), db.ItemCreate{
			ChannelID:   m.ChannelID,
			ChannelName: channel.Name,
			Link:        link,
			Meta:        info,
			ServiceID:   m.GuildID,
			ServiceName: guild.Name,
			UserID:      m.Author.ID,
			UserName:    m.Author.Username,
		}); err != nil {
			return 0, xerrors.Errorf(": %w", err)
		}

		n++
	}

	return n, nil
}
