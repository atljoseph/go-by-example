package main

import "fmt"

// "go" supports methods defined on struct types
// this is a lot how javascript works in a way
type rect struct {
	width, height int
}

// now that we have a struct,
// we can add methods to it
// notice that "area()" uses a pointer to r (r *rect)
// go automatically dereferences pointers in method calls
func (r *rect) area() int {
	return r.width * r.height
}

// notice that "perim()" does NOT use a pointer
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// playing around
// this function modifies the original scope of "r"
// requires the usage of the pointer
func (r *rect) smaller() {
	r.width = r.width - 1
	r.height = r.height - 1
}

func (r *rect) bigger() {
	r.width = r.width + 1
	r.height = r.height + 1
}

func main() {

	// do NOT use the pointer
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// now use the pointer
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	// run "smaller" on "r" and then recheck the area nad perimeter
	r.smaller()
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// run "bigger" on "rp" and then recheck the area nad perimeter
	rp.bigger()
	rp.bigger()
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

}
