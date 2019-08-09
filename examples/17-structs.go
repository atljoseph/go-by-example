package main

import "fmt"

// "structs" are useful for grouping data together
// it is essentially a way to compose a new type
// similar to a JSON in js
// similar a type in typescript
type person struct {
	name string
	age  int
}

func main() {

	// can assign values in order of occurrence in the struct
	fmt.Println(person{"Bob", 20})

	// can also assign values explicitly by name
	// similar to js
	fmt.Println(person{name: "Alice", age: 30})

	// undeclared values are zero-valued
	fmt.Println(person{name: "Fred"})

	// a pointer rto the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// access fields by using "."
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// can assign a pointer to a new variable
	// struct pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age++
	fmt.Println(sp.age)
}
