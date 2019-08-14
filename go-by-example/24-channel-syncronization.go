package main

import "fmt"
import "time"

// use channels to syncronize execution across concurrent goroutines
// this example uses a blocking receive to wait for a goroutine to finish
// use a "WaitGroup" to wait for many goroutines to finish
// this example is very similar to await in C# and other languages
// except that it is go's version of accomplishing this task

// define a function that will take in a chanel as a argument
// this function will use the channel
// to notify another goroutine of completion
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// notify that we are in fact done
	done <- true
}

func main() {
	fmt.Println("start")

	// make a channel a nd pass it into a function on a new goroutine
	done := make(chan bool, 1)
	go worker(done)

	// block the main channel until the channel receives a notification
	// note: if the "<-done" line were omitted, then the "go worker(done)" thread would not even execute"
	r := <-done
	fmt.Println(r)

	fmt.Println("end")
}
