package main

import "fmt"
import "time"

// we use "timers" to schedule
// one-time-only actions in the future
// BUT "tickers" allow us to schedule
// actions in the future
// repeated at regular intervals

func main() {

	// "tickers" act the same way as "timers"
	// "time.NewTicker(duration)" returns a new channel
	// every time each new interval expires,
	// this channel is sent a value
	ticker := time.NewTicker(500 * time.Millisecond)

	// fire this function every time the ticker ticks
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// sleep with a block for 1.6 seconds
	// then stop the ticker
	time.Sleep(2600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stoped")
}
