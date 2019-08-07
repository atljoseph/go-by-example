package main

import "fmt"

func main() {

	// arrays are somewhat limiting
	// their type is defined by their type
	// once created, an array's dimensions cannot change

	// must know the size of an array when you create it
	// is initialized with all 0's
	var a [5]int
	fmt.Println("init:", a)

	// get and set values as you would expect
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// builtin "len()" length operator
	fmt.Println("len:", len(a))

	// declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// multi-dimensional composition
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
