package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go supports pointers allowing you to pass REFERENCES to values and records
// within you program. Basically allowing in-line changes rather than passing
// a value and returning a new value, etc.

// We'll see how pointers work in contrast to values with 2 functions: zeroval and
// zeroptr. zeroval has an int param, so args will be passed to it by value.
// zeroval will get a copy of the ival distinct from the one in the calling func.
func zeroval(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int (pointer type) param, meaning that it takes an
// int pointer. the *iptr code in the function body then DEREFERENCES the pointer
// from its memory address to the current value at the address. Assigning a value
// to a dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	p("initial value:", i)
	p()
	zeroval(i)
	p("After calling zeroval:", i)
	p()
	// The &i syntax gives the memory adderss of i, i.e. a pointer to i.
	zeroptr(&i)
	p("After calling zeroptr:", i) // note here we are still using normal i as the pointer is
	// referencing it.
	p()
	// pointers can be printed too
	p("The pointer:", &i)
	fmt.Printf("Using the pointer verb: %p\n", &i)

}
// zeroval doesn't change the i in main, but zeroptr does because it has a reference
// (pointer) to the memory address for that variable.
