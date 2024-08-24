package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Functions are central to Go. We'll learn about functions with a few
// different examples.

// First, here's a function that takes two ints and returns their
// sum as an int. Note: I put this in the main func and got an error because
// you can't declare a normal func inside another. Only an anony func which we'll
// get into later for sure.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won't automatically return the value
	// of the last expression
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit
// the type name for the like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// function calls are as expected 'name(args)'
	res := plus(1, 2)
	p("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	p("1 + 2 + 3 =", res)

}

// There ar several other features to Go functions. One is multiple return values
// which we've seen already, which we'll look at next.
