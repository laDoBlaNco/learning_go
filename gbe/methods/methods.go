package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go supports methods (associated funcs) defined on struct types

type rect struct {
	width, height int
}

// this area method has a RECEIVER type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here's an
// example of the value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct
	p("area:", r.area())
	p("permieter", r.perim())

	// Go automatically handles conversion between values and pointers for method
	// calls. You may want to use a pointer receiver type to avoid copying on every
	// method call or to allow the method to mutate the receiving struct
	rp := &r
	p("area:", rp.area())
	p("perimeter:", (*rp).perim())

}

// Pretty simple, next let's look at Go's mechanism for grouping and naming related
// sets of methods: interfaces ("job descriptions" and contracts)
