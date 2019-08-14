package main

import "fmt"

// "go" supports pointers
// A "pointer" is a reference to a variable's memory address

// define a function who receives a copy of the value passed as ival
func zeroval(ival int) {
	ival = 0
}

// define a function who receives the actual variable passed as iptr
func zeroptr(iptr *int) {
	*iptr = 0
}

// kitty example
type cat struct {
	name      string
	meowIndex int
}

func quietKitty(meowPtr *int) {
	*meowPtr = 0
}

func main() {

	// in go, when you call a function without a pointer
	// the paramter value(s) available inside the function
	// are actually copies of the original

	// to pass the actual variable and not a copy),
	// use a Pointer

	// our goal (below) is to "zero out" the value of "i"
	i := 1
	fmt.Println("inital:", i)

	// does nothing, because a copy of i is passed into the function
	zeroval(i)
	fmt.Println("zeroval:", i)

	// this function actually "zeroes out" the value of "i"
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointers can also be printed
	fmt.Println("pointer:", &i)

	// usage of pointers can dramatically increase performance in some cases

	// in this way, "quietKitty()" operates directly on myCat.meowIndex
	var myCat = cat{name: "kitty", meowIndex: 1}
	fmt.Println(myCat)
	quietKitty(&myCat.meowIndex)
	fmt.Println(myCat)

}
