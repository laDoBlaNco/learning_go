package main

import (
	"fmt"
)

// ...for quick debugging
var pl = fmt.Println

// closures are basically functions that have access to to variables in other scopes
// that aren't internal to its own. These are usually anonymous as they are returned
// without name from other functions.

// we can also pass functions to functions with closures/anonys
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer:", f(x, y))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {

	intSum := func(x, y int) int {
		return x + y
	}
	pl("5 + 4 =", intSum(5, 4))

	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}
	changeVar()
	pl(samp1)
	// this works because same as in a func that return a func. That returned func
	// has access to the vars in the parent func. here changeVar has access to the
	// vars in the greater func (main). Remember that we can't normally create a
	// real declared func within main. But doing an anony or closure, we are able to
	// have access to this samp1 var without the use of a declared pointer or
	// passing by value.
	
	pl()
	pl("Using a function taking a function:") 
	useFunc(sumVals,5,8) 

}
