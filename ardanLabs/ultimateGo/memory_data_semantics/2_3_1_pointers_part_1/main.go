package main

/*
	Understanding pointers in go is critical for undestanding the costs of our design decisions.
	First we need to understand the main model for Go which is PASS BY VALUE.

		- Everything in Go is about making a copy of data as we transition across these program boundaries
		- When a Go program starts it looks to see how many CPUs and hardward threads we have on the machine
			- For example: a i7 CPU with 4 cores has 8 hardward threads, so thats 8 main Go applications
			- We work with a Processor over a Machine where the OS schedules on which HW thread the operation
			  will happen.
			- On top of that M (machine) we also have our G (goroutine) which is very similar to the machine
			  thread. Machine uses the hardware thread to work and our Goroutine (Application thread) uses
			  the Machine thread to work
			- Goroutine starts with a stack of around 2k. Every thread gets its own stack, includiing the M
			  which starts with a 1mg stack (Win). This is a huge difference in memory allocation. We could kick
			  off 50,000 goroutines at 2k each which no issue, and they'll grow as needed. If we try that with
			  Windows and its thread of 1mg stack, that would kill the machine. This is why goroutines are
			  considered green threads or light threads and part of the reason we can be so flexible when it
			  comes to concurrency.
			- So in summary, our go program starts up and we have our P and our M and our G with a 2k stack
			  to start executing its instructions. The M of course takes those and is working directly with
			  the hardware.
*/
// Sample pgoram to show the basic concept of pass by value.

// So every function in Go is a program boundary. When we call a function we are basically having the main
// Go thread cross over the program boundaries. With each function the stack is sliced and we havre a 'frame'
// or sandbox in which our goroutine will do its work. This protects the rest of our program from any modifications
// thus guaranteeing the INTEGRITY of our design.

// On  each function call there is some data transformation happening, we get our input (or create it within
// the function or frame if we don't need to get it from outside) and we put out output, or some side effect.
// Go always starts with our first 'sandbox' being the 'main' func, which typically has no inputs. If a
// function is called then our gr must cross over program boundaries into that next slice or sandbox (function)
// which again is a slice on the current stack. Its not in that sandbox and only has effect on what's happening
// in that frame. It'll need any input for it to work with, as we see below 'increment(count)'
func main() {

	// Declare variable of type int with a value of 10
	count := 10

	// display the 'value of' and 'address of' count
	println("count:\tValue of [", count, "]\tAddr Of [", &count, "]")

	// Pass the 'value of' the count
	increment(count) // this is 'pass by value' a copy of count is created and passed over the boundary
	// with the mechanics of a parameter. Then inside the increment function go works with a copy of that
	// count (inc). So again working with copies is the safest way to work with data, so any modification
	// only impacts what's in that function from

	// Once the gr returns from the function, its back into the 'main' frame. This is call 'Value Semantics'
	// This gives us some huge wins on integrity and even performance, but there's always a price. What's the
	// cost of 'value semantics'? Efficiency and at times this creates code complexity. If we pass by value
	// 4 different copies to 4 different boundaries and they all modify things in their  boundaries, but then
	// we want to get back an accurate view of what the data is, we now have 4 possibly different values. So
	// it may be unreasonable to use the value semantics and give everyone a copy since there is so much
	// modification going on.

	// The opposite of 'value semantics' is 'pointer semantics', where we share data across all the program
	// boundaries, so we gain efficiencies, as we don't need to worry about reassembling the data and its
	// changes, but we lose the benefits of these isolated frames working on their own copies. With these
	// pointer semantics we get side effects and this is what causes bugs. That's why functional programming
	// tries to get rid of these pointers and their side effects. 
}

// increment declares count as a pointer variable whose value is always an address and points to values
// of type int
//
//go:noinline
func increment(inc int) {
	// Increment the 'value of' inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
