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
	
	// Go has various value  types including strings, integers, floats, booleans
	// etc. here are some basic examples.
	
	// strings, which can be added together with '+'
	p("go" + "lang") 
	
	// integers and floats
	p("1+1 =",1+1) 
	p("7.0/3.0 =",7.0/3.0) 
	
	// Booleans, with boolean operators as you'd expect
	p(true && false)
	p(true || false)
	p(!true) 	
	
}
