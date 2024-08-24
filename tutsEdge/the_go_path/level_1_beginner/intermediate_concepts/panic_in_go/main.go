package main

import (
	"errors"
	"fmt"
)

func IPanicEasily() {
	defer func() {
		fmt.Println("I am deferred")
	}()
	panic("I panic easily")
}

func MyAwesomeFunction() (err error) {
	defer func() {
		fmt.Println("My awesome function defered")
		if r := recover(); r != nil {
			fmt.Println("recovered from a small panic") // here with defer we 'recover' from our
			// panic, rather than a complete crash.
			// this also helps us to identify errors for better debugging
			err = errors.New("I panic easily panicked")
		}
	}()
	IPanicEasily()
	return nil
}

func main() {
	defer func() {
		fmt.Println("Main function deferred")
	}()
	fmt.Println("Panic! In the Go App")
	// panic("arghhhhh") // panic func is a builtin func that allows us to terminate immediately
	// when we find something we can't recover from or an error we want to kill the app at for
	// debugging.

	// IPanicEasily() // the  stacktrace on this one takes us back to the erring func
	MyAwesomeFunction() // again we see the trace  goes all the way up the chain to the orig error

	if err := MyAwesomeFunction(); err != nil {
		fmt.Println("My awesome function returned an error")
		fmt.Println(err.Error()) 
	}
	fmt.Println("We've reached the end of 'main' successfully") // if we don't recover, we never
	// get to this  point.
}

// we show the deferred examples above because regardless  if we panic or have failures, the
// deferred funcs are still being called, so we can clean up even if our application fails.
// This means we can also use defer to 'recover' from panics

// The builtin 'recover' func is great for debugging and recovering from failures.
