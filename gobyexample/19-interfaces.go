package main

import (
	"fmt"
	"math"
)

// interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// structs we will implement "geometry" on
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// extend rect and circle
// notice how these extensions on "rect" and "circle"
// are not yet associated to "geometry" yet
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// size up the geometries
// this is where "geometry" is associated to rect or circle
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
