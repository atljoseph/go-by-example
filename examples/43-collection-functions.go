package main

import "strings"
import "fmt"

// "go" does not support generics
// instead, we must use "collection functions"
// to filter collections by a predicate
// OR transform a collection into another format

// below are some collection functions
// for the type []string (string slice)

/*
Index ... returns the index of string inside of []string, and -1 if not found
*/
func Index(entire []string, part string) int {
	for idx, val := range entire {
		if val == part {
			return idx
		}
	}
	return -1
}

/*
Include ... returns true if a []string contains a string
*/
func Include(entire []string, part string) bool {
	return Index(entire, part) >= 0
}

/*
Any ... returns true if a []string contains a string
*/
func Any(entire []string, predicate func(string) bool) bool {
	for _, val := range entire {
		if predicate(val) {
			return true
		}
	}
	return false
}

/*
All ... returns true if all strings in []string matches the predicate
*/
func All(entire []string, predicate func(string) bool) bool {
	for _, val := range entire {
		if !predicate(val) {
			return false
		}
	}
	return true
}

/*
Filter ... returns a []string of items which match the predicate
*/
func Filter(entire []string, predicate func(string) bool) []string {
	matches := make([]string, 0)
	for _, val := range entire {
		if predicate(val) {
			matches = append(matches, val)
		}
	}
	return matches
}

/*
Map ... returns a []string of transformed items from the source []string
*/
func Map(entire []string, transformer func(string) string) []string {
	transformed := make([]string, len(entire))
	for idx, val := range entire {
		transformed[idx] = transformer(val)
	}
	return transformed
}

func main() {

	// test everything out!

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(part string) bool {
		return strings.HasPrefix(part, "p")
	}))
	fmt.Println(All(strs, func(part string) bool {
		return strings.HasPrefix(part, "p")
	}))
	fmt.Println(Filter(strs, func(part string) bool {
		return strings.Contains(part, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}
