package main

/*
	So we've talked about
		- value vs pointer semantics
		- heap vs stack memory
		- optimizing for laziness vs correctness
		- how our main routine moves on the stack and we can share up and down and the implications

	Now we need to talk about growth. If we want to get up to 50k goroutines, we need to understand
	how our program (stack) grows.
*/

// Every goroutine starts out at 2k. If we make function calls down the frame and get to the
// end, then we need to grow. This means that Go needs to go out and get a new stack twice as
// large and copy everything over. This is the same way Go works withs slices of varible growth.
// Here we have values on the stack that would need to then MOVE to another stack. But again this
// shows that there is always a cost. Acquiring more contiguous memory means extra cost in latency.

// Go uses a contiguous stack implementation to determine how stacks grow and shrink. One alternative
// is to use segmented stack implementations as do other OSs. Every function asks "Is there enough stack
// space for this new frame?"If there is there's no problem. If not, then a larger stack must be constructed and
// the memory on the exisiting stack must be copied over to the new one. This required changes to pointers
// that reference memory on the stack. The benefits of continguous memory and linear traversals with
// modern hardware is the tradeoff for the cost of the copy.

// Due to this, no Goroutine can have a pointer to some other Goroutine's stack. There would be too much
// for the runtime to keep track of every single pointer to every stack and readjust those pointers to the
// new locations. 

// Example
// Number of elements to grow each stack frame. Run with 10 and then 1024
const size = 1024

// main is the entry point as always
func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size of the stack.
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}
	stackCopy(s, c, a)
}
