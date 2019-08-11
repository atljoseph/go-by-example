// channels are the pipes that connect concurrent goroutines
// channels allow us to send results from one goroutine to another goroutine
// by default, channels are unbuffered
// that means that the send (messages <- "ping")
// can ONLY occur if there is a corresponding receive (msg := <-messages)
// there is a thread block that occurs if no corresponding receive is made
// "<-messages" is REALLY similar to "await", except that we are not awaiting a function result
// we are waiting on the async function to mutate the scope

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
