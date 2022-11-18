package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"golang.org/x/net/websocket"
	"log"
)

type SendHandler struct {
	clients map[string]*models.Client
}

func NewSendHandler(clients map[string]*models.Client) *SendHandler {
	return &SendHandler{clients: clients}
}

func (sh *SendHandler) Send(c echo.Context) error {
	// TODO: get user_id fron echo.Context
	userID := "test_user"

	s := websocket.Server{
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			sh.clients[userID] = models.NewClient(ws)
			for {
				message := <-sh.clients[userID].Messages
				err := websocket.JSON.Send(sh.clients[userID].Connection, message)
				if err != nil {
					log.Println(err)
				}
			}
		}),
	}

	s.ServeHTTP(c.Response(), c.Request())
	return nil
}
