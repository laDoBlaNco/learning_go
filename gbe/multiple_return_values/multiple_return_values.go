package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go has builtin support for multiple return values. This feature is used often
// in idiomatic Go, for example to rturn both result and error values from a
// function.

// The (int,int) in this function sig shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we have to use the 2 different return values from the call with
	// multiple assignment. This is basically using tuple assignment
	a, b := vals()
	p(a)
	p(b)
	p()

	// if you only want a subset of the returned vals, you the blank identifier
	// to ignore what you don't need
	_, c := vals()
	p(c)

}

// Accepting a variable number of args is another nice feature of Go functions.
// let's look at that now.
