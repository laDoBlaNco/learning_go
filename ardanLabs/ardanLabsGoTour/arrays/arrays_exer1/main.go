// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

func main() {

	// Declare an array of 5 strings set to its zero value.
	var myStrings [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	myStrings2 := [5]string{"one","two","three","four","five"}

	// Assign the populated array to the array of zero values.
	myStrings = myStrings2 // since they are the same type, this works. The type being [5]string

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for _,v := range myStrings{
		fmt.Printf("Value:%s\tAddress:%p\n",v,&v) 
	}	
}
