package main

import "fmt"

/**TAKING SLICES OF SLICES
 *
 * Now that we understand the mechanics of slices its time to understanding a few things about
 * how they got there names. Also some of the 'gotchas' because even though we move them around
 * with those value semantics, we read and write them with pointer semantics and that's where
 * things start to get tricky
 *
 * When we slice a slice we are using the same backing array for efficiency. This way we aren't
 * recopying of the same data. We use the syntax 'slice1[a:b]' 
 * which means set our [*index2]:[*index2+length of our new slice]. Then then changes our slice
 * stucture to [2][4][6]. The cap is 6 because capacity is counted from starting index to the end
 * of the backing array. Also NOTE that they memory addresses ARE THE SAME. We aren't dealing with 
 * new blocks of memory but looking through a window, so to speak, at our original backing array.
 *
 * Why is this cool and important? Because since we aren't creating more and more copies, we aren't
 * polluting our heap with a lot of allocations and thus we will have less GCs and in turn better
 * performance. 
 *
 * The trick part here is that remembering that slices are REFERENCE TYPES, when we need to read or
 * write from them, we are actually using POINTER semantics and thus there is the potential for 
 * side effects, or data being mutated that we didn't expect. In our example below we change the
 * data at index 0 in slice2 and that changes the data in the backing array. If that wasn't expected
 * we've just created a bug. 
 *
 * But what if we try to avoid this by using append? All it does is move it down two index as
 * the append mechanics only looks at your current slice to see if len and cap are the same. If
 * they aren't, then a new backing array won't be allocated, it'll simply overwrite what's in the
 * backing array at that next spot.
 * 
 * Now if we can make it so that call to append sees a slice where the len and cap are the same, then
 * this will trigger a new allocation which will give us what we  want, as the append mechanics say
 * "if you don't have space then MAKE A COPY of what you want to a new block of memory. 
 * We do this by using a 3rd argument to our slice request [start:start+length:capicity] 
 * 
 * 
 *
 */

func main() {

	// create a slice with a length of 5 elements and a capacity of 8
	slice1 := make([]string, 5, 8)  // [*slice1][5][8]
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	inspectSlice(slice1)
	fmt.Println("****************************************")

	// Now let's take a slice of slice1. We only want indexes 2 an 3. 
	// parameters are [starting_index:(starting_index + length)] that being length of 
	// the new slice we want
	slice2:=slice1[2:4:4] // so start_index = 2 and (starting_index + length) = 2 + 2 = :4
	// and now we add in the 3rd arg which is the capacity of our new slice which forces the algo
	// to make a copy into a new allocation and our CHANGED doesn't impact our backing array
	inspectSlice(slice2) 
	fmt.Println("****************************************")

	// change the value of the index 0 of slice 2
	// slice2[0]="CHANGED"
	slice2=append(slice2,"CHANGED")  
	
	// Display the change across all exisiting slices
	inspectSlice(slice1)
	inspectSlice(slice2) 
	fmt.Println("****************************************")
	// NOTE the new capacities and memory addresses with the use of that 3rd slice arg.
	
	
	
	
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
