package main

import "fmt"

/*
	Key Map Restrictions:
	Not all types can be used as a key
	
	A slice is a good example of a type that can't be used as a key. Only values
	that can be run through the hash function are eligible. A good way to recognize
	types that can be a key is if the type can be used in a comparison operation. You
	can't compare two slice values, so they can't be keys.

*/

// User represents someone using the program
type user struct{
	name string
	surname string
}

// users defines a set of users
type users []user

func main(){

	// Declare and make a map that uses a slice as a key
	u := make(map[users]int) 
	
	// map_examp3/main.go:23:16: invalid map key type users
	
	// Iterate over the map
	for k,v := range u{
		fmt.Println(k,v) 
	}
	
	
}

// Notes:
// 	- Maps provide a way to store and retrieve key/value pairs
// 	- Reading and absent key returns the zero default value for the map's value type
// 	- Iterating over a map is always random
// 	- The map key must be a value that is comparable
// 	- Elements in a map are not addressable 
// 	- Maps are a reference type

