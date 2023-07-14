package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

//Go supports anonymouse funcs which can form closures. Anony funcs (as I call them)
// are useful when you want to define a func inline without having to name it.

// This func intSeq returns another function, which we define anonymously in the
// body of intSeq. the returned func closes over the variable i to form a CLOSURE
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call intSeq, assigning the result (a func) to nextInt. This func
	// value captures its own i value, which will be updated each time we
	// call nextInt.
	nextInt := intSeq()

	// Calling it a few times we see the effect of the closure o sea the closed over
	// universe of the variable i.
	p(nextInt())
	p(nextInt())
	p(nextInt())
	p(nextInt())
	p(nextInt())
	p()

	// To confirm that the state is unique to that particular func, create and
	// and test a new one
	newInts := intSeq()
	p(newInts())
	p(newInts())
	p(newInts())
}

// The last feature well look at here is recursion.
