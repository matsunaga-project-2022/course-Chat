package models

// Message is a data model of chat message
type Message struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Text        string `json:"text"`
	AvatarURL   string `json:"avatar_url"`
	MessageType string `json:"message_type"`
}
