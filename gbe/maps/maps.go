package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Maps are Go's built-in associative data type (hashes or dicts in other langs)
func main() {

	// To ceate a map we can use the same builtin 'make' as we do with slices
	// and arrays.
	m := make(map[string]int)

	// set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map will show all of its key/value pairs
	p("map:", m)
	p()

	// Get a value for a key with name[key]
	v1 := m["k1"]
	p("v1:", v1)
	p()

	// and of course built-in len returns the number of key/value pairs when
	// called on maps
	p("map length:", len(m))
	p()

	// The builtin 'delete' removes key/value pairs from a map
	delete(m, "k2")
	p("map after key delete:", m)
	p()

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map. This can be used to
	// disambiguate between missing keys and keys with zero values like
	// 0 or "". Here we didn't need the value itself, so we ignored it
	// with the blank identifier _
	_, prs := m["k2"]
	p("present?:", prs)
	p()

	// You can also declare and initialize a new map in the same line with
	// the map literal syntax
	n := map[string]int{"foo": 1, "bar": 2}
	p("map literal:", n)

}

// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println
