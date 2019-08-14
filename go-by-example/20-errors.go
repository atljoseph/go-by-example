// go communicates errors through function return values
// specifically via multiple returns
// by convention, the last return value is the error

package main

import "fmt"
import "errors"

// define a function that returns an errors sometimes
// when there is no error, then return some value
func functionReturningBasicError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant' work with 42")
	}
	return arg + 3, nil
}

// define an "argError" struct
// and define a function that extends this error type
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// define a function that returns the custom error
// use "&" to instantiate a new struct of argError ("&argError" is like "new argError()")
func functionReturningCustomError(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant' work with 42"}
	}
	return arg + 3, nil
}

func main() {

	// the usage of the following code is common in go:
	// if r, e := myFunc(); e != nil {
	//	// there was an error
	// } else {
	//	// succeeded
	// }

	// basic error
	for _, i := range []int{7, 42} {
		if r, e := functionReturningBasicError(i); e != nil {
			fmt.Println("functionReturningBasicError failed", e)
		} else {
			fmt.Println("functionReturningBasicError worked", r)
		}
	}

	// custom error
	for _, i := range []int{7, 42} {
		if r, e := functionReturningCustomError(i); e != nil {
			fmt.Println("functionReturningCustomError failed", e)
		} else {
			fmt.Println("functionReturningCustomError worked", r)
		}
	}
}
