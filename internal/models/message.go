package models

import "time"

type Message struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}
