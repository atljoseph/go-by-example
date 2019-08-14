package main

import "fmt"

// for basic channel usage
// sends and receives are blocking operations
// this means that the thread they run on
// waits for their execution to complete before it advances
// one way to get around blocking is to use a "select"
// in combination with a "default" branch
// this is a way to make a non-blocking goroutine call

func main() {

	// non-buffered channels
	messages := make(chan string)
	signals := make(chan bool, 1)

	// "messages" has no value set yet
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// when we try to send "msg" to "messages"
	// we do not get an error, but the message is not sent
	// because (at the time of execution of the "select")
	// there is no receiver for the message on the "messages" channel
	// therefore the default branch is selcted
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// comment this line in or out )if you like)
	// to alter the outcome branch below
	signals <- true

	// down here, message is received from, but not above
	// however, there is no message yet in "messages"
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
