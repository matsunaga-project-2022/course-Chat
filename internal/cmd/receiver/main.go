package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matsunaga-project-2022/course/chat/internal/controllers"
	"github.com/matsunaga-project-2022/course/chat/internal/usecases"
	"net/http"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	receiveHandler := controllers.NewReceiveHandler(messageReceiver)
	e.POST("/post", receiveHandler.Receive())

	e.Logger.Fatal(e.Start(":8080"))
}
