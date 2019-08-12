package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// this is an alternative to managing state via "mutex"

// each peice of data should be owned by one goroutine
// the goal of this script
// is to manage state on one goroutine
// then use 100 goroutines to read the state
// then use 10 goroutines to write to the state

// define a read operation metadata struct
type readOp struct {
	key  int
	resp chan int
}

// define a write operation metadata struct
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// count the number of operations we perform
	var readOps uint64
	var writeOps uint64

	// define 2 channels that will be used
	// to communicate read and write request data
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// a single goroutine should "own" the state
	go func() {
		// the state variable is "private"
		// to the scope of this anonymous goroutine
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// create 100 goroutines
	// each goroutine will execute a read
	// and send it to the state goroutine
	for r := 0; r < 100; r++ {
		go func() {
			// forever loop
			for {

				// ok, we "read" something
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}

				// send the value down the channel
				// then await the action inside the select (above)
				// (above - to read.resp)
				reads <- read
				<-read.resp

				// update our counters and sleep
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// create 10 goroutines to write
	for w := 0; w < 10; w++ {
		go func() {
			for {

				// ok, we "wrote" something
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}

				// send the value down the pipe
				// then wait on the update in the select
				// (above - to write.resp)
				writes <- write
				<-write.resp

				// update our counter and sleep
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// give the program time to execute many operations
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
