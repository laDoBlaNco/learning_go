package main

import "fmt"

/*
	SLICE OF INTERFACE

	So when I declare a slice of an interface type, I'm capable of grouping different concrete
	values together in a slice. Because we are basing that grouping on what they can do and now
	what they are. This is why Go doesn't need the concept of sub-typing. Its not about a
	common DNA, its about a common behavior.

*/

// Sample program to show how the concrete value assigned to the interface is what is stored
// inside the interface.
// Our printer displays information
type printer interface {
	print()
}

// cannon defines a cannon printer
type cannon struct {
	name string
}

// print displays the printer's name
func (c cannon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

// espon defines an epson printer
type epson struct {
	name string
}

// print displays the printer's name.
func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}

func main(){
	
	// Create a cannon and epson printer
	c := cannon{"PIXMA TR4520"} 
	e := epson{"WorkForce Pro WF-3720"} 
	
	// add the printers to the collection using both value and pointer semanatics
	printers := []printer{
		// Store a copy of the cannon printer value
		c,
		// Store a copy of the epson printer value's address
		&e,
	}
	
	// change the name field for both printers
	c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100" 
	
	// Iterate over the slice of printers and call print against the copied interface
	// value
	for _,p := range printers{
		p.print() 
	}
	
	// When we store a value, the interface value has its own copy of the value. Changes
	// to the original value will not be seen. When we store a pointer, the interface value 
	// has its own copy of the address. Changes to the original value will be seen.
}
/*
	In this code we also see how a slice of interface type printer allows me to create a 
	collection of different concrete printer types. Iterating over the collection and leveraging
	polymorphism since the call to p.print changes its behavior depending on the concrete
	value the code is operating against. 
	
	The example also shows how the choice of data semantics changes the behavior of the 
	program. When storing the data using value semantics, the change to the original value
	is not seen. This is because a copy is stored inside the interface. When pointer semantics
	are used, any changes to the original value are seen.
	
	
	NOTES:
		- The method set for a value, only includes methods implemented with a value recvr
		- The method set for a pointer, includes methods implemented with both pointer and
		  value recvrs
		- Methods declared with a pointer recvr, only implement the interface with pointer
		  values.
		- Methods declared with value recvr, implement the interface with both a value and
		  pointer recvr.
		- The rules of method sets apply to interface types
		- Interfaces are reference types, don't share with a pointer
		- This is how we create polymorphic behavior in Go.
		
		
	QUOTES:
		"Polymorphism means that you write a certain program and it behaves differently depending
		on the data that it operates on." - Tom Kurtz (inventor of BASIC) 
		
		"The empty interface says nothing" - Rob Pike
		
		"Design is the art of arranging code to work today, and be changeable forever." - Sandi Metz
		
		"A proper abstraction decouples the code so that every change doesn't echo throughout the
		entire code base." - Roma Steinburg
	
*/
