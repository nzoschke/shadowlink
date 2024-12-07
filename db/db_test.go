package db_test

import (
	"context"
	"os"
	"testing"

	"github.com/nzoschke/shadowlink/db"
	"github.com/stretchr/testify/assert"
)

func TestItemLifecycle(t *testing.T) {
	t.Skip()

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
