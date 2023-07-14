package main

import "fmt"

func main() {
	//This is the normal way to a var
	var speed int

	fmt.Println(speed)
}

/*
In this exercise I learned that the order of the var declaration and
using it matters. Just like in JS, Golang runs from top to bottom. So you
can't use a var before declaring it. Variable hoisting.
*/
