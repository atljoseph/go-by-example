package main

import (
	"errors"
	"fmt"
)

// example of:
// ----------------------------
// - errors and idiomatic error flow
// - panic usage
// - multiple return
// - variadic function w/ optional arrgument

// define a function
// with an optional argument
// "go" has no way to acheive this directly
// but we can use variadic function
// in an imaginative way
// to accomplish the same affect
func addOne(startWith ...int) (int, error) {

	// validate startWith
	l := len(startWith)
	if l > 1 {
		return 0, errors.New("too many args - pleasse use one for this function")
	}

	// default to 0 if no args
	s := 0
	if l == 1 {
		s = startWith[0]
	}
	return s + 1, nil
}

// implement same function as above
// panic upon error
func addOnePanic(startWith ...int) int {
	r, e := addOne(startWith...)
	if e != nil {
		panic(e)
	}
	return r
}

func main() {

	// -----------------------------
	// panic if there is an error
	// -----------------------------

	// this one returns 1
	count1 := addOnePanic()
	fmt.Println("count1", count1)

	count1 = addOnePanic(count1)
	fmt.Println("count1", count1)

	count1 = addOnePanic(count1)
	fmt.Println("count1", count1)

	// -----------------------------
	// check the error after running
	// -----------------------------

	count2_1, err1 := addOne()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("count2", count2_1)

	count2_2, err2 := addOne(count2_1)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("count2", count2_2)

}
