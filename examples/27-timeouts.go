package main

import "time"
import "fmt"

// time and timeouts are important in any language
// because many processes depend on precise timing

func main() {

	// set up a goroutine with an anonymous function
	// that takes a few seconds to complete
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// set up a timeout with "select"
	// this one will always timeout
	// when "c1" takes over 1 second
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// set up another goroutine with an anonymous function
	// that takes a few seconds to complete
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// set up a timeout with "select"
	// this one will always timeout
	// when "c1" takes over 1 second
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
