package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/controllers"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"github.com/matsunaga-project-2022/course/chat/internal/usecases"
	"os"
)

func main() {
	ctx := context.Background()
	clients := make(map[string]*models.Client, 0)

	pubsub := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	// running worker on background
	messagePublisher := usecases.NewMessagePublisher(pubsub)
	go messagePublisher.Publish(ctx, "chat", clients)

	e := echo.New()
	sendHandler := controllers.NewSendHandler(clients)
	e.GET("/chat", sendHandler.Send)

	e.Logger.Fatal(e.Start(":8080"))
}
