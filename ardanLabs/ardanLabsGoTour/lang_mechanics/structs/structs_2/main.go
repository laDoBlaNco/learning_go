// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and initialize anonymous
// struct types.

// Now in this example we are looking at anony structs. These are similar to anony funcs,
// an unnamed literal type set to its non-zero value first. Its an anonymous  literal because
// we assign it directly to the variable with var on creation of the actual struct. 

package main

import "fmt"

func main() {

	// Declare a variable of an anonymous type set
	// to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init
	// using a struct literal. NOTE for anony structs with initializing values we use
	// struct{}{} syntax. The first {} has the structure and the second {} has the initialization
	// values. 
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
// The idea of a literal construction is just that, to construct something without a name. Again we
// should use 'var' for zero value and := with {} or {}{} for non-zero value construction.




