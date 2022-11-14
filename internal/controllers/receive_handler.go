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
		userID := c.Get("user_id")
		text := c.FormValue("text")

		err := rh.receiver.Receive(c.Request().Context(), "", userID.(string), text)

		if err != nil {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	}
}
