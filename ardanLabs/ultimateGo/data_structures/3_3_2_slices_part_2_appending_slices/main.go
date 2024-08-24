package main

import "fmt"

/*
	Now let's start looking at some other things about this data structure as it is the most
	import data structure in Go.

	What makes the reference types interesting is that all reference types have a zero default
	of 'nil', just like a pointer set to nothing, Now with a string, its said to be empty not
	nil. This is why we say a string is a 'built-in' type and not a reference type.
	The zero default reference type [nil][0][0] for example has nil for its pointer word, but an
	empty slice [*][0][0] does have a pointer.
	Another difference is that we use our pointer semantics when reading and writing to reference types
	even though we move them with value semantics. Remember that strings are immutable, so we
	are creating new strings and changing the pointer, not using the pointer to read and write
	That's the difference.

	Slices represents collections. This means it can possibly NOT EXISTS (nil) or it might just be
	EMPTY but it still exists. This is one of the reasons why we use 'var', cuz it gives us the
	consistency.

	So where does the empty pointer point to?  struct{} 'the empty struct' (the literal struct{}{})
	Its a global variable, zero allocation type embedded into the runtime. It has a location and
	even with 100,000 empty structs, its all in the same location and this gives us the empty
	semantic.

	So where we are right now if we are with a nil pointer it means we have no backing array yet.
	so we then append to it.

*/

func main() {
	// declare a nil slice of strings (nil being the zero default for a slice which makes sense
	// since its a reference type)
	var data []string // 3 word structure [nil] [0] [0]

	// Capture the capacity of the slice.
	// lastCap := cap(data)

	// Append ~100k strings to the slice
	for record := 1; record <= 1e5; record++ {

		// the built-in function APPEND to add to the slice, using the index as an ID
		value := fmt.Sprintf("Rec: %d", record) // Sprint returns a strings
		data = append(data, value)              // her we add the string to the slice
		// 1. Start with a nil slice and make a call to 'append'

		// 2. Append takes its own copy of our 3 word slice val to work with. It then checks
		//    to see if the length and cap are the same. If they are then it needs to create
		//    some. It go out and create its first backing array, set it to 1 with capacity to 1
		//		- How does Go know how much to grow the capacity by? If the backing array is under
		// 		  1000 elements then it'll double the size. If its over 1000 elements then it will
		// 		  grow by 25%. That simple. What this is actually doing is creating efficiency since
		//        getting up to a certain point we'll start to have maybe 10 values but capacity for 16
		//        or something. so the further appends won't need to create new backing arrays and the
		//        GC won't need to come clean anything up.
		// 		- When we append to a slice that DOES have a backing array we do the same process
		// 		  but this time we need to copy everything from the old backing array to the new one
		// 		  before we move on to the next step.

		// 3. It then returns a new version of the slice value and this in turn takes the old and
		// 	  resets the pointer to the new back array and then updates its len/cap

		// 4. Once this is complete, the data in the append function is ready to be wiped on the
		//    next function call.

		// This called a Value Semantic Mutation and is the safest way to do mutation. Really it
		// reminds me of functional programming. Everyone works on their own copies and returns
		// new values rather than sharing data and mutating things real-time or inline. Everything
		// uses isolation. But we can only use the value semantics when it calls for it, meaning
		// moving around these built-in types. But later when we read or write to reference types
		// we have to use our pointer semantics.

		// NOTE: one thing we do need to worry about here is 'memory leaks'. How does this work in Go?
		// In Go this means that there's a value on the heap and it has a reference to it that it shouldn't
		// have. With that reference then the GC doesn't wipe it. We can detect this with a GC trace.
		// If we find there are then we need to look at 4 things:
		// 1. Goroutines - that are left blocking so they have references to them and just sit there
		// 2. Maps - when using maps as caches and not cleaning up or deleting keys
		// 3. Append - if we don't replace our variables with append results (bill = append(data,value)) then
		//    in reality we still have references to data that we don't need to have
		// 4. Any API with a 'close' function that we forgot to use.

		// NOTE that this was a contrived example. If in reality we know that we are doing 100k
		// appends, then we should just start our slice with a capacity of 100k and save on the
		// electrons of doing all those appends. So instead of using var zero default to nil
		// we just use 'make'

		// NOTE another good bug to look out for is if we know we might need to use append we
		// be can't set our length in make because make uses that length to tell it where to start
		// appending. if we do this (make([]string,1e5)) then now when do append, it'll start appending
		// at index 100,000, even if everything else is empty.
		
		// append is a great example of a value semantic mutation. Append has been built around 
		// the value semantics since We use value semantics to move all of these slices around our
		// programs.

	}

	inspectSlice(data)

}

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
