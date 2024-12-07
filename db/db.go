package db

import (
	"context"
	"time"

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

func (db *DB) ItemDelete(ctx context.Context, in ItemKey) error {
	_, _, err := db.client.From("items").Update(ItemDelete{
		DeletedAt: time.Now(),
	}, "", "").Eq("service_id", in.ServiceID).Eq("url", in.URL).Execute()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (db *DB) ItemDestroy(ctx context.Context, in ItemKey) error {
	_, _, err := db.client.From("items").Delete("", "").Eq("service_id", in.ServiceID).Eq("url", in.URL).Execute()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (db *DB) ItemUpsert(ctx context.Context, in ItemUpsert) error {
	in.UpdatedAt = time.Now()
	_, _, err := db.client.From("items").Insert(in, true, "service_id,url", "", "").Execute()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
