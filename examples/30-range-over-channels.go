package main

import "fmt"

// "range" can be used with "for"
// to iterate over slice, array, or map values
// BUT it can also be used with channels
// to iterate over values received by the channel
// note: channel must be closed first

func main() {

	// iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// if we comment this line out, then error occurs
	// fatal error: all goroutines are asleep - deadlock!
	close(queue)

	// range iterates over all values in the channel
	// like running "append(myArray, <-queue)" many times
	// mote: range over channel depends on a channel being closed
	for elem := range queue {
		fmt.Println(elem)
	}
}
