package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

func main() {

	path := os.Args[0]

	fmt.Println(path)
}

// The only difference was that I didn't build to  -o myprogram on my first version
