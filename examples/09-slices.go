package main

import "fmt"

func main() {

	// slices are like arrays, but different
	// while arrays are defined by their size and the type of the items they contain,
	// slices are defined only by the type of items they contain
	// must use the built-in make() command
	// with required length argument (but can be 0)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// get / set just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// same length as when originally created
	fmt.Println("len:", len(s))

	// add some elements to the array
	// these must be of the same type as the slice's type
	// NOTE: append does not mutate the slice in place
	// you MUST capture the product with a new assignment
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// new length
	fmt.Println("len:", len(s))

	// copy a slice
	// copy into c from s
	// first, intialize c
	c := make([]string, len(s))
	// this ends up with nothing copied:
	// c := make([]string, 0)
	// nothing to copy into
	// so evidently copy does not reserve new memory
	// in the receiver for the memory from the donator
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices support "slice" operator
	l := s[2:5]
	fmt.Println("sl1:", l)

	// get everything BEFORE the fifth index
	l = s[:5]
	fmt.Println("sl2:", l)

	// get everything AFTER the second index
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize in the same line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slicesa can be multidimensional, like arrays
	// notice the use of := inside of the for loop
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
