// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var slice []int 

	// Append numbers to the slice.
	for i:=0;i<10;i++{
		slice=append(slice,i)	
	}

	// Display each value in the slice.
	for _,v:=range slice{
		fmt.Println(v)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"kevin","xavier","brittney","kelen","odalis"}

	// Display each index position and slice value.
	for i,v:=range names{
		fmt.Printf("Index: %d\tName: %s\tAddress: %p\n",i,v,&v)
	}
	fmt.Println()

	// Take a slice of index 1 and 2 of the slice of strings.
	nameSlice := names[1:3]

	// Display each index position and slice values for the new slice.
	for i,v:=range nameSlice{
		fmt.Printf("Index: %d\tName: %s\tAddress: %p\n",i,v,&v)
	}
	
}
