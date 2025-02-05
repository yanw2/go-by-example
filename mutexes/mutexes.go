package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// For more complex state we can use a mutex to safely access data across multiple goroutines.
	// For our example the `state` will be a map. This `mutex` will synchronize access to `state`.
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	// We'll keep track of how many read and write operations we do
	var readOps uint64
	var writeOps uint64

	// Here we start 100 goroutines to execute repeated reads against the state, once per miliseconds
	// in each goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// For each read we pick a key to access, `Lock()` the `mutex` to ensure exclusive access to
				// the `state`, read the value at the chosen key, `Unlock()` the mutex, and increment the
				// `readOps` count.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We'll also start to goroutines to simulate writes, using the same pattern we did for reads.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(10)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	// Let the 10 goroutines work on the `state` and `mutex` for a second
	time.Sleep(time.Second)

	// Take and report final operation counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)

	// With a final lock of state, show how it ended up.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	// Running the program shows that we executed about 90,000 total operations against our
	// mutex-synchronized state.
}
