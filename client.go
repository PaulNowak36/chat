package main

import (
	"github.com/gorilla/websocket"
)

// client represents a signle chatting user.
type client struct {
	//socket is the web socket for this client.
	socket *websocket.Conn
	//send is a channel on which messages are sent.
	send chan []byte // room is the room this client is chatting in.
	room *room
}
