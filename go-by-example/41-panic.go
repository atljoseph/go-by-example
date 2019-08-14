package main

import "os"

// use builtin "panic(error)" to throw an error
// used especially when they are not expected
// and when we do not want to handle the error in the current scope

func main() {

	// hardcoded "panic"
	// panic("a problem")

	// sometimes when we get an error, we know how to handle it
	// but if we don't, we use "panic(err)", as below
	// note the use of idiomatic go below - it uses the error return
	_, err := os.Create("/tempster/filester.file")
	if err != nil {
		panic(err)
	}
}
