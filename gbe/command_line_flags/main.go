package main

import (
	"flag"
	"fmt"
)

// Go by Example; Command line flags
// command line flags are a ommon way to specify options for cli programs. For exmaple, in
// wc -l the -l is a command flag.

// Go provides a flag package supporting basic command-line flag parsing. We'll
// use this package to implement our example command-line program.

func main() {

	// Basic flag declarations are available for string, int, and bool options. Here
	// we declare a string flag 'word' with a default value 'foo' and a short description.
	// This flag.String function returns a string pointer (not a string value); we'll
	// how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// This declares 'numb' and 'fork' flags, using a similar approach to the word flag.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Its also possible to declare an option that uses an existing var declared
	// elsewhere in the program. NOTE that we need to pass in a pointer to the flag
	// declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and any trailing positional args. NOTE that we
	// need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	// That's why the 'flag.Functions' return pointers to strings.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// To test this one we'll going to build it first.
