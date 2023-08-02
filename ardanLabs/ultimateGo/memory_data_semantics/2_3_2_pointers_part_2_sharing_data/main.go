package main

// with pointer semantics the code looks pretty much the same, but the behavior is very different.
// We still have or main routine kick off in our main func frame and the value of count:=10 is created
// then in that frame our goroutine comes across our function call and then passes over the program boundary
// into another slice of the stack (frame)

// Sample program to show the basic concept of using a pointer to share data

func main() {
	// Declare variable of type int with a value of 10
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of [", count, "]\tAddr Of [", &count, "]")

	// Pass the "address of" count to our function as a param
	increment(&count) // here we don't make a copy of the count value, but we make a copy of the actual
	// address of the voiue (pointer). NOTE: this is still a PASS BY VALUE not a pass by reference???
	// Its a copy of the address that being passed over the boundary so if its real data we need to store it
	// in a var which is what we do with (inc *int), by placing that * in front of the named type of the var
	// we are saying that we want to store the ADDRESS of an integer value (pointer) in that parameter.
	// A pointer variable is not a named type but its a literal type, so we get the pointer for free in that
	// it'll always be that 4byte or 8 byte value, not the actual memory allocation for the value.

	// This means that when our goroutine moves to that new frame in the function on our stack we now have
	// INDIRECT access to memory in the main frame. NOTE: the goroutine only had DIRECT access to the memory
	// in the frame we are currently in. So pointer semantics allow the goroutine to reach outside of its
	// sandbox to read and write.

	println("count:\tValue Of [", count, "]\tAddr Of [", &count, "]")

}

// increment declares count as a pointer variable whose value is always an address and points
// to values of type in.
//
//go:noinline
func increment(inc *int) {

	// Increment the "value of" count that the "pointer points to".
	*inc++ // read/modify/write the integer value that we have indirect access to with *. This is the
	// side effect. Its more efficient, but this is the scary part as we are modifying things outside
	// of the frame we are working with.

	println("inc:\tValue Of [", inc, "]\tAddr Of [", &inc, "]\tValue Points To [", *inc, "]")
	// the fact that the address we are passing in (inc *int) is different from the address of the var &inc
	// shows us that we are still passing by value, because we are working with a copy of the address, not
	// the address itself. ****
}

// NOTE: WHEN DEALING WITH POINTER SEMANTICS, WE ARE SAYING THAT THE GOROUTINE HAS SHARED ACCESS TO SOMETHING
// OUTSIDE OF ITS FRAME/SANDBOX, VALUE SEMANTICS MEANS THE GOROUTINE HAS ITS OWN COPY OF THE DATA. THESE
// SEMANTICS ARE BURIED VERY DEEP INTO THE LANGUAGE. SO WHEN DO WE USE ONE OVER THE OTHER?
// WE SHOULDN'T WORRY ABOUT THAT RIGHT NOW. WE'LL GET TO THOSE GUIDELINES LATER. RIGHT NOW WE NEED TO
// TRULY UNDERSTAND THE DIFFERENCE BETWEEN ONE SEMANTIC AND THE OTHER. 
