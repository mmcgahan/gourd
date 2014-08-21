package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/mmcgahan/gourd/app/stream"
	"github.com/revel/revel"
	"strconv"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) StreamSocket(ws *websocket.Conn) revel.Result {
	// activate new Stream, with NewPoints channel
	watcher := stream.Watch()
	defer watcher.Unwatch()

	// In order to select between websocket messages and newPoint data, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		// listener on socket
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new data
	// on the websocket (data) and newMessages (msg, ok).
	for {
		select {
		case point := <-watcher.NewPoints:
			// broadcast NewPoint to socket as JSON, e.g.
			// {"XVal":1,"YVal":2,"Timestamp":1407452575,"Label":"C"}
			if websocket.JSON.Send(ws, &point) != nil {
				// They disconnected.
				return nil
			}
		case rate, ok := <-newMessages:
			// process received message on the socket
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}
			// Otherwise, say something.
			i, _ := strconv.Atoi(rate)
			stream.Throttle(i)
		}
	}
	return nil
}
