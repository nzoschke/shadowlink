package db_test

import (
	"context"
	"os"
	"testing"

	"github.com/dyatlov/go-oembed/oembed"
	"github.com/nzoschke/shadowlink/db"
)

func TestItemCreate(t *testing.T) {
	t.Skip()
	d, err := db.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
	if err != nil {
		t.Fatal(err)
	}

	if err := d.ItemCreate(context.Background(), db.ItemCreate{
		ChannelID:   "1",
		ChannelName: "test",
		Link:        "https://example.com",
		Meta: oembed.Info{
			Description: "test",
		},
		ServiceID:   "1",
		ServiceName: "test",
		UserID:      "1",
		UserName:    "test",
	}); err != nil {
		t.Fatal(err)
	}
}
