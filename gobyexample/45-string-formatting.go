package main

import (
	"fmt"
	"os"
)

// string formatting possibilities in go
// styled in the printf tradition
// here are some common tasks

// helper struct
type point struct {
	x int
	y int
}

func main() {

	// Printf()
	// prints a string as dsired to the terminal

	// define something we can print
	p := point{1, 2}

	// %v
	// value
	fmt.Printf("%v\n", p)

	// %+v
	// prints the value with the field names
	fmt.Printf("%+v\n", p)

	// %#v
	// go syntax representation
	fmt.Printf("%#v\n", p)

	// %T
	// variable type
	fmt.Printf("%T\n", p)

	// %t
	// boolean format
	fmt.Printf("%t\n", true)

	// %d
	// integer, standard base 10 formatting
	fmt.Printf("%d\n", 123)

	// %b
	// outputin binary format
	fmt.Printf("%b\n", 14)

	// %c
	// character representation
	fmt.Printf("%c\n", 33)

	// %x
	// hex encoding
	fmt.Printf("%x\n", 456)
	fmt.Printf("%x\n", "hex this")

	// %f
	// integers, floats, basic decimal formatting
	// note: "-" flag means to left-justify
	fmt.Printf("%d\n", 78.9)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// %e
	// %E
	// scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// %s
	// string
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("|%6s|%6s|\n", "foo", "bee")

	// %q
	// string, double-quoted
	fmt.Printf("%q\n", "\"string\"")

	// %p
	// pointer
	fmt.Printf("%p\n", &p)

	// -------------------------

	// Sprintf()
	// returns a string formatted as desired
	s := fmt.Sprintf("a %s", "strange string")
	fmt.Println(s)

	// -------------------------

	// Fprintf()
	// print to a file, buffer, etc
	// here, we write to stdError
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
