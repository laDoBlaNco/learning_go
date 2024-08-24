package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// In Go, an arry is a numbered sequence of elements of a SPECIFIC length.
// In typical Go code, slices are much more common to see; arrays are useful
// in some special scenarios.
func main() {

	// Here we create an array 'a' that will hold EXACTY 5 ints. The
	// type of elements and length are both part of the array's type name
	// By default an array is zero-valued, which for ints means 0s
	var a [5]int
	p(a)
	p()

	// We can set a value at an index using the array[index] = value syntax,
	// and get a value with array[index].
	a[4] = 100
	p("set:", a)
	p("get:", a[4])
	p()

	// The builtin 'len' returns the length of an array
	p("len:", len(a))
	p()

	// Use this syntax to declare and initialize an array in one line: array literals
	b := [5]int{1, 2, 3, 4, 5}
	p("array literal:", b)
	p()

	// Array types are one-dimensional, but you can COMPOSE types to build
	// multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	p("2d:", twoD)
}

// Note that arrays appear in the form [v1 v2 v3 ...] when printed with
// fmt.Println
