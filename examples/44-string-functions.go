// go uses UTF-8 encoding on strings

package main

import "fmt"

// alias "strings as "s"
import s "strings"

// alis "fmt.Println" as "p"
var p = fmt.Println

func main() {

	// these are methods which come from the "strings" package itself
	// and are not properties that exist on strings
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// get length in bytes
	p("Len: ", len("hello"))

	// get a byte by it's index
	p("Char:", "hello"[1])

	// see strings, bytes, runes, and characters
}
