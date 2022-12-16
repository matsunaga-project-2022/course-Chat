package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/matsunaga-project-2022/course/chat/internal/usecases"
	"log"
	"net/http"
)

type ReceiveHandler struct {
	receiver usecases.IMessageReceiver
}

func NewReceiveHandler(receiver usecases.IMessageReceiver) *ReceiveHandler {
	return &ReceiveHandler{receiver: receiver}
}

func (rh *ReceiveHandler) Receive() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.FormValue("user_id")
		text := c.FormValue("text")
		avatarURL := c.FormValue("avatar_url")
		messageType := c.FormValue("message_type")
		if userID == "" || text == "" || avatarURL == "" || messageType == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		err := rh.receiver.Receive(c.Request().Context(), "", userID, text, avatarURL, messageType)

		if err != nil {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	}
}
