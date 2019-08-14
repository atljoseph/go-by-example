package main

import "bytes"
import "fmt"
import "regexp"

// go supports regular expressions by default

func main() {

	regexpPattern := "p([a-z]+)ch"

	// test if a string matches a pattern
	match, _ := regexp.MatchString(regexpPattern, "peach")
	fmt.Println("MatchString:", match)

	// above was a very basic example
	// most of the time
	// you will need to compile an optimized regexp struct
	r, e := regexp.Compile(regexpPattern)
	fmt.Println("Compile:", r, e)

	// use "MustCompile" when compilation failure
	// should in turn cause appilcation failure
	// this method panics upon failure
	// and has only one return value
	rc := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("MustCompile:", rc)

	// here are some example methods that we can use
	// on the compiled regexp struct

	// does a string match a pattern?
	fmt.Println("MatchString:", r.MatchString("peach"))

	// this is what "MatchString" does under the hood
	fmt.Println("Match:", r.Match([]byte("peach")))

	// find the first matching text in string
	fmt.Println("FindString:", r.FindString("peach punch"))

	// find all matching text in string
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", -1))

	// start searching from the middle of the string
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", 4))

	// find the match's start and end indexes
	fmt.Println("FindStringIndex:", r.FindStringIndex("peach punch"))

	// return information about submatches
	// this means partial and full regexp matches
	fmt.Println("FindStringSubmatch:", r.FindStringSubmatch("peach punch"))

	// same as "FindStringSubMatch" (above)
	// but returns start and end indexes
	fmt.Println("FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))

	// and the "All" version of "FindStringSubMatchIndex"
	fmt.Println("FindAllStringSubmatchIndex:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// we can also use "Replace"
	// according to the regexp pattern "r"
	// replace all matches with "<fruit>"
	fmt.Println("ReplaceAllString:", r.ReplaceAllString("a peach", "<fruit>"))

	// we can use a function to replace text
	// use the "...Func" variant
	// can see where usage of a closured function can be useful here
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("ReplaceAllFunc:", string(out))
}
