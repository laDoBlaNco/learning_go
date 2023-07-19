package main

import "fmt"

// POINTERS
// Pointers hold the memory address of another variable

// & as ONE meaning
// The & symbol means: The address of the variable it points to

// Example: if a:= 25 and we set b := &a
// The value of b is the memory address of a
// And its type is a pointer to an int - or a pointer type *int

// * has TWO meanings
// 1: The * symbol next to a VARIABLE means: Get the value of the variable that this
// pointer is pointing to aka DEREFERENCING or INDIRECTING

// Example: if we set n := *b
// The value of n is 25. Its the value stored at the memory address that b is
// pointing to.

// 2: The * symbol next to TYPE means the variable  being declared is a pointer
// type and it points to an address holding the type followed by the *

// Example: var myName *string
// myName is a variable which holds the memory address
// of a string variable.

// WHY USE POINTERS?
// One use case is to alter the variables passed into functions.
// When we call a function that takes an argument, that argument is copied to the function:
func zero(x int) {
	x = 0
}
func zeroPtr(x *int) {
	*x = 0
}

// Another case is to reveal if the value has actually been set
// or if its just the default value
var isSeventeen bool

// default value for bool is 'false'
// default value for interfaces, slices, channels, maps, functions and POINTERS is 'nil'
var isSeventeenPtr *bool

func main() {
	a := 25        // implicit dec
	b := &a        // b is not a pointer to a
	fmt.Println(b) // nothing special about these memory address, we can print them out normal
	//c := &b        // a pointer to a pointer to an int -- *(*int)
	// Could also write it like this.
	var c **int = &b // stll works the same

	fmt.Println(*c) // note its the same address as a
	fmt.Printf("c is now a %T\n", c)

	fmt.Println("====================================================================")
	x := 5
	zero(x) // we work with a copy but aren't changing the origina var
	fmt.Println(x)

	zeroPtr(&x) // we work with the original value dereferenced from x and alter that value
	fmt.Println(x)
	fmt.Println("====================================================================")

	fmt.Println(isSeventeen)    // note the default value for bool is hard to tell if set or not
	fmt.Println(isSeventeenPtr) // the default for a pointer to a bool is <nil> not false. So never set

	input := false
	isSeventeenPtr = &input

	fmt.Println(isSeventeenPtr)  // here we get a memory address
	fmt.Println(*isSeventeenPtr) // here we get the dereference value when dereference, no nil values
}

// Pointers in a nutshell.
