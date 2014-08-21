package stream

import (
	"container/list"
	"math/rand"
	"time"
)

type Point struct {
	XVal  int
	YVal  int
	Label string
}

type Stream struct {
	NewPoints chan Point // New points coming in
}

func (s Stream) Unwatch() {
	// Owner of a Stream must Unwatch it when they stop listening to events.
	// This gets called automatically with .defer after Watch()
	unwatch <- s       // unwatch the channel - send the channel to the unwatch channel
	drain(s.NewPoints) // Drain it, just in case there was a pending publish.
}

func Watch() Stream {
	// resp := make(chan Stream) // make new channel of Streams
	// watch <- resp             // put the new Stream into the watch channel
	// return <-resp             // return the sending-only channel
	stream := Stream{make(chan Point)}
	watch <- stream
	return stream
}

func Publish(point *Point) {
	publish <- *point
}

func Throttle(rate int) {
	throttle <- rate
}

var (
	// Send a stream here to connect to data feed.
	watch = make(chan Stream) // receiving channel

	// Send a stream here to to disconnect from data feed.
	unwatch = make(chan Stream)

	// Send points here to publish them to all streams.
	publish = make(chan Point)

	// send rate integers here to throttle rate of new points
	throttle = make(chan int)
	rate     = 1
)

// This function loops forever, handling the data stream pubsub
func stream() {
	// private list of Streams, one per websocket
	watchers := list.New()

	for {
		select {
		// listen to all 'global' channels
		case stream := <-watch:
			// put stream into list of watcher streams
			watchers.PushBack(stream)

		case stream := <-unwatch:
			// watcher disconnected - find the channel and remove it
			for watcher := watchers.Front(); watcher != nil; watcher = watcher.Next() {
				// watcher.Value.(Stream) gets the underlying element, with type Stream
				if watcher.Value.(Stream) == stream {
					watchers.Remove(watcher)
					break
				}
			}

		case point := <-publish:
			// new point, send it to the NewPoints channel of all watcher streams
			for watcher := watchers.Front(); watcher != nil; watcher = watcher.Next() {
				watcher.Value.(Stream).NewPoints <- point
			}

		case newRate := <-throttle:
			rate = newRate
		}
	}
}

// go routine for generating random points
func generate() {
	rand.Seed(42)
	labels := []string{"A", "B", "C"} // all possible Point Labels
	for {
		// idx := rand.Intn(len(labels))
		Publish(&Point{rand.Intn(10), rand.Intn(10), labels[0]})
		// Publish(&Point{rand.Intn(10), rand.Intn(10), labels[1]})
		// Publish(&Point{rand.Intn(10), rand.Intn(10), labels[2]})
		time.Sleep(time.Second * time.Duration(rate))
	}
}

func init() {
	go stream()
	go generate()
}

// Helpers

// Drains a given channel of any points.
func drain(ch <-chan Point) { // <-chan can only release data
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
