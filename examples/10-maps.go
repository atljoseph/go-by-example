package main

import "fmt"

func main() {

	// maps are analogous to a dynamic in C#, or a JSON in JS
	// called hashes or dictionaries in other languages
	// use make() to create an object in memory
	// syntax: map[<key-type>]<val-type>
	m := make(map[string]int)

	// get & set (just like slices / arrays)
	// maps have the possibility of non-numeric keys
	// set
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// get example
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len operates on maps in a similar way as slices
	// returns the number of key-value pairs
	fmt.Println("len:", len(m))

	// delete a member from a map
	delete(m, "k2")
	fmt.Println("map", m)

	// presence of key in map
	// when declaring, use the "blank identifier (_)"
	// to indicate a blank spot for a variable
	//  whose value you are not interested in
	// here we had to use it
	// because the get function returns multiple values
	// we are checking "k2" since we deleted it above
	// no need to use "_" if we potentially want both values
	// "isPresent" (below) is useful in order to
	// distinguish between "0 value" items and items missing from the map
	_, isPresent := m["k2"]
	fmt.Println("isPresent:", isPresent)

	// declare and initialize a map in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
