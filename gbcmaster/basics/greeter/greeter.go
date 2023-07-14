package main

import (
	"fmt"
	"os"
)

// os package allows you to talk to your operating system. its the os that controls
// all the input peripherals and Go's OS package allows us to talk to it.

// Intro to slices - in the os package we ahve the Args variable which is a slice
// Args []string - a alice a strings or a series of string values
// You get the arg values with Args[0] Args[1] Args[2] etc
// Args[0] is the path to the running programming
// Args[1] stores the first arg given to the program
// Args[2] stores the second arg and so on...

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path;", os.Args[0])
	fmt.Println("1st argument;", os.Args[1])
	fmt.Println("2nd argument;", os.Args[2])
	fmt.Println("3rd argument;", os.Args[3])

	fmt.Println("Number of items inside of os.Args:", len(os.Args))
}
