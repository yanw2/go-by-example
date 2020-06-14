package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The done channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// This `ping` function only accepts a channel for sending values. It would be a compile-time error
// to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// This `pong` function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels
	// from one goroutine and receive those values into another goroutine.

	// Create a new channel with `make(chan val-type)`.Channels are typed by the values they convey
	messages1 := make(chan string)

	// Send a value into a channel using the `channel <-` syntax. Here we send "ping" to the messages
	// channel we made abvoe, from a new goroutine
	go func() { messages1 <- "ping" }()

	// The `<-channel` syntax receives a value from the channel. Here we'll receive the "ping" message
	// we sent above and print it out
	msg1 := <-messages1
	fmt.Println(msg1)

	// By default sends and receives block until both the sender and receiver are ready. This property
	// allowed us to wait at the end of our program for the "ping" message without having to use any
	// other synchronization.

	// By default channels are unbuffered, meaning that they will only accept sends (`chan <-`) if there
	// is a corresponding receive (`<- chan`) ready to receive the sent value. Buffered channels accept a
	// limited number of values without a corresponding receivers for those values

	// Create a channel of strings buffering up to 2 values
	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	// We can use channels to synchronize execution across goroutine. Here's an example of using a block
	// receive to wait for a goroutine to finish.

	// Start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	// If you removed the `<-done` line from this program, the program would exit before the `worker`
	// even started
	<-done

	// When using channels as function parameters you can specify if a channel is meant to only send
	// or receive values. This specificity increases the type-safety of the program.
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
