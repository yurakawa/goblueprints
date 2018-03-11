package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている一人のユーザを表す
type client struct {
	socket *websocket.Conn // socketはこのクライアントのためのWebSocket
	send chan []byte // sendはメッセージが送られるチャネル
	room *room // roomはこのクライアントが参加しているチャットルームです
}

// WebSocketへの読み書きを行うメソッド

func (c * client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err :=c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
