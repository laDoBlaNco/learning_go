package main

import "fmt"

/*
Now with slices, this is the most important data structure in Go. We get both the mechanical sympathy with
Arrays and the dynamics of vector.

Remember that there are 3 types in go:
	- Built-in types
	- Struct types
	- Reference types (maps, slices, functions)



*/

func main() {
	// make is a built-in function that's purpose is to allocate memory for the backing array.
	// make also works for maps and channels, which kinda tells me already how those are implemented.

	// Create a slice with a length of 5 elements
	fruits := make([]string, 5, 8) // note that there is no value in [] that tells us this is a slice type
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// as far as memory representation we have a slice (3 word structure (24 bytes))
	//		- The first word being the POINTER to the backing array (set to its 0 value).
	// 		- The second word is the LENGTH of the backing array, which means everything we can access
	// 		  through our slice (window into the array) Even if the backing array is 100, we can only access
	// 		  5 indexes of it through this slice.
	// 		- The third word is the CAPACITY. This is the full amount of indexes in the backing array, whether
	// 		  we can access them or not. I could also think that this is the 'length' of the backing array
	// 		  while 'length' above is the length of the slice. Also good to note that capacity is really
	// 		  for efficiency in future  growth.

	// You can't access an indexof a slice beyond its length
	// fruits[5] = "Runtime Error" // panic: runtime error: index out of range [5] with length 5

	fmt.Println(fruits) // and to put this in the print function (across) the program boundary we use
	// value semantics to move it around. So even though this is a reference type, we still use value
	// semantics to move between program boundaries. So Println gets a copy of that 3 word structure.
	// NOTE: that when reading and writing we are using pointer semantics. So again we could have 5 copies
	// of our 3 word structure but they all point to the same memory.

	// Above I added 8 as a 3rd arg to 'make' and that's the CAPACITY. so our backing array. so we still
	// only have a slice of length 5, but we now have an backing array of length/capacity 8.

	inspectSlice(fruits)
}

// inspectSlice exposes the slice headerr for review showing us that the backing array is a contiguous
// block of a memory on a predictable stride.
func inspectSlice(slice []string) { // the arg its asks for a a slice VALUE not a pointer to a slice
	fmt.Printf("Length[%d]  Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { // and here we use the VALUE semantic version of the for/range
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
	// How do we know the form of the for/range to use:
	// What is this a collection of? Slice of strings.
	// What is a string? A built-in type
	// What do we use for built-in types? Value semantics
}

// when we don't have this type of data semantic consistency we are creating bugs.
