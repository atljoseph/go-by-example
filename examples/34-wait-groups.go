package main

import (
	"fmt"
	"sync"
	"time"
)

// define a function which accepts a waitGroup pointer
// notice the *. When you see this, think &
func worker(id int, wg *sync.WaitGroup) {

	// simulate an async process
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// notify the "wait group" that the workrer is done
	// calling wg.Done(), decrements the WaitGroup counter
	wg.Done()
}

func main() {

	// increment the waitGroup counter for each worker's gotroutine
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// wait until wg.Done() is called on all the waitGroups
	// the thread will no longer be blocked until the counter is back to 0
	wg.Wait()

	// the execution of goroutines and waitgroups will likely change every time
}
