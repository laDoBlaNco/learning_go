package main

/*
	Now let's see what we can use these pointer semantics for with some user defined types
*/

// user represents a user in a system
type user struct {
	name  string
	email string
}

// Main is the entry point for the application - With all the different calls here in main it
// looks as if we are sharing code up and down the stack with the pointers we are using in the
// func as well as in the calls to println. But one of the great things about a stack is that its
// self cleaning. After the main routine passes back up the boundery, if it needs another frame
// it'll take the frame below it and initialize (or overwrite) it. so this being the case, if we
// were sharing back up the stack, we would have bugs cuz our values would be wiped. So what's
// happening? Heap memory? Static code analysis. Escape analysis is done at compile time to
// decide where the value should be constructed on the memory of stack or heap. That depends on how
// a value is shared. So values that return pointers (shared up) need to be created on the heap
// not the stack.

// So in other words, anything we share down the stack is fine cuz the memory above our current
// routine location is always valid. But anything shared UP the heap has issues, cuz the memory
// BELOW our current position is not valid. It might or might not be wiped. So with the static
// escape analysis we don't have to worry about where these values are created. Once the heap is
// involved, then the GC gets involved. Remember that stack memory is self cleaning, but anything
// that goes to the heap will be impacted by our INTERNAL LATENCY (our garbabe collector)
func main() {
	u1 := stayOnStack()
	u2 := escapeToHeap()

	println("u1", &u1, "u2", u2)

}

// We don't have constructors in Go but we do have factory functions to create instances of our
// types and return them back to the caller
// createUserV1 creates a user value and PASSES A COPY back to the caller
// For factory functions its good to look at the return type cuz this tells us a lot about
// the data semantics that are in plan. This one for example is a value semantic function as it
// returns a copy of hits own instance of the type. This is cuz its returning a copy not a pointer
// So the func does what it needs to and returns a copy for other to work on.
//
//go:noinline
func stayOnStack() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V1", &u)

	return u
}

// createUserV2 creates a user value and SHARES the value with the caller. so here we are returning
// shared access to the data that our function created.
//
//go:noninline
func escapeToHeap() *user { // NOTE: How we return a shared value or a pointer
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u

}

// How we use pointer semantics can cost us  readability. If we see our code reuturning a pointer
// then we know that there's cost and heap memory involved. But if we use the pointer up higher in
// the code so that when we return the 'value', we aren't sure what it is without having to go
// and read more code, then we are hurting our readability. For that we need to code with some
// guidelines.

// 1. If we are using construction to a variable we shouldn't use pointer semantics as it costs us
//    readability. If we construct directly to a reutrn value or using pointer semantics in a function
//    call, that's ok, cuz its obvious what we are doing. Don't optimize for laziness (easy code),
//    optimize for correctness and readability.

// 2. Use these -gcflags -m=2 compiler flags so that we know what's happening with out code, if values
//    are being created on the stack or on the heap during the profiling.

/*
	Escape analysis:
	The algorithm the compiler uses to determine if a value should be constructed on the stack or heap
	is called "escape analysis". The name of the algo makes it sound like values are constructed on the
	stack and then moved to the heap memory. This is NOT the case. The construction of the value only happens
	once, and the escape analysis algorithm decides where that will be (stack or heap). Only construction
	on the heap is considered an ALLOCATION in Go, since anything on the stack is self-cleaning, we don't worry
	about it.

	Understanding escape analysis is about understanding value ownership. The idea is, when a value is
	constructed within the scope of a function, then that function owns the value. From there we ask the
	question, does the value being constructed still have to exist when the owning function returns? If the
	answer is no, the value can be constructed on the stack. If the answer is yes, the value must be constructed
	on the heap.

	NOTE: the ownership rule is a good base rule for identifying code that causes allocations. However
	We must appreciate that escape analysis has flaws that can result in non-obvious allocations. Also
	the algo takes opportunitiesto leverage compiler optimizations to save on allocations.

	This is also a key reason as to why we think about the stack as being self-cleaning. Zero value intialization
	helps every stack frame that we need to be cleaned withou the use of the GC. The stack is self cleaning
	since a frame is taken and initialized for the execution of each function call. the stack is cleaned
	during function calls and not on returns because the compiler doesn't know if that memory on the stack
	will ever be needed again.

	Escape analysis decies if a value is constructed on the stack (the default) or the heap (escape). With
	one of the above functions we pass a copy of the value back to the caller, so it's safe to keep the value
	on the stack. With the other function, we pass a copy of the value's address back to the caller (sharing
	up the stack)so its not safe to keep the value on the stack. 
*/
