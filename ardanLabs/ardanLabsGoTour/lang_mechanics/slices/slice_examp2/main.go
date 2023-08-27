package main

import "fmt"

/*
	DATA SEMANTIC GUIDELINE FOR SLICES:
	As a guideline, if the data, remembering that data rules everything for us in Go, if the
	data we are workign on is a slice, then we use VALUE SEMANTICS to move the data around
	our program. This includes declaring fields in a type and all other references types like:
		- slices
		- maps
		- channels
		- interfaces
		- functions -this comes up constantly with functions in two ways really. First because
		  functions are in fact reference types, but also when we talk about moving data, its
		  typically with functions that this happens. Moving program boundaries as data gets
		  passed into functions (and methods).

		  func Foo(data []byte) []byte {...}  Here we move slice in and out of the func by
		  value and not with pointer semantics.

	One valid reason to switch to pointer semantics would be if we need to share the slice for a
	decoding or unmarshaling operation. We see this in the stdlib a lot. Using pointers for these
	types of operations are ok, but it needs to be documented if its not obvious what is going on.

*/

type Foo struct {
	x []int // this means that we are passing slices into our struct by value and not pointer
	y []string
	z []bool
}

func main() {

	/*
		CONTIGUOUS MEMORY LAYOUT - The idea behind the slice is that we get the most efficient data structure
		as it relates to the machine, which is the array, but ast the same time we have the ability to be
		dynamic and efficient with the amount of data we need at runtime and future growth. That efficiency
		comes in the way appending, which we'll talk about in a bit.
	*/

	// Here let's create a slice with 'make' with a length of 5 and a capacity of 8. Remembering that
	// the length is the amount of elements we have access to while the capacity is the number of
	// total elements tha the backing array has. This is important when we talk about appending and
	// growth efficiencies a little later.
	fruits := make([]string, 5, 8)
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "banana"
	fruits[3] = "grape"
	fruits[4] = "plum"

	inspectSlice(fruits)

}

// inspectSlice is a func that will expose the slice header for review and I'll use for other examples
// in this lesson - notice how the for/range only iterates over the length of the slice and not the 
// capacity.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d]  Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d]  %p  %s\n", i, &slice[i], s)
	}

}
