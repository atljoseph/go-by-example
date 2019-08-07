package main

import "fmt"

func main() {

	// type infered from initil value
	var a = "initial"
	fmt.Println(a)

	// can initialize and set more than one variable
	var b, c int = 1, 2
	fmt.Println(b, c)

	// type infered from initil value
	var d = true
	fmt.Println(d)

	// "zero" value
	var e string
	fmt.Println(e)

	// "zero" value
	var f int
	fmt.Println(f)

	// "zero" value
	var g bool
	fmt.Println(g)

	// shorthand for declaring and initializing
	h := "apple"
	fmt.Println(h)
}
