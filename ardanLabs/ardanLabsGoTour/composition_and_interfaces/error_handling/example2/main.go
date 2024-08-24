package main

// Note the import of errors. We didn't do this on the first example since we built
// out our own errors interface and New func. But moving fwd we just use the stdlib
// package 'errors'
import (
	"errors"
	"fmt"
)

// So what if its important to know WHAT error value exists inside the err interface
// and not just if one exists or not? Then we an use error variables.

// Error variables provide a good mechanic to identify what specific error is being
// returned. They have an idiom of starting with the prefix Err and are based on the
// concrete type errorString from the errors package (o sea we just provide them with
// a string using the errors.New() func)

var (
	// ErrBadRequest is returned when there are problems with the request
	ErrBadRequest = errors.New("Bad Request")

	// ErrPageMoved is returned when a 301/302 is returned
	ErrPageMoved = errors.New("Page Moved")
)

func main() {

	/*
		In this application after the call to webCall is made, a check can be performed
		to see if there is a concrete value stored inside the err interfacde variable.
		If there is, then a switch statement is being used to determine which error it
		was by comparing err to the different ERROR VARIABLES

		In this case, the context of the error is based on which error variable was
		returned.
	*/

	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occured")
			return
		case ErrPageMoved:
			fmt.Println("The page moved")
			return
		default:
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Life is good")

}

// In this new version of webCall, the function returns one or the other error
// variable. This allows the caller to determine which error took place, o sea
// giving more context to the user.
func webCall(b bool) error {
	if b {
		return ErrBadRequest // again the use of early returns helps us not use 'else'
	}
	return ErrPageMoved
}

// Now what if an error variable is not enough context?? What if some special state
// needs to be checked, like with networking errors?? In these cases, a custom concrete
// error type is the answer and we'll see that in the next example.
