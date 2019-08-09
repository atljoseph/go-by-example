package main

import "fmt"

// define a function that returns a pair of ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// get both return values from the function
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// get on the 2nd return value and ignore the other
	_, c := vals()
	fmt.Println(c)

	// common usage of this feature is to pass back errors
	// and for ranges/loop, index and value
}
