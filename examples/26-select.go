package main

import "fmt"
import "time"

// at first glance,
// select looks just like switch
// but beyond the code structure, it is quite different
// instead of branches representing a set of logic paths
// where only one path will ultimately be followed upon execution
// "select" empowers developers to condense logic
// for multiple related channels
// into one statement
// each "case" reads from a channel

func main() {

	// setup channels
	c1 := make(chan string)
	c2 := make(chan string)

	// anon function on a new goroutine
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// anon function on another goroutine
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// defaine what happens when each of the channels receives a message
	// this select will only take 2 seonds to exec, because c1 and c2 exec simultaneously
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received 1", msg1)
		case msg2 := <-c2:
			fmt.Println("received 2", msg2)
		}
	}

	// i could see how the select could lead
	// to logic where many complex channels and operations
	// could cause the select to act in effect like a switch
	// but that is not the intended usage
}
