package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// channel communications is how state is managed in "go"
// this example uses the "sync/atomic" package

func main() {

	// define an unsigned int to represent our counter
	// unsigned becausae it will always be positive
	var ops uint64

	// start 50 goroutines
	// which will increment the counter once each millisecond
	// since we are on 50 separate goroutines,
	// we need to increment the counter
	// use "atomic.AddUint64()" to increment a uint64 via "sync/atomic"
	// by using it's pointer "&ops"
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let it run for a second or so
	time.Sleep(time.Second)

	// check the number of operations
	// use "atomic.LoadUint64()" to to safely get a copy of the instantaneous value of "opsFinal"
	// using "&ops" directly might have a negative effects
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
