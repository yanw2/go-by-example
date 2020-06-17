package main

import (
	"fmt"
	"time"
)

func main() {
	// Go's `select` lets you wait on multiple channel operations.
	// For our example, we'll select across two channels. Each channel will receive a value after
	// some amount of time.
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// Timeouts are important for programs that connect to external resources or that otherwise need to
	// bound exectuion time. For example, suppose we're exectuing an external call that returns its
	// result on a channel c3 after 2s. Note that the channel is buffered so the send in the goroutine is
	// nonblocking. This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "result 3"
	}()

	// Here's the `select` implementing a timeout. Since `select` proceeds with the first receive that's
	// ready, we'll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive from c4 will success and we'll print the result
	c4 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c4 <- "result 4"
	}()

	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	// Basic sends and receives on channels are blocking. However, we can use `select` with a `default` clause
	// to implement non-blocking sends, receives, and even non-blocking multi-way selects.
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use mulitple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// Closing a channel indicates that no more values will be sent on it.
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	// We await the worker using the synchronization apporach
	<-done

	// Range over Channels
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
