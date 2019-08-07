package main

import "fmt"

func main() {

	// only loop in "go" is the for loop

	// basic loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// intial / condition / after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for loops without conditions will loop without end
	// must use "break" or "return"
	for {
		fmt.Println("loop")
		break
	}

	// use "continue" to skip to next loop iteration
	for n := 0; n <= 5; n++ {
		// don't have to use parenthesis around if statements
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
