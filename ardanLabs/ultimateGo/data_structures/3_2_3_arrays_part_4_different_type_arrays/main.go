package main

import "fmt"

/*
An intesting part of go is when we work with arrays of the same time but
differnet lengths
*/
func main() {
	var five [5]int
	four := [4]int{10, 20, 30, 40}

	fmt.Println(four, five)
	// five = four

	six := [6]string{"annie", "betty", "charley", "doug", "Bill", "Kevin"}
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p]  IndexAddr[%p]\n", v, &v, &six[i])
	}
}

// Its important be be clear on what the compiler is saying here. Its saying that
// an array of 4 integers and an array 5 integers represent data of different types
// The size of an arry is part of its type information. In Go, the size of an array
// has to be known at compile time.

// Finally let's look at how an array does in fact provide contiguous layout of memory
