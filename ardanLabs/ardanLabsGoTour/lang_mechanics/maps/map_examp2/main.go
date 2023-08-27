package main 

import "fmt"

// Sample program to show how maps behave when you read an absent key.

func main(){

	// first let create a map to track scores for players in a game again.
	scores := make(map[string]int) 
	
	// Read the element at key "anna". It is absent so we get the zero default for this
	// map's value type
	score := scores["anna"] 
	fmt.Println("Score:",score) 
	
	// If we need to check for the presence of a key we use the 2 var assignment. The 2nd var
	// is a bool
	score,ok := scores["anna"]
	fmt.Println("Score:",score,"Present:",ok) 
	
	// We can leverage the zero default behavior to write convenient code like this
	scores["anna"]++
	
	// Without this behavior we would ahve to code in a defensive way like this
	if n,ok:=scores["anna"];ok{
		scores["anna"]=n+1
	}else{
		scores["anna"] = 1
	}
	
	score,ok = scores["anna"]
	fmt.Println("Score:",score,"Present:",ok) 
	
	
	
}
