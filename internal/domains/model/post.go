package model

import "time"

// Post ユーザの投稿
type Post struct {
	ID        string
	Author    string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
