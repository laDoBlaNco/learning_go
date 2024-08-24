package main

import "fmt"

// First let's declare a struct that reps a ball player, including fields for
// name, atBats and hits
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a  method that calculates the batting average for a player
func (p player) average() float64 {
	if p.atBats == 0{
		return 0.0 
	}
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Let's create a slice of players and populate each player with field values
	players := []player{
		{"Kevin", 70, 50},
		{"Odalis", 50, 20},
		{"Kelen", 100, 30},
		{"Xavier", 100, 90},
	}

	// Display the batting average for each player in the slice
	for _, p := range players {
		fmt.Printf("%s has a batting average of %.3f\n", p.name, p.average()*1000)
	}
}

// In the end there s question as to why we use the for range semantic that we did. I didn't
// use the same one that Bill used. I'm not sure why he uses pointer semantics in this case
// since we are just pulling information, not mutating anything, and the fields of the type
// used value semantics as well. I used value semantics in both the method as well as the 
// for/range
