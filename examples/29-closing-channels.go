package main

import "fmt"

// sometimes a channel needs to be closed
// can be useful to communicate completeness to receivers

func main() {

	// when we have no more jobs, we will close the channel and call done
	jobs := make(chan int, 5)
	done := make(chan bool)

	// on a separate thread, continually listen for new "jobs"
	go func() {
		// this runs all the time (endless for loop)
		for {
			// this is a special form of the receive syntax
			// "more" gives a boolean which tell us if there are more values in the channel to read
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send some jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// uncommenting the code below will produce an error
	// panic: send on closed channel
	// jobs <- 5

	// this causes us to wait on completion
	<-done
}
