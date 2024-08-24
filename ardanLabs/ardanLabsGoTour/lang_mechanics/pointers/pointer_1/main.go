// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.

// So pointers are about SHARING values across program boundaries,
// remembering that program boundaries are when we move from one program
// frame to another, in which frame we only have read/write access to the
// elements in that frame. That's why we say boundaries. Its our sandbox
// or safe zone. There are several types of program boundaries, but the
// most common that we'll deal with is the one before funcs. There is also
// one between goroutines, which we'll look at later. But before we really
// get into pointers we need to understand the semantics around how the
// main goroutine and its memory management, etc.

// We'll get into the code in a bit, but when we start up a go program, the
// runtime create an inital Goroutine. This is the main routine of the
// program or the main thread as some say. Goroutines are lightweight app
// level threads (green threads) with many of the same semantics as system
// threads. Every Go program has at least or is at least 1 goroutine which
// is the MAIN routine.

// The key here is that each routine gets its OWN BLOCK of memory call the
// stack. In Go each stack starts out at 2k (2048) byte allocation. Its very
// small, but stacks grow in size over time. This is also a key point for us
// when we start to learn about Heap, GCs, and slices, etc which all have to
// do intimately with pointers and their impact on the world. So our program
// starts, we get our 2k stack and start processing in our MAIN 'frame' at
// top of our stack. I'll get into the 'frames' in a moment.

// Everytime a function is called, a block of the stack space is taken to
// or slice off, to help the main goroutine execute the instructions
// associated with that particular function. Each individual blick of memory
// sliced off is what we call a 'frame'. The size of a frame for a given
// func is calculated at compile time. No values can be constructed on the
// stack unless the compiler knows the size of that value at compile time.
// If the compiler doesn't know the size, then that value would need to be
// constructed on the 'heap', and we'll talk more about that in a bit, but
// I learned a lot about that when studying the Go garbage collector.

// Another thing interesting about stacks, since I did mention GC, is that
// they are self-cleaning and zero default helps with the initialization
// of the stack. Every time we make a function call, and a frame of memory
// is sliced off or blocked out, the memory for that frame is initialized
// or set to zero, and that's how its self-cleaning. On a function return
// the memory for the frame is left alone since its unknown if that memory
// will even be needed again, so it would be inefficient to initialize
// memory on every return, cuz again we don't know if we are going to use
// it or not.

// PASS BY VALUE:
// Prior to understanding passing by reference (pointer) we need to get
// in our minds passing by value. ALL DATA IN GO IS PASSED AROUND THE
// PROGRAM BY VALUE. This is the case even in when working with pointers
// and reference types and we'll see this in a little bit. So this means
// that the data is being passed across program boundaries in COPIES
// between each goroutine of function. Each has its own copy of the data
// to work with. This is the safest environment to work and mutate data.
// There are 2 types of data that we'll be working with. Either the value
// itself (int, string, user, etc) or the value's address. Addresses ARE
// ALSO DATA that needs to be copied and passed by value, stored and
// moved across program boundaries. This is why we say that in Go ALL DATA
// IS MOVED AROUND THE PROGRAM BY VALUE regardless if its a pointer or
// the value itself. The first example below attempts to explain this more.

// Sample program to show the basic concept of PASS BY VALUE.
package main

func main() {

	// Declare variable of type int with a value of 10. Note the use
	// of the short declaration :=
	count := 10

	// Display the "value of" and "address of" count. Note here the two
	// different types of data (value of/address of) that get moved around
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	// the use of &means we are getting the address of, or pointer to.

	// Pass the "value of" the count. When we see a function, we need to
	// start to think about it as going from one program boundary to another.
	// The value 'count' is being copied into our increment function to
	// be worked with. Its working on its own copy of count.
	increment(count)

	// here we print out the 'value of' and 'address of' count. The value
	// of count will not change after the function call. Again cuz this
	// function just worked on its own copy, printed what it needed to
	// (in which case that call to println also got its own copy) and then
	// returned nothing so it all went away.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//
//go:noinline
func increment(inc int) {

	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// Now let's dive into a pointer example.
