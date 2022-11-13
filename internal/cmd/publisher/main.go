package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/controllers"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"github.com/matsunaga-project-2022/course/chat/internal/usecases"
)

func main() {
	ctx := context.Background()
	messages := make(chan models.Message, 1)
	errors := make(chan error, 1)

	pubsub := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// running worker on background
	messagePublisher := usecases.NewMessagePublisher(pubsub)
	go messagePublisher.Publish(ctx, "chat", messages, errors)

	e := echo.New()
	sendHandler := controllers.NewSendHandler(messages, errors)
	e.GET("/chat", sendHandler.Send)

	e.Logger.Fatal(e.Start(":8000"))
}
