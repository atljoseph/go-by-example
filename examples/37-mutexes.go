package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// "mutexes" allow us to safely accrss data accress multiple goroutines
// "mutex.Lock()" gives us exclusive access to the state when called

func main() {

	// for this example, the state will be a map
	var state = make(map[int]int)

	// use mutex to syncronize access to state
	var mutex = &sync.Mutex{}

	// we want to keep track of how many read and write operations are performed
	// "unint64" because the number will always be positive
	// ... and it might get pretty large
	var readOps uint64
	var writeOps uint64

	// exec 100 goroutines
	// each one will access the state
	// and increment a counter
	// inside of an endless for loop
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// pick a random number as the key
				// between 1 and 5
				key := rand.Intn(5)

				// request exclusive access to the state
				mutex.Lock()

				// do something with the state to read it out
				total += state[key]

				// relenquish exclusive access to state
				mutex.Unlock()

				// atomically update our counter
				// to track our reads
				atomic.AddUint64(&readOps, 1)

				// pause for a millisecond
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// exec 10 goroutines to simulate a bunch of writes
	for w := 0; w < 100; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the 10 goroutines we just started
	// to work on the state and mutex
	// for one second
	time.Sleep(time.Second)

	// to safely read our data out of the pointer addresses
	// use sync.LoadUint64(pointer)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}
