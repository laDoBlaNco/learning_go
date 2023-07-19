// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how variables of an unnamed type can
// be assigned to variables of a named type, when they are
// identical.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	// flag2 bool // this would change the type of the struct or better yet it's shape.
	pi      float32
}

func main() {

	// Declare a variable of an anonymous type and init
	// using a struct literal and the same shape as our 'example' above.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Create a value of type example. So e is now a zero default instance
	// of our example type.
	var ex example
	fmt.Println(ex)

	// Assign the value of the unnamed struct type
	// to the named struct type value and this can be done cuz they are the same shape, same fields.
	ex = e

	// Display the values.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
