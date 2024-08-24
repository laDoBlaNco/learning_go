package main

import(
	"fmt"
)
// ...for quick debugging
var pl = fmt.Println

// Recursion is the act of a function calling itself. And anytime you use recursion
// you must must must think of the "base case" or the condition that ends the 
// process this function calling itself. 

// Let's do a factorial to show this.

func factorial(num uint64)uint64{
	// start with you base case
	if num==0{
		return 1
	}
	return num * factorial(num-1) // this ensures that it will eventually get to base case.
} 

func main(){

	pl("Factorial 4 =",factorial(4))
	pl("Factorial 30  =",factorial(30))

}
