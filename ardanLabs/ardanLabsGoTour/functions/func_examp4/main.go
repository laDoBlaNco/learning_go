package main

/*
	Now let's take a quick look at anony funcs. They have all the normal closure goodness
	that I've come to love from anonymous functions

*/

import(
	"fmt"
)

func main(){
	
	var n int
	
	// Declare an anony func and call it
	func(){ // note is just a func without a name
		fmt.Println("Direct:",n)
	}() // then as expected we call it with '()'
	
	// Declare an anony func and assign it to a variable
	f := func(){ 
		fmt.Println("Variable:",n)
	}
	
	// Now we need to call the anony func through the var 'f'
	f() 
	
	// Now let's defer a call to the anony func till after main returns, thus showing the
	// closure aspects
	defer func(){
		fmt.Println("Defer 1:",n) 
	}() 
	
	// set the value of n to 3 before the return
	n=3
	
	// call the anony func through the variable again
	f()
	
	// Defer the call to the anony func till after main returns again
	defer func(){
		fmt.Println("Defer 2:",n)
	}() 
	
}
// Result came out exactly as I thought it would. Main points to NOTE are that defer calls stack
// LIFO and the anony funcs still ahve access to the internal state of main, meaning that they can
// see our n var change from 0 to 3.

