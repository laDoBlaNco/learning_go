package main

import "fmt"

/*
APPENDING WITH SLICES - Efficiency in growth

The language provides a built-in function, much like 'make', called 'append' to add values to an
existing slice.
*/

func main() {

	var data []string // this creates a zero default slice, our 3 word struture [*nil][0][0]

	lastCap := cap(data)

	for record := 1; record <= 102400; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		if lastCap != cap(data) {
			// calc the % of change
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100
			// save the new values for capacity
			lastCap = cap(data)
			// display the results
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)

		}
	}

	// The append function works with a slice even when the slice is initialized to its zero default
	// state. The API design of append is what's interesting to us though because it uses a value
	// semantic mutation approach. Value semantic mutation is the safest mutation that you can use and
	// its really part of the basis of Functional Programming. Append gets its own copy of a slice value
	// it then mutates its own copy, and then it returns a copy back to the caller.

	// Why was this designed this way? This is becasue the idiom is to use value semantics to move a slice
	// value around our program. This must still be respected even when we are in a mutation operation
	// and working with reference types. Plus, again, value semantic mutation is the safest way to
	// perform mutation since the mutation is being performed on the function's own copy of the data
	// isolation.

	// Also append ALWAYS MAINTAINS a contiguous block of memeory. As we remember the benefits of arrays
	// are the mechanical sympathies we get with the cache lines and the contiguous block of memory
	// giving us a predictive pattern. Slices keep with this in always maintaining a contiguous block
	// of memory for the slice's backing array, EVEN AFTER GROWTH. This is again importantisimo for
	// the hardware.

	// Every time the append function is called, the function checks to see if the len and cap are the
	// same or not. If its the same, it means there's no more room in the backing array for the new
	// value. In that case, append creates a new backing array (doubling its size if its elements are
	// less than 1000, or growing by 25% is > 1000) and then copies the values from the old array into
	// the new one. Then the new value can be appended and the old backing array can be wiped. That old
	// backing array being wiped depends on there not being anymore references to it. This is why we
	// use append inside functions and make sure it returns to the original var and not to a new var.
	// if it returns to a new var than more than likely the old var is still referencing the old
	// backing array and that will be a memory leak.

	// So if we have a slice of 3 elements  with len == cap the slice structure would be:
	// [*backingArray][3][3] and after append it would be [*newBackingArray][4][6] NOTE the efficiencies
	// If we then append a 5th element, len != cap so no new allocation needed, the new structure will
	// simply be [*backingArray][5][6] and we can continue to append until len==cap again. Then cap
	// will be doubled to 12 in new backing array [*newBackingArray][7][12] etc.

	// To be more exact, when a backing array has 1024 elements of capacity or less, new backing arrays
	// are constructed by doubling the size of the existing array. Once the backing array grows past
	// 1024, growth happens at 25%. - NOTE this was changed in version 1.18, per the spec : "The built-in
	// function append now uses a slightly different formula when deciding how much to grow a slice
	// when it must allocate a new underlying array. The new formula is less prone to sudden transitions
	// in allocation behavior." This can be seen in the results of running this example.

	/*
			ADDENDUM - Bill also had the new formula cited in his notes. From the Go team we have the
		 	following:

		 	runtime: make slice growth formula a bit smoother

		 	Instead of growing 2x for < 1024 elements and 1.25x for >= 1024 elements,
		 	use a somewhat smoother formula for the growth factor. Start reducing
		 	the growth factor after 256 elements, but slowly.

		 	starting cap    growth factor
		 	256             2.0
		 	512             1.63
		 	1024            1.44
		 	2048            1.35
		 	4096            1.30

		 	(Note that the real growth factor, both before and now, is somewhat
		 	larger because we round up to the next size class.)

		 	This CL also makes the growth monotonic (larger initial capacities
		 	make larger final capacities, which was not true before). See discussion
		 	at https://groups.google.com/g/golang-nuts/c/UaVlMQ8Nz3o

		 	256 was chosen as the threshold to roughly match the total number of
		 	reallocations when appending to eventually make a very large
		 	slice. (We allocate smaller when appending to capacities [256,1024]
		 	and larger with capacities [1024,...]).
	*/

}
