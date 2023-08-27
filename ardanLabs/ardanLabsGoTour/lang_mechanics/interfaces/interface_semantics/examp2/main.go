/*
	METHOD SETS

	Before moving beyond the value/pointer semantics provided with interfaces. Its important to review
	the rules for method sets again. Method set rules help describe when a piece of data of a given type
	can be stored INSIDE of an interface. These rules are all about keeping the integrity of our programs.

	A simple way to look at it:

	IMPLEMENT INTERFACE					STORE VALUE					STORE ADDRESS
	Value Receivers						Yes							Yes (mixing semantics)
	Pointer Receivers					No (mixing semantics) 		Yes

	Basically what the above is saying is the following as stated by the method set rules:
	1. When an interface is implemented using a value receiver (value semantics), copies of values
	   and addresses can be stored inside the interface. However...
	2. When the interface is implemented using a pointer receiver (pointer semantics), only copies of
	   addresses can be stored.
	   		- when we say 'stored' we mean the concrete value or instance of a type that will be used with the
			  the interface. So if a function takes an interface, whatever arg we give to that function, which
			  will be either a copy of a value or a copy of an address, is said to be stored inside of the
			  interface. (apparently we can also do this outside of a function by simply using a '=')

	So then the question now comes up, Why doesn't Go allow copies of values to be stored inside the interface
	when pointer semantics are being used? The answer is a two-part integrity issue.

	FIRST REASON IS, we can't guarantee that every single value is addressable. If we can't take a value's address
	then it can't be shared or pointed to and therefore a pointer receiver method can't be used. Almost every
	value in Go is addressable, but not 100% of them as we see below
*/

package main

import "fmt"

type notifier interface {
	notify()
}

// a type of duration is declared to implement the notifier interface above. We do so with pointer semantics
// as seen in the notify method below.
type duration int

func (d *duration) notify() {
	fmt.Println("Sending notification in",*d) 
}

func main(){
	// We then convert the LITERAL VALUE (remember literals are compile time values, meaning they aren't placed
	// in any vars, they are just being used directly) into a value of type 'duration' then call the notify 
	// method on it. This call isn't allowed though becasue we are using an UNADDRESSABLE LITERAL VALUE instead
	// of placing it in a var and creating a copy on local stack mem. 
	// This method requires the duration value to be shared, but without an address, that's not possible. This
	// explains the first reason why an interface that is implemented using pointer semantics only allows copies
	// of addresses to be stored. The compiler can't assume that it can take the address of any given value for
	// a type that implemented the interface using pointer semantics because it won't be able to in 100% of the 
	// cases.  
	duration(42).notify() 
}

/*
	The SECOND REASON is just as important, and a huge win for integrity. If we look at the method set rule
	for pointer semantics again:

	Pointer receiver interfaces can't store values but they can store addresses. 

	This part of the rule is preventing us from storing copies of values (value semantics) inside the interface
	if we've implemented the interface using pointer semantics. The rule enforces the idea that, if we change
	the semantic from pointer to value, it crosses a dangerous line. We can only share values with the interface
	and never store actual values if we implement the interface with a pointer receiver. WE CAN NEVER ASSUME
	THAT IT IS SAFE TO MAKE A COPY OF ANY VALUE THAT IS POINTED TO BY SOME POINTER AND POSSIBLY CHANGING
	UNDERNEATH US. 

	When we look at the rule for value semantics again:
	Value receiver interfaces can store values but they can also store addresses (mixing semantics)

	This part of the rule is allowing us to store copies of the values (value semantis) and addresses (pointer
	semantics) inside the interface if we implement the interface using value semantics. The rule supports the
	idea that if we change the semantic from value to pointer, it can be safe. However, there is a word of
	caution related to this idea. MIXING SEMANTICS IS AA CONSISTENCY ISSUE THAT MUST BE PERFORMED AS A 
	CONCIOUS EXCEPTION. Consistency is everything, and mixing semantics can create unexpected side effects
	in our code. 

	Next let's look at why we consider interfaces as valueless:
*/
