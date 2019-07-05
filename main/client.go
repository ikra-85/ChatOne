// client
// First edition
// 1.1
package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn //web socket for the client
	send   chan []byte     //a channel on which messages are sent
	room   *room           //the room this client is chatting in
}
