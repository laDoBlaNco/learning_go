package main

import "fmt"

/*
	So up to know we've talked about the latencys from GC and also how
	we get data in and out of the machine. But we also need to look at the
	mechanics of it all and knowing when to use what in order for us to start
	to understand the cost of decisions we make as engineers. We don't hide
	cost and this helps us optimize for correctness. Keeping in mind all of the
	costs and the sympanthies will make us better Go devs.
*/

func main() {

	// Declare an array of five strings that is initialized to its zero default
	var fruits [5]string // here we allocate an array of 5 strings. But what does this
	// really mean. A contiguous block of about 40 bytes (each string is a 2
	// word value (5*2) and each word is 4 bytes (10*4)) each of which is set
	// to its 0 default value ""
	fruits[0] = "Apple" // this is really a copy. If we look at the 2 word value
	// that will rep this, its a pointer to the actual bytes for the characters that
	// make up the word "apple" and the 2nd word is the length of 5. So in reality
	// because our string values using pointer semantics in that first word there
	// is an efficiency in this sharing as our assignment above isn't creating its
	// on copy of "apple" to work with, but using the same. O sea we have both
	// value semantics and pointer semantics at the same time.
	//	- The 2 word value is copied as Go works with its own copy of that
	// 	- Both pointers (from both 2 word copies) are now point at the same "Apple" string.

	fruits[1] = "orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// iterate over the array of strings
	// There are 2 flavors of for/range and they have different value semantics.
	// this one is using value semantics as we are getting a copy of every
	// element to work with. If we leave off 'fruit' below and just use
	// for i := range fruits{}, this would be the pointer semantics. We aren't copying
	// any values but just indexing directly into our array. But back to using
	// our call below, on each iteration we are creating yet another copy of our
	// 2 word values, with another pointer pointing to the same string. so for
	// each iteration will have a total of 3 copies of each word value all pointing
	// to their respective string. But with all these copys, we still share the
	// uknown cost, which is the actual string.
	for i, fruit := range fruits {
		fmt.Println(i, fruit) // now here we have a call to print meaning that
		// we are going to move across the program boundary to another frame and
		// of course it will have yet another copy of our word value. so now on
		// each iteration and call to print, we have 4 copies of each word value,
		// but again efficiently sharing the same respective string values.
	}
}

// all of this helps us to come up with the following guideline. When our data is
// a number a string or a bool, we'll want to use our VALUE SEMANTICS to move that
// data around the program. Everyone gets their own copy and we don't share it
// across  program boundaries. this will help us reduce memory on the heap and
// then lower the latency costs with the GC. When everyone has their own copy of
// they value, all of that is on the stack in cache and the only thing that may
// need to go to the heap is where we have pointer semantics helping us with the
// efficiencies, as in above with 4 copies of each of our string word values, but
// pointing to only 5 actual underlying arrays on the heap. If it  was all on the
// heap, our GC would have to manage it all. But now its just managing the actual
// underlying array. the one exception to this is if we were trying to represent
// 'null' in which we would use pointer semantics, since the zero default of a pointer
// is null.
// all of this includes not just function parameters and vars, but fields and 
// types as well. 
