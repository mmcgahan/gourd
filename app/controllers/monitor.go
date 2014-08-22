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

func listen(ws *websocket.Conn, newMessages chan string) {
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
}

func (c WebSocket) StreamSocket(ws *websocket.Conn) revel.Result {
	// websocket now open, let up general listener
	newMessages := make(chan string)
	go listen(ws, newMessages)

	// 1. publish what data sources are available
	// websocket.JSON.Send(ws, &arrayOfSources)
	//
	// 2. listen for which sources should be published
	// for { select { case sources, ok := <-newMessages } }
	//
	// 3. activate new Stream, with NewPoints channel for this socket
	// watcher := stream.Watch(sources)
	watcher := stream.Watch()
	defer watcher.Unwatch()

	// Now listen for new data:
	// 	a. new points from stream (watcher.NewPoints)
	// 	b. throttling values from client (newMessages)
	for {
		select {
		case point := <-watcher.NewPoints:
			// broadcast NewPoint to socket as JSON, e.g.
			// {"XVal":1,"YVal":2,"Timestamp":1407452575,"Label":"C"}
			if websocket.JSON.Send(ws, &point) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// throttle
			if !ok {
				return nil
			}
			i, _ := strconv.Atoi(msg)
			stream.Throttle(i)
		}
	}
	return nil
}
