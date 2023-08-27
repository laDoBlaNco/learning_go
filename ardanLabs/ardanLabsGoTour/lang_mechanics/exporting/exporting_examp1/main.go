package main

import (
	"fmt"

	"github.com/ladoblanco/ardanLabsGoTour/exporting/examples" 
)

/*
	Exporting provides us with the ability to declare if an identifier is accessible to code
	outside of the package its declared in. Its just that simple

	EXPORTING MECHANICS
	A pacakge is the basic unit of compiled code in  the Go language. It represents a physical
	compiled unit of code, usually as a compiled libary on the host operating system.
	Exporting determines access to indentifiers across package boundaries.

	For example we created a file on the sode of this one in a folder called counters
	with some notes on capital and lowercase letters for exporting. The above import is mine

*/

func main() {

	// Let's create a var of the exported type and initialize it, showing that we do have access to
	// its since we exported it.
	counter := examples.AlertCounter(10)
	fmt.Printf("Counter: %d\n", counter)
	
	
}

