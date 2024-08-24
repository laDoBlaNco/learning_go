package main

import(
	"fmt"
	_"log"
	_"os"
	_"os/exec"
)

// ... for quick debugging
var p = fmt.Println

func main(){
	
	// In go, variables are explicitly declared and used by the compiler to e.g.
	// check type-correctness of function calls.
	
	// var declares 1 or more variables
	var a = "initial"
	p(a) 
	
	// you can declare multiple variables at once.
	var b,c int = 1,2 
	p(b,c) 
	
	// Go will infer the type of initialized variables
	var d = true
	p(d) 
	
	// Variables declared without a corresponding initialization are zero-valued
	// For example, the zero value of an int is 0
	var e int
	p(e) 
	
	// the := syntax is shorthand for declaring and initializing a variable, e.g.
	// for var f string = "apple" in this case
	f := "apple"
	p(f) 
	
}
