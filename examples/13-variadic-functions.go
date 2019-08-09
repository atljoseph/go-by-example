package main

import "fmt"

// Variadic Function:
// A function that can be called with any number of trailing arguments

// define a function
// this will take any number of integers as a list of arguments
func sum(desc string, nums ...int) int {
	fmt.Println(desc)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum("No soup for you!", 1, 2))
	fmt.Println(sum("Where you at?", 1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum("Freeze, this is a stick up!", nums...))
}
