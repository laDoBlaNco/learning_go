// Go By Example: Command-Line Arguments
// command line args are a common way to parameterize execution of programs. For
// example, fo run hello.go uses 'run' and 'hello.go' arguments to the go program.

package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args provides eash access to raw command-line args. Note that the first value
	// in this slice is the path to the program, and 'os.Args[1:]' holds the args
	// to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// you can get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}

// to experiment with the command-line args it's best to build a binary with go build first
// but you don't have to, as I've done this before with run as well. You just get the 'run .'
// included in the []args

// Next we'll look at some more complex commandline processing with flags.
