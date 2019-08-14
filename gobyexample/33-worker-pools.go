package main

import "fmt"
import "time"

// "worker pools"
// depends upon "goroutines" and "channels"

// define a worker function
// this will be used concurrently by many goroutines
func worker(id int, jobsIn <-chan int, resultsOut chan<- int) {
	for j := range jobsIn {
		fmt.Println("worker", id, "started job", j)

		// simulate doing something in this job
		time.Sleep(time.Second)

		// finish up the job and send result to the "resultsOut" channel
		fmt.Println("worker", id, "finished job", j)
		resultsOut <- j * 2
	}
}

func main() {

	// jobs and results channels
	// need a channel for the jobs and a channel for the results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// use concurrecnt goroutines
	// to send our workers some work
	// note: the first job is nott usually executed on the first goroutine thread
	for w := 1; w <= 3; w++ {
		// fmt.Println("create work", w)
		go worker(w, jobs, results)
	}

	// receive results from each channel
	// similar to "await a response from each task" in c#
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// close the jobs channel
	// this step is optional here
	// but good practiche
	close(jobs)

	// read data from the "results" channel
	for a := 1; a <= 5; a++ {
		// if we do not run this line,
		// then no "workers" will run
		<-results
	}

}
