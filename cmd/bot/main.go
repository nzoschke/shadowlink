package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nzoschke/shadowlink/bot"
	"github.com/nzoschke/shadowlink/db"
	"golang.org/x/xerrors"
)

func main() {
	if err := mainErr(); err != nil {
		log.Fatalf("ERR: %+v\n", err)
	}
}

func mainErr() error {
	db, err := db.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	_, close, err := bot.Open(os.Getenv("TOKEN"), db)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := close(); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
