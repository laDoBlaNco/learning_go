package main

import "fmt"

// This example mainly reiterates the fact that for/range has a value semantic version as well 
// as a pointer semantic version. The deciding factor for us should be what type of data 
// is in the collection that we are ranging over.

func main(){
	
	// using the value semantic form of the for/range
	// Even though we alter the original slice, it doesn't impact us cuz we are working on
	// a copy on each iteration
	friends := []string{"Annie","Betty","Charley","Doug","Edward"}
	for _,v := range friends{
		friends = friends[:2]
		fmt.Printf("v[%s]\n",v) 
	}
	
	// using the pointer semantic version of for/range
	// this version panics out since we alter the original slice during the ranging
	// and since its shared, it impacts what we are printing out.
	friends = []string{"Annie","Betty","Charley","Doug","Edward"}
	for i := range friends{
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i]) 
	}
}
// Another thing to NOTE is the way we are accessing the element in each iteration. When
// working with the value semantic version Go gives us the index of a copy of the slice
// as well as a copy of the actual element to work with in a var v, so we just use v.
// In the pointer semantic version, we only get an index, meaning we need to use it to 
// go into the original slice to access the values from there.
