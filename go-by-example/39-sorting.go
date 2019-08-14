package main

import "fmt"
import "sort"

// "sort" lib
// can use with builtin types, and also user-defined types
// NOTE: sorting happens in-place ... there is no return value from sort

func main() {

	// sort slice of strings (in-place)
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// sort slice of ints (in-place)
	ints1 := []int{7, 2, 4}
	sort.Ints(ints1)
	fmt.Println("Ints:", ints1)

	// test if strings are sorted in ascending order
	s := sort.IntsAreSorted(ints1)
	fmt.Println("Sorted:", s)

	// Note: All sorts above are ascending order
}
