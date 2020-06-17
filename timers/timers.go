package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future. You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel C until it sends a value indicating that the
	// timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used `timer.Sleep`. One reason a timer may be useful
	// is that you can cancel the timer before it fires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timer enough time to fire, if it ever was giong to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}
