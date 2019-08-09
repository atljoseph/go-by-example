package main

import "fmt"

// define a function
// types are mandatory
func plus(a int, b int) int {
	return a + b
}

// define a function
// when a, b, and c are all int's
// we can write as below
func plus3(a, b, c int) int {
	return a + b + c
}

func main() {
	// call the function you just created
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// and the other
	res = plus3(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
