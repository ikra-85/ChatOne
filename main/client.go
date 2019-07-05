// client
// Second edition
// 1.2: read and write methods
package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn //web socket for the client
	send   chan []byte     //a channel on which messages are sent
	room   *room           //the room this client is chatting in
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
