package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Branching with if and else in Go is straight-forward
func main() {
	// Here's a basic example:
	if 7%2 == 0 {
		p("7 is even")
	} else {
		p("7 is odd")
	}

	// You can have an if statement without an else
	if 8%4 == 0 {
		p("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables declared in
	// this statement are then available in the current and all subsequent
	// branches
	if num := 9; num < 0 {
		p(num, "is negative")
	} else if num < 10 {
		p(num, "has 1 digit")
	} else {
		p(num, "has multiple digits")
	}

}

// Note: you don't need ()s around conditionals in Go, but {}s are required.
// There is no ternary operator in Go, so you'll need to use a full 'if' statement
// even for basic conditions, but I think it'll allow you to do one liners most
// times without reformatting you, and even if gofmt does reformat, if won't impact
// your productivity but rather help it, so no complaints from me on that one.
