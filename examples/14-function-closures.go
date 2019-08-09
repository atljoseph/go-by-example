package main

import "fmt"

// Anonymous Functions are declared without a specific name

// Closures attach the scope of the function creator
// to the scope of the function instaniator

// "i" still gets incremented each time the resulting function is called
// this is analogous to creating a class in old js
// very similar to a class
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// newInt1, with it's own scope
	// newInt2, with it's own scope
	newInt1, newInt2 := intSeq(), intSeq()
	fmt.Println(newInt1())
	fmt.Println(newInt1())
	fmt.Println(newInt1())
	fmt.Println(newInt2())

	// newInts, with yet another scope
	newInts := intSeq()
	fmt.Println(newInts())
}
