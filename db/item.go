package db

import (
	"time"

	"github.com/dyatlov/go-oembed/oembed"
)

type ItemCreate struct {
	ChannelID   string      `json:"channel_id"`
	ChannelName string      `json:"channel_name"`
	Link        string      `json:"link"`
	Meta        oembed.Info `json:"meta"`
	ServiceName string      `json:"service_name"`
	ServiceID   string      `json:"service_id"`
	UserName    string      `json:"user_name"`
	UserID      string      `json:"user_id"`
}

type Item struct {
	ID          int         `json:"id"`
	ChannelID   string      `json:"channel_id"`
	ChannelName string      `json:"channel_name"`
	CreatedAt   time.Time   `json:"created_at"`
	DeletedAt   time.Time   `json:"deleted_at"`
	Link        string      `json:"link"`
	Meta        oembed.Info `json:"meta"`
	ServiceName string      `json:"service_name"`
	ServiceID   string      `json:"service_id"`
	UserName    string      `json:"user_name"`
	UserID      string      `json:"user_id"`
}
