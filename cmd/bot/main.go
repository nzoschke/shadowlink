package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nzoschke/shadowlink/bot"
	"github.com/nzoschke/shadowlink/db"
	"github.com/nzoschke/shadowlink/web"
	"golang.org/x/xerrors"
)

func main() {
	if err := mainErr(); err != nil {
		log.Fatalf("ERR: %+v\n", err)
	}
}

func mainErr() error {
	ctx := context.Background()

	db, err := db.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	_, close, err := bot.Open(os.Getenv("DISCORD_TOKEN"), db)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	shutdown := web.Start()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := close(); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := shutdown(ctx); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
