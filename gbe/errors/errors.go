package main

import (
	"errors"
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// In Go its idiomatic to communicate errors via an explicit, separate return value
// This contrasts with the exceptions concept in many other langs like Java and Ruby
// and the overloaded single result/error value sometimes used in C. Go's approach
// makes it easy to see which functions return errors and to handle them using
// the same language constructs employed by any other, non-error tasks.

// by convention, errors are the last return value and have type error, a built-in
// interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given string.
		return -1, errors.New("can't work with 42")
	}
	// a Nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// It's also possible to use custom types as errors by implementing the Error()
// method on them. Here's a variant on the example above that uses a custom
// type to explicitly rep an arg error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		// in this case we use the &argError (pointer to our custom type) to build
		// a new struct, supplying values for  the two fields needed.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// The two loops below test out each of our error-returning functions. Note
	// that the use of an inline error check on the if line is a common idiom in
	// Go code. (as I mentioned earlier)
	for _, i := range []int{63, 42, 66} {
		if r, e := f1(i); e != nil {
			p("f1 failed:", e)
		} else {
			p("f1 worked:", r)
		}
	}

	for _, i := range []int{63, 42, 66} {
		if r, e := f2(i); e != nil {
			p("f2 failed:", e)
		} else {
			p("f2 worked:", r)
		}
	}

	// If you want to programatically use the data in a custom error, you'll
	// need to get th error as an instance of the custom error type via type
	// assetion
	for _, i := range []int{63, 42, 66} {
		_, e := f2(i)
		if ae, ok := e.(*argError); ok {
			p(ae.arg)
			p(ae.prob)
		}
	}
}

// I can find more info here: https://go.dev/blog/error-handling-and-go
