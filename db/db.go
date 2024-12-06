package db

import (
	"context"

	"github.com/supabase-community/supabase-go"
	"golang.org/x/xerrors"
)

type DB struct {
	client *supabase.Client
}

func New(url, key string) (DB, error) {
	client, err := supabase.NewClient(url, key, &supabase.ClientOptions{})
	if err != nil {
		return DB{}, xerrors.Errorf(": %w", err)
	}

	return DB{client: client}, nil
}

func (db *DB) ItemCreate(ctx context.Context, item ItemCreate) error {
	_, _, err := db.client.From("items").Insert(item, false, "", "", "").Execute()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
