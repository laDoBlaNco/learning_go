package main

import "fmt"

// variables are really at the heart of the Go language, but we aren't going to focus on what
// a var is but more what key things in Go we need to know. for example with Go vars the important
// piece is the type.

// Types - we have to understand types. It means two things. The amount of memory we need for reading and
// writing and what that memory represents. With these 2 things there is no integrity. A byte is just a byte
// until we know what type it is. A variable as we know is a name that identifies that memory allocated

// There are 3 classes of types in Go:
// 	- builtin types
// 		- Types provide integrity and readability by answering 2 questions:
// 			- What is the amount of memory to allocate? (1,2,4,8 bytes)
// 			- What does that memory represent? (int,uint,bool,....)
//  - reference types
//  - custom user-defined struct types

func main() {

	// Declare variables that are set to their zero value.
	var a int // this is interesting cuz we can be more precise (int8, int16, int32, int64) but only if needed
	// this var tells us we are representing integers but it doesn't tell us size.Also remember that the main
	// thing we are doing when coding is reading and writing to memory and that's integer based. So remembering
	// the machine, the architecture, we want to use the most efficient integers for our architecture. So by using
	// int, we let the compiler choose the precision depending on the architecture.
	var b string // strings are intresting in Go because of their implemenation. In go they are a 2 'word'
	// value. A word is a generic allocation. In the case of the 'go playground' its 4 bytes, in the case of
	// an 64bit machine its 8 bytes for an integer or an address, and in the case of a string its 2 words and
	// the first word in this 2 word value is a pointer to nil and the number of bytes.
	var c float64 // this tells us both the size (memory) and what it reps
	var d bool    // here we know what its used for, but the memory allocation is just 1 byte of the 8

	// This stuff about the word size is important because its also at the heart of understanding memory
	// allocation and references, etc. The word size represents the amount of memory allocation required
	// to store integers and pointers for a given arch.
	// 		- 32 bit arch: word size is 4 bytes of memory allocation
	// 		- 64 bit arch: word size is 8 bytes of memory allocation
	// This is key because in Go we have internal data structures (maps, channels, slices, interfaces, and
	// functions) that store integers and pointers. Everything is integers and pointers. The size of these
	// data structures will be based on the arch being used to build the program. So the amount of memory
	// used to build an int, a pointer, a word, etc will always be the same on the same arch.

	// part of the integrity of Go is always setting its zero default, if the engineer doesn't do so. One of
	// the warts of Go is that there are many ways to create a var, so the team needs  to use conventions to
	// deal with that. For example above, we use 'var' when its a zero default declaration and below we use
	// the short declaration operator := when are initializing.
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	// A string is a 2 word internal data structure in Go. The 2 word value above which was nil over 0 will now be
	// a pointer to where that "hello" starts in memory over the number of bytes 5. This is key as we'll
	// see it again with arrays and slices. So a string is 2 words (a pointer / length or number of bytes)
	cc := 3.14159
	dd := true
	fmt.Println(aa, bb, cc, dd)

	// The other item we need to discuss is that Go favors conversion over casting. With casting we assume
	// there's space to expand values and with that we get bugs in our codebase by overwriting. Go would
	// rather Convert the value to a new allocation rather than guess or assume that there's space
	// Specify type and perform a conversion:
	aaa := int32(10) // so here we get the new allocation rather than just the casting other langs do
	// so we change this 10 to a 4byte allocation from its 1 byte.
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
	// Go hasn't eliminated casting as you can still get to it in the 'unsafe' pkg, but since integrity is more
	// important, conversion over casting

	/*
		RECAP:
		- Go has its types (builtin, reference, and struct)
		- There are way too many ways to assign in go so we will be using
			- var for zero default value allocations, that's Go setting all the bits to 0000000
			- := for everything else
		- Coversion over casting

		- TYPE IS LIFE and it tells us two pieces of information
			- size is the amount of memory in bytes that we'll be reading and writing
			- it's representation
		- This is where we get our compiler protection, if we know the type we can ensure integrity
			- Accurate, consistent, and efficient

	*/
}
