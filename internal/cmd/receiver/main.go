package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/controllers"
	"github.com/matsunaga-project-2022/course/chat/internal/usecases"
)

func main() {
	pubsub := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// running worker on background
	messageReceiver := usecases.NewMessageReceiver(pubsub)

	e := echo.New()
	receiveHandler := controllers.NewReceiveHandler(messageReceiver)
	e.POST("/post", receiveHandler.Receive())

	e.Logger.Fatal(e.Start(":8080"))
}
