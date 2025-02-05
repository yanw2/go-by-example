package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine. Note that a WaitGroup must be passed to functions
// by pointer.
func worker(id int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutine launched here to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they're done.
	wg.Wait()
}
