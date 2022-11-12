package models

import (
	"github.com/go-redis/redis/v9"
	"time"
)

type Message struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

type IMessageEventer interface {
	Get(id string) (Message, error)
	Save(message Message) error
}

type MessageEventer struct {
	manager *redis.Client
}

func NewMessageEventer(manager *redis.Client) *MessageEventer {
	return &MessageEventer{manager: manager}
}

func (m MessageEventer) Get(id string) (Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageEventer) Save(message Message) error {
	//TODO implement me
	panic("implement me")
}
