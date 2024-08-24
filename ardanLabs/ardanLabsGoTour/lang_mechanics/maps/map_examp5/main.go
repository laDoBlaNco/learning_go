package main

import(
	"fmt"
	"sort"
)

// this is a sample program to show how we can walk through a map and make it sorted
// even though its designed to be random. If that's the functionality that we need.

// let's start with our user struct
type user struct{
	name string
	surname string
}

func main(){
	
	// get our map going
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}
	
	// let's first pull the keys from the map
	var keys []string // create a zero default slice of strings
	for key := range users{
		keys = append(keys,key) // append all the keys to our slice
	}
	
	// then we sort that slice with our sort pkg
	sort.Strings(keys) 
	
	// now we can use that slice to walk through the map and pull things out in alphabetical order
	for _,key := range keys{
		fmt.Println(key,users[key]) 
	}
	
	// Nothing novell here.
	
}
