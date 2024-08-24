package main

import(
	"fmt"
	"math" 
	_"log"
	_"os"
	_"os/exec"
)

// ... for quick debugging
var p = fmt.Println

	// Go supports constants of character, string, boolean, and numeric
	// values
	
	// const declares a constant value.
	const s string = "constant" 

func main(){
	
	p(s) 
	
	// a const statement can appear anywhere a var statement can.
	const n = 500000000
	
	// constant expressions perform arithmetic with arbitrary precision
	const d = 3e20/n
	p(d) 
	
	// a numeric constant has no type until it's given one, such as 
	// by an explicit conversion
	p(int64(d)) 
	
	// a number can be given a type by using it in a context that requires one, 
	// such as a variable assignment or function call. for example, here 
	// math.sin expects a float64
	p(math.Sin(n)) 
	
}
