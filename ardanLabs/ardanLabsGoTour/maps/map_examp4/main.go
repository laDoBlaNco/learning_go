package main

import "fmt"

// Sample program to show how to declare, initialize and iterate over a map.
// Also reiterates that a map is random

// user represents someone using the program
type user struct{
	name string
	surname string
}

func main(){
	
	// Declare an initialize the map literal with values
	users := map[string]user{
		"Roy":{"Rob","Roy"},
		"Ford":{"Henry","Ford"},
		"Mouse":{"Mickey","Mouse"},
		"Jackson":{"Michael","Jackson"},
	}
	
	// Now let's iterate over the map print each key and value
	for k,v := range users{ // this is the value semantic version of course
		fmt.Println(k,v)
	}
	
	fmt.Println()
	
	// Iterate over the map printing just the keys.
	// NOTE the results are different (random order)
	for key := range users{ // this is the pointer semantic version, but we are just retrieving data
		fmt.Println(key) 
	}
}
