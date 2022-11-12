package models

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

// Message is a data model of chat message
type Message struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

// IMessageEventer is the interface of manage publish and subscribe
type IMessageEventer interface {
	Subscribe(ctx context.Context, subID string) (*Message, error)
	Publish(ctx context.Context, message Message) error
}

// MessageEventer is the implementation of IMessageEventer by using Redis Client
type MessageEventer struct {
	manager *redis.Client
}

// NewMessageEventer is the constructor that generate new MessageEventer instance
func NewMessageEventer(manager *redis.Client) *MessageEventer {
	return &MessageEventer{manager: manager}
}

// Subscribe is the method of MessageEventer that subscribe message
func (m MessageEventer) Subscribe(ctx context.Context, subID string) (Message, error) {
	//TODO implement me
	panic("implement me")
}

// Publish is the method of MessageEventer that publish message
func (m MessageEventer) Publish(message Message) error {
	//TODO implement me
	panic("implement me")
}
