package db

import (
	"time"

	"github.com/dyatlov/go-oembed/oembed"
)

type Author struct {
	Channel string `json:"channel"`
	Name    string `json:"name"`
	Service string `json:"service"`
}

type Item struct {
	Author    Author      `json:"author"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt time.Time   `json:"deleted_at"`
	ID        int         `json:"id"`
	Meta      oembed.Info `json:"meta"`
	ServiceID string      `json:"service_id"`
	URL       string      `json:"url"`
}

type ItemDelete struct {
	DeletedAt time.Time `json:"deleted_at"`
}

type ItemKey struct {
	ServiceID string `json:"service_id"`
	URL       string `json:"url"`
}

type ItemUpsert struct {
	Author    Author      `json:"author"`
	Meta      oembed.Info `json:"meta"`
	ServiceID string      `json:"service_id"`
	UpdatedAt time.Time   `json:"updated_at"`
	URL       string      `json:"url"`
}
