package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.

// Here's a function that will take an arbitrary number of ints as args.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// within the function, the type of nums is equivalent to []int (a slice of ints)
	// We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	p(total)
	p()
}

func main() {

	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// if you already have multiple args in a slice, apply them to a variadic func
	// using func(slice...) like this:
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	// so the secret to variadic functions is  the ... either before the type name
	// or on the end of your slice or array.
}

// Another key aspect of functions in Go is their ability to form closures
// These are formed wither as callbacks or as anony funcs returned from other funcs
// Let's look at that next.
