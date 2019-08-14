package main

import "fmt"
import "os"

// "defer" is used to delay execution of the defered function
// to the end of the enclosing function
// this allows a visual reorganization of cleanup code, etc

func main() {

	// basic example
	// create, write, and close a file
	// "closeFile(f)" will be excuted at the end
	// Note: "f" is a pointer
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
