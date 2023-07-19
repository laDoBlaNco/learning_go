// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Struct types are a way of creating complex types that group fields of data together. They are a great
// way of organizing and sharing different aspects of the data your program consumes. Similar to the 
// Dart record type also used to group a lot of different types under one variable. 

// A computer arch's potential performance is DETERMINED PREDOMINATELY BY ITS WORD LENGTH (the number of bits
// that can be processed per access) and, more importantly, memory size, or the  number of words that it can
// access. 


// First let's look at how we declare create and intialize them.
// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a concrete user-defined type as a composite of different fields and types. 
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// We declare a variable of type example set to its
	// zero value using the 'var' keyword.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1) // here we can see the value of this struct with its zero defaults

	// Declare a variable of type example and init using
	// a struct literal.
	// The difference here being that rather than creating with 'var' and leaving zero default, we create it
	// with its initial values using := and 'example{}'. Similar to the constructors in Dart.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values. NOTE that we can access each field with the . operator
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
