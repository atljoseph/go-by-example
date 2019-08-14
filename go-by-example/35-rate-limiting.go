package main

import "fmt"
import "time"

// "rate limiting" for "channels"
// is like "piping" and "throttling" for "rxjs observables"

func main() {

	// simulate 5 requests
	// handle the requests via a channel called "requests"
	fmt.Println("input requests simultaneously")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
		fmt.Println("request in", i, time.Now())
	}
	close(requests)

	// define a channel which will receive a value every 200 milliseconds
	// read the values out with a rate limiter
	// note: this rate limiter is thread-blocking
	fmt.Println("read requests via thread-blocking rate limiter")
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request out", req, time.Now())
	}

	// --------------------

	// define a "bursty rate limiter"
	// essentially, we are building a channel with metadata
	// that will help with processing below

	// add 3 batched times
	// these times will be virtually the same
	fmt.Println("input requests simultaneously")
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// on an async thread (goroutine)
	// place items into the limiter via rate limiting
	// these times will be 200 ms apart
	fmt.Println("add items to burstLimiter on a goroutine")
	go func() {
		for t := range time.Tick(2000 * time.Microsecond) {
			// fmt.Println("add rate limiter timing to channel")

			// COMMENT OUT ONE OF THE FOLLING OPTIONS:

			// (option 1)
			// after the initial burst
			// start reading out 1 at a time
			// burstyLimiter <- t

			// (option 2)
			// after the initial burst
			// reload another burst
			for i := 0; i < 3; i++ {
				burstyLimiter <- t
			}
		}
	}()

	// define a requests channel with 5 slots
	// input 5 requests to it
	// then close the channel
	fmt.Println("input requests simultaneously")
	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
		fmt.Println("request in", i, time.Now())
	}
	close(burstyRequests)

	// read the requests out via rate limiting
	fmt.Println("read out requests via a dynamically-changing rate limiter")
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request out", req, time.Now())
	}

	// this didn't work in liu of above (see below)
	// burstyResponses := make(chan time.Time, 50)
	// go test(burstyLimiter, burstyRequests, burstyResponses)
	// <-burstyResponses
}

// this doesn't work for some reason
// when two goroutines are execed, they aren;t yet communicating
// func test(limiterOut <-chan time.Time, requestsOut <-chan int, responsesIn chan<- time.Time) {
// 	fmt.Println("read out requests via a dynamically-changing rate limiter")
// 	for req := range requestsOut {
// 		<-limiterOut
// 		responsesIn <- time.Now()
// 		fmt.Println("request", req, time.Now())
// 	}
// }
