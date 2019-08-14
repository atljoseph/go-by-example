package main

import (
	"fmt"
	"sort"
)

// "sort" is a good package, but it is very basic
// for example, we do not yet have a way
// to sort strings by their length
// we can only use the ascending alphabetic sort currently

// this is a helper type for sorting
type sortByLength []string

// satisfy "sort.Interface"
// "Len()" and "Swap()" will usually always be the same
// but "Less()" here is imortant,
// as this is the heart of our sorting definition
// these are extension functions we need to define in order to properly sort our array
func (s sortByLength) Len() int {
	return len(s)
}
func (s sortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortByLength) Less(i, j int) bool {
	// ascending
	// return len(s[i]) > len(s[j])
	// descending
	return len(s[i]) > len(s[j])
}

func main() {

	// implement sort by length
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(sortByLength(fruits))
	fmt.Println(fruits)
}
