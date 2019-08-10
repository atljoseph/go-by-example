// channels are the pipes that connect concurrent goroutines
// channels allow us to send results from one goroutine to another goroutine

package main

import "fmt"

func main() {

	// create a new channel to pipe a string value from a goroutine
	messages := make(chan string)

	// send a value into a channel by using the  "<-" syntax
	// this is done asyncronously
	go func() { messages <- "ping" }()

	// capture the result of the goroutine
	// here, the "<-" syntax receives the value from the messages channel
	msg := <-messages

	// print the result
	fmt.Println(msg)

	// by default, sends and receives ("<-") block
	// until both the sender and the receiver are both ready
	// thread blockage is how we were able to capture the value "ping" at the end of the script
	// there are other methods of syncronization
}
