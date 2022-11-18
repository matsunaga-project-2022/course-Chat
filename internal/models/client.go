package models

import "golang.org/x/net/websocket"

type Client struct {
	Connection *websocket.Conn
	Messages   chan Message
}

func NewClient(connection *websocket.Conn) *Client {
	messages := make(chan Message, 1)
	return &Client{Connection: connection, Messages: messages}
}
