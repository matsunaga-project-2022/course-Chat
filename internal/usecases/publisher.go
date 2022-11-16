package usecases

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"log"
	"sync"
)

// IMessagePublisher is the interface of manage publish
type IMessagePublisher interface {
	Publish(ctx context.Context, subID string, messages chan<- models.Message, errors chan<- error)
}

// MessagePublisher is the implementation of IMessagePublisher by using Redis Client
type MessagePublisher struct {
	manager *redis.Client
	mutex   sync.Mutex
}

// NewMessagePublisher is the constructor that generate new MessagePublisher instance
func NewMessagePublisher(manager *redis.Client) *MessagePublisher {
	return &MessagePublisher{manager: manager, mutex: sync.Mutex{}}
}

// Publish is the method of MessagePublisher that subscribe message
func (m *MessagePublisher) Publish(ctx context.Context, subID string, messages chan<- models.Message, errors chan<- error) {
	subscribe := m.manager.Subscribe(ctx, subID)

	ch := subscribe.Channel()
	for subscription := range ch {
		var message models.Message
		if err := json.Unmarshal([]byte(subscription.Payload), &message); err != nil {
			errors <- err
		}

		log.Println(message)
		m.mutex.Lock()
		messages <- message
		m.mutex.Unlock()
	}
}
