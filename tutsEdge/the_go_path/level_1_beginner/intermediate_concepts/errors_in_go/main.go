package main

import (
	"errors"
	"fmt"
)

// Go is different when it comes to errors. As Go doesn't do exceptions at all. You could use
// panic as other langs use "Try and Catch" structures, but go devs don't do this. Instead we
// create errors and pass them around as values. That's why typical funcs return 2 or more values
// with the last one being an error

// With more complex functios we'll probably multple error types that we want to use.
// handling it different depending on that error type.
var (
	ErrIsLessThanZero = errors.New("number less than zero")
	ErrIsNotEven      = errors.New("number is not even")
)

func IsEven(num int) error {
	if num%2 != 0 {
		// return errors.New("number not even")
		return ErrIsNotEven
	}

	// remember that when we are working with errors we need to resolve both the sad and
	// happy paths. So above we return our new error but if the func gets to this point then
	// we return nil
	return nil
}

func main() {

	fmt.Println("Errors in Go!")
	println()

	err := IsEven(25) // we should be seeing and handling errors at the point that they
	// might occur. Go will allow us to simply run the function and do nothing with it.
	if err != nil { // with named errors we can do an 'assertion'
		if err == ErrIsNotEven {
			fmt.Println("Err is not even thrown")
		}
		// errors are values and have methods
		fmt.Println(err.Error())
	}

}
