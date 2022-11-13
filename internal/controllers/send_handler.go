package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/models"
	"golang.org/x/net/websocket"
)

type SendHandler struct {
	messages chan models.Message
	errors   chan error
}

func NewSendHandler(messages chan models.Message, errors chan error) *SendHandler {
	return &SendHandler{messages: messages, errors: errors}
}

func (sh *SendHandler) Send(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// 初回のメッセージを送信
		err := websocket.Message.Send(ws, "Server: Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			// Client からのメッセージを読み込む
			message := <-sh.messages
			payload, err := json.Marshal(message)
			if err != nil {
				c.Logger().Error(err)
			}
			// Client からのメッセージを元に返すメッセージを作成し送信する
			err = websocket.Message.Send(ws, payload)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
