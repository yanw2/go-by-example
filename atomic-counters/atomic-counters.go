package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// The primary mechanism for managing state in Go is communication over channels.
	// Here we'll look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.

	// We'll use an unsigned integer to represent our (always positive) counter.
	var ops uint64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly 1000 times.
	// To aotmically increment the counter we use `AddUnit64`, giving it the memory address of our ops counter
	// with the & syntax.
	for i := 0; i < 50; i++ {
		// Increment WaitGroup counter by 1 for each goroutine
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done.
	wg.Wait()

	// It's safe to access ops now because we know no other goroutine is writing to it. Reading aotmics safely
	// while they are being updated is also possible, using functions like `atomic.LoadUint64`.
	fmt.Println("ops:", ops)
}
