package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// for is Go's only looping construct. Here are some basic types of for loops
func main() {

	// The most basic type with a single condition:
	i := 1
	for i <= 3 {
		p(i)
		i += 1
	}
	p()

	// A classic initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		p(j)
	}
	p()

	// for without a condition will loop repeatedly until you break out
	// of the loop or return from the enclosing function
	for {
		p("loop")
		break
	}
	p()

	// you can also continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		p(n)
	}

}

// We'll see some other for forms later when we look at range statements, channels
// and other data structures.
