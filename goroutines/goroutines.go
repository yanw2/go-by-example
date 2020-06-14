package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Call a function f(s) in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a goroutine, use `go f(s)`. This new goroutine will
	// execute concurrently with the calling one
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutine now. Wait for them
	// to finish (for a more robust approach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}
