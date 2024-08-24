package main

// Go By Example: Environment Variables
// Environment variables are a universal mechanism for conveying
// configuration information to Unix problems. let's look
// at how to set, get, and list environment vars.

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use os.Setenv. To get a value for a
	// key, use os.Getenv. This will return an empty string if the
	// key isnt' present in the environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
