package main

import "fmt"

// here's an example showing that we can't take an address of a an element in a map
// as I have in my notes, because Elements in a map are not addressable. They have to
// be pulled and assigned to a local var in memory first

type player struct {
	name  string
	score int
}

func main() {

	// declare our map with initial vals
	players := map[string]player{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 21},
	}

	// let's try to take an address of an element
	// anna := &players["anna"]
	// anna.score++
	// map_examp6/main.go:21:11: invalid operation: cannot take address of players["anna"] (map index expression of type player)

	// what we can do is take out the value, assign it to a var locally, work with it, and then
	// put it back
	player := players["anna"]
	player.score++
	players["anna"] = player 

	fmt.Println(players)

}
