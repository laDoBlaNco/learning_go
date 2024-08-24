// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0


// STACK VS HEAP:
// Let's start with a quote from Ayan George - "The STACK is for data that needs to persist only for the 
// lifetime of the function that constructs it, and is reclaimed without any cost when the function exits.
// The HEAP is for data that needs to persist after the function that constructs it exits, and is reclaimed
// by a sometimes costly garbage collection." - Therefore we can deduce that a better design would be avoiding
// escapes to the heap and thus the need for stack growth

// STACK GROWTH - The size of each frame (sandbox) for every function is calculated at compile time. This makes
// sense since the compiler can tell what size the value is and therefore can place it on the stack. This also
// means if the compiler doesn't know the size of a value at compile time, the value must be constructed on
// the heap. An example of  this is using the built-in function 'make' to construct a slice whose size is
// based on a variable. 
// b := make([]byte,size) // this backing array would be created on the heap

// Go uses a contiguous stack implemenation to determine how stacks grow and shrink. One alternative Go could
// have used is a segmented stack implementation, which is used by some OS's. 

// Every function call comes with a little question, "Is there enough stack spacd for this new frame?". If yes
// then no probelem the frame is taken (sliced off) and initialized. If not, then a new larger stack must be
// constructed and the memory on the exisiting stack must be copied over to the new one. This requires 
// changes to pointers that reference the memory on the stack. The benefits of contiguous memory and linear
// traversals with modern hardware is the tradeoff for the cost of making that copy. Because of the use of
// contiguous stacks, no Goroutine can have a pointer to some other Goroutine's stack. There would be too
// much overhead for the runtime to keep track of every pointer to every stack and readjust those pointers
// to the new location on stack growth. 

// Sample program to show how stacks grow/change.
package main

// Number of elements to grow each stack frame.
// Run with 1 and then with 1024
const size = 1024

// main is the entry point for the application.
func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
//
//go:noinline
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
// The results of this experiment show that the stack block in memory changes. When the size if 1, on recursion
// the memory (stack) address is the same almost completely for all 10 runs. When we change it to 1024 then
// the memory location changes at least 2 times during the 10 runs, meaning that it was copied to a new 
// location at least twice due to the need to grow larger than 2k.

