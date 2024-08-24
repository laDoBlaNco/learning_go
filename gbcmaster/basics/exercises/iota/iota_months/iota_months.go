// iota exercises -
package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// EXERCISE: iota months

// 1. Initialize the constants using iota
// 2. You should find the correct formula for iota

// RESTRICTIONS:
// 1. Remove the initializer values from all constants
// 2. Then use iota once for initializing one of the constants.

// EXPECTED OUTPUT:
// 9 10 11

func main() {

	const (
		Nov = -iota + 11 // here iota is -0 + 11 = 11
		Oct              // then -1 + 11 = 10
		Sep              // then -2 + 11 = 9
	)
	p(Sep, Oct, Nov)

}

// done deal!
