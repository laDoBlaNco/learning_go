package main

// Declare and make a map of integer values with strings as keys
// Populate the map with five values and iterate over the map
// to display the key/value pairs.

import "fmt"

func main(){
	
	// declare and make a map of integer type values
	stuff := make(map[string]int) 
	
	// Initialize some data into the map
	stuff["one"]=1
	stuff["two"]=2
	stuff["three"]=3
	stuff["four"]=4
	stuff["five"]=5
	
	// Display each key/value pair
	for k,v := range stuff{
		fmt.Println(k,v) 
	}
}
