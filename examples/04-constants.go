package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	// var, or const
	// constants can be character, string, boolean, numeric
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// usage of const here will produce an error
	e := math.Sin(n)
	fmt.Println(e)
	fmt.Println(int64(e))

	var f = math.Acosh(n)
	fmt.Println(f)

}
