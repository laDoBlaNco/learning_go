package main

/*
	So we are going to talk briefly again about the range mechanics because most Go devs
	don't ever realized that the value vs pointer semantics run so deep in go that there
	2 different version of for/range, one with value and one with pointer semantics.
	
	
*/

import "fmt"

func main(){
	
	// Using the value semantic form of for/range
	friends := []string{"Annie","Betty","Charley","Doug","Edward"}
	for _,v := range friends{ // value  semantic form with its key/val vars making copies on each iteration
		friends = friends[:2]
		fmt.Printf("v[%s]\n",v) 
	}
	
	// Using the pointer semantic form for the for/range
	friends = []string{"Annie","Betty","Charley","Doug","Edward"}
	for i := range friends{
		friends=friends[:2]
		fmt.Printf("v[%s]\n",friends[i]) 
		// this panics out because we are working on shared data and mutate it, which causes
		// me to get an out of index error. We are doing the same mutation in the loop above
		// but there we making copies, mutating in isolation, and passing values, so it doesn't
		// impact our result, but does cost effiency.		
	}
}
