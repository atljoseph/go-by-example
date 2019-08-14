package main

import "fmt"

// by default, channels are unbuffered
// to be buffered means that the channel can accept
// a limited number of values without
// a corresponding receiver for each of the values

func main() {

	// make a channel that can store up to 2 values
	// can buffer with a length of 1, too
	// Note: make(chan string, 1) != make(chan string)
	messages := make(chan string, 2)

	// without needing a corresponding receive,
	// we can send these two values into the channel
	messages <- "buffered"
	messages <- "channel"

	// what happens if we try to send more than it can hold?
	// ERROR: fatal error: all goroutines are asleep - deadlock!
	// messages <- "extra"

	// and now we read the values out of the channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// now that we read some messages out,
	// we can send new ones in,
	// then read the new ones out again
	// each subsequent read/receive action gets the oldest message in memory
	messages <- "more1"
	messages <- "more2"
	fmt.Println(<-messages)
	messages <- "messages"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
