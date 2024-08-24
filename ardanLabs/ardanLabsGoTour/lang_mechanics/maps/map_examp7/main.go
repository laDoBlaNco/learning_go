package main

import "fmt"

func main(){
	
	// here's an example to show how maps are reference types
	
	scores := map[string]int{
		"anna":21,
		"jacob":12,
	}
	
	// pass the map to a function for some mutation
	// remember the map is passed by value, but its a reference type.
	// so reading and writing is done by pointer
	double(scores,"anna") 
	
	// see the change in our map
	fmt.Println("Score:",scores["anna"])
	fmt.Println(scores) 
}

// double finds the score for a player and multiplies it by 2
func double(scores map[string]int,player string){
	scores[player] *=  2
}
// again nothing novel here

