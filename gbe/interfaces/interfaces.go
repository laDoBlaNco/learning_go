package main

import (
	"fmt"
	_ "log"
	"math"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Interfaces are named collections of method signatures. They are also thought of
// (with the arrival of generics) as named collections of type sets. So interfaces
// are said to define a type set or a method set. Or in other words a list of methods
// a list of types that implement the methods that are defined by the interface. This
// is also intimately related to Go generics.

// Here's a basic interface for a geometric shape
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on rect and circle types:
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To "implement" the interface in Go, we only need to implement all the methods
// in the interface. This guarantees that anyone with knowledge of the  interface
// or has built any funcs using the interface will be able to use our types
// Here we will implement geometry on rects (o sea create the methods needed)
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Here we do the same for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Now the key to using interfaces is in the variables that are of that interfaces
// type. Ifa variable has an interface type, then we can call methods that are in
// the named interface. Here's a generic measure function taking advantage of this
// wot work on any geometry as we know anything that fits the interface will have 
// the methods we need for our func.
func measure(g geometry){
	fmt.Printf("%#v\n",g)  
	p("The area:",g.area()) 
	p("The perimeter:",g.perim()) 
}

func main() {

	r:=rect{width:3,height:4}
	c:=circle{radius:5}
	
	// The circle and rect struct types both implement the geometry interface so
	// we can use instances of these structs as args for our measure func
	measure(r) 
	p() 
	measure(c) 
	
}
// here some more to learn about Go interfaces:
// https://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go
