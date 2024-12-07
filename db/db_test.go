package db_test

import (
	"context"
	"os"
	"testing"

	"github.com/nzoschke/shadowlink/db"
	"github.com/stretchr/testify/assert"
)

// func TestItemCreate(t *testing.T) {
// 	t.Skip()
// 	d, err := db.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := d.ItemCreate(context.Background(), db.ItemCreate{
// 		ChannelName: "test",
// 		Link:        "https://example.com",
// 		Meta: oembed.Info{
// 			Description: "test",
// 		},
// 		ServiceID:   "1",
// 		ServiceName: "test",
// 		UserName:    "test",
// 	}); err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestItemLifecycle(t *testing.T) {
	ctx := context.Background()

	d, err := db.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
	assert.NoError(t, err)

	key := db.ItemKey{
		ServiceID: "s1",
		URL:       "https://soundcloud.com/avalonemerson",
	}

	err = d.ItemDestroy(ctx, key)
	assert.NoError(t, err)

	in := db.ItemUpsert{
		ServiceID: "s1",
		URL:       "https://soundcloud.com/avalonemerson",
	}

	err = d.ItemUpsert(ctx, in)
	assert.NoError(t, err)

	err = d.ItemUpsert(ctx, in)
	assert.NoError(t, err)

	err = d.ItemDelete(ctx, key)
	assert.NoError(t, err)
}
