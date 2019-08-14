package main

import "fmt"

// when channels are used as function parameters
// you need to specify each channel as a send-channel or a receive-channel
// this increases the type safety

// define a function that sends a message into a receiving channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// define a function that reads the last value out
// a compile-time error results from trying to receive from "pongs" anywhere in this function
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	// define channels to hold our ping and pong messages
	// we will define a ping, then call pong
	// to simulate a queue or some other buffered service
	// these channels will hold a maximum of 1 value
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// ping a new message
	ping(pings, "passed message")

	// then pong it out
	pong(pings, pongs)

	// wait on the pongs to be populated
	fmt.Println(<-pongs)
}
