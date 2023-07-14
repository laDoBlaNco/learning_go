package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

func main() {
	age := 2
	fmt.Println(7.5 + float64(age)) 
}

// Not bad. The solution was convert 'age' and I converted the actual number 2
// which in the end wouldn't make much sense converting the number directly.
