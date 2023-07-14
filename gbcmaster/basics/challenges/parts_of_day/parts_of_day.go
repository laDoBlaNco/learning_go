package main

import (
	"fmt"
	"time"
)

// ... for quick debugging
var p = fmt.Println

func main() {

	switch now := time.Now().Hour(); {
	case now < 12:
		fmt.Println("Good Morning!")
	case now < 18:
		fmt.Println("Good Afternoon")
	case now < 20:
		fmt.Println("Good Evening")
	case now >= 20:
		fmt.Println("Good Night")
	}
	// Inanc's solution was slightly different. He used >= and came down instead
	// of going up. And he considered anything from midnight to 6 am as good night
	// rather than what I did considering that as morning.

}
