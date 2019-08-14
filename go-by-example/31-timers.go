package main

import "fmt"
import "time"

func main() {

	// NewTimer returns a channel
	// when the time expires, the channel is notified
	timer1 := time.NewTimer(2 * time.Second)

	// start a timer
	// "timer1.C" blocks the timer's C channel until the time is expired
	fmt.Println("Timer 1 started")
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// look, this timer is cancellable
	// it operates on a different thread
	// thus, it is not blocking
	// ... and cancellable
	timer3 := time.NewTimer(3 * time.Second)
	go func() {

		// wait on timer3 to finish
		fmt.Println("Timer 3 started")
		<-timer3.C
		fmt.Println("Timer 3 expired")
	}()

	// define timer2 to simulate another async process
	fmt.Println("I'm hungry!")
	timer2 := time.NewTimer(time.Second)

	// wait on timer2 to finish
	fmt.Println("Timer 2 started")
	<-timer2.C
	fmt.Println("Timer 2 expired")

	//  then stop timer3
	stopped3 := timer3.Stop()
	if stopped3 {
		fmt.Println("Timer 3 stopped")
	}

}
