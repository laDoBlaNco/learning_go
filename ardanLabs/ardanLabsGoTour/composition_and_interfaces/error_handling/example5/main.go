package main

// Always use the Error Interface
// One mistake Go devs tend to make is when they use the concrete error type and not the
// error interface for the return type for handling errors. If we were to do this, bad
// things could happen.

import (
	"fmt"
	"log"
)

// customError is just an empty struct for this example
type customError struct{}

// Error implements the error interface
func (c *customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types
func fail() ([]byte, *customError) { // this should be 'error' 
	return nil, nil
}

func main() {

	var err error
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}

func reason() {
	var err error
	fmt.Printf("Type of value stored inside the interface: %T\n", err)

	if _, err = fail(); err != nil {
		fmt.Printf("Type of value stored inside the interface: %T\n", err)
	}

	log.Println("No Error")
}

// so why did this code fail when we returned nil for the error?? Its because the fail
// function is using the concrete error type and not the error interface. In this case, there
// is a nil pointer of type customError stored inside the err variable. That is not the same
// as a nll interface value of type error
