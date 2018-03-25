package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// clientはチャットを行っている一人のユーザを表す
type client struct {
	socket   *websocket.Conn // socketはこのクライアントのためのWebSocket
	send     chan *message   // sendはメッセージが送られるチャネル
	room     *room           // roomはこのクライアントが参加しているチャットルームです
	userData map[string]interface{}
}

// WebSocketへの読み書きを行うメソッド

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
