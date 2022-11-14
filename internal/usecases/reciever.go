package usecases

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"time"
)

type IMessageReceiver interface {
	Receive(ctx context.Context, channel string, userID string, text string) error
}

type MessageReceiver struct {
	redisClient *redis.Client
}

func NewMessageReceiver(redisClient *redis.Client) *MessageReceiver {
	return &MessageReceiver{redisClient: redisClient}
}

func (r *MessageReceiver) Receive(ctx context.Context, channel string, userID string, text string) error {
	id := uuid.New()
	structMessage := &models.Message{
		ID:        id.String(),
		UserID:    userID,
		Text:      text,
		CreatedAt: time.Now(),
	}

	byteMessage, err := json.Marshal(structMessage)
	if err != nil {
		return err
	}
	message := string(byteMessage)

	_, err = r.redisClient.Publish(ctx, "chat", message).Result()
	if err != nil {
		return err
	}

	return nil
}
