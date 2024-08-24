package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go supports recursive functions. Here's the classice example.

// This fact function calls itself until it reaches the base case of
// fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	p(fact(7))

	// closures can also be recursive, but this requires the closure
	// to be declared with a typed var explicitly before its defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// since fib was previously declared in main, Go knows which function
		// to call with fib here.
		return fib(n-1) + fib(n-2)
	}

	p(fib(7))
}

// I need to think about this one a bit. so when we say we need to declare the
// closure (anony func) with a type var before it defined that's doing this:
// var fib func(n int) int - so here we are using that func sig as a type. so this
// is still a closure(anony func) but we are giving it a FUNCTION TYPE and then
// we can declare it as any other anony and use it, since it was declared (named)
// recursively.

// So the bottom line is that if I ever need to do recursion from inside of a parent
// function or main for example, I need to declare a var of function type prev.
// is basically the way to give a name to a anony func. Similar to when you create
// a lambda function or an anony in other langs and passing them to a declared
// var.
