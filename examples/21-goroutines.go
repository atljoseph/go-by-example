package main

import "fmt"

// goroutines are lightweighht threads that are employed by go to exec asyncronous tasks

func f(from string) {
	// example shows the concurrency after the 1st line,
	// but i did not see it unti at least the 10th (and sometimes 20th)
	for i := 0; i < 25; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// we can call a function in the normal way
	// this runs the function syncronously
	f("direct")

	// OR we can call it with a go routine
	// which executes concurrentlywith the calling scope
	go f("goroutine")

	// goroutines can also be started from anonymous functions
	go func(msg string) {
		fmt.Println(msg)
	}("going ... Press any key ...")

	// wait for user entry and then exit
	// if we don't do the wait, then we will not see ethe output in the terminal
	fmt.Scanln()
	fmt.Println("done")

}
