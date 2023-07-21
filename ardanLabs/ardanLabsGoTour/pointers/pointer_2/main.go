// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0


//Now in this example we are passing still a copy but not of the value it self
// but the ADDRESS  of the value
// Sample program to show the basic concept of using a pointer
// to share data.
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Note how we use &count here to pass the address of our count rather than
	// the value of. But again, its still passing a copy of that address.
	// so its still technically passing by value.
	// Pass the "address of" count.
	increment(&count)

	// now we print out the 'value of' and 'address of' count and note that 
	// its changed after the function call where in our first example it didn't
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

// In this example increment declares the func to accept its own copy of an 
// address (rather than a copy of a value) that points to the integer value
// Pointer variables are literal types and are declared using *
//go:noinline
func increment(inc *int) {

	// Increment the "value of" count (caller's int value) that the "pointer points to".
	*inc++
	// so here  our 'value of' is the pointer or the address of the value. Now here's
	// the important part. Note that our addr of is also different. this means that
	// we passed in a COPY of the address, so now we have a different address for our
	// address but all are pointing back to the same integer value.
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

// NOTE: 
// 		- Use pointers to share data
// 		- Values in Go are ALWAYS passed by value
// 		- "Value of", what's IN the box. "Address of" (&), WHERE's the box
// 		- The (*)operator declares a pointer variable and the "Value that the pointer
// 		  points to". When we say declares a pointer variable its because *var is 
//        a literal type, so for example 'inc *int' is a var in of type pointer to int
