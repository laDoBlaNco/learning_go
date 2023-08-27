/*
INTERFACES ARE VALUELESS

We might be thinging, since the second word of the interface structure (2-word structure) is always an
address to the concrete value being stored inside of it, there there is always going to be an address
that can be used to call the pointer receiver method. Why then would storing a value when pointer semantics
are used to implement the interface be restricted???

The fact of the matter is that the second word of the interface value storing an address is irrevelant. If
we consider this implementation detail when defining the method set rules, we are essentially allowing
implementation details to creep into the spec. From the specification point of view, how anything is
implemented is irrevelant as the implemenation is always changing. If we allow that to creep into our
design decisions then those changes will break our code down the line.

Actually a change was made in version 1.4 with the interface implementation. "In earlier releases, the
interface contained a word that was either a pointer or a one-word scalar value, depending on the type
of the concrete object stored. This implementation waas problematical for the garbage collector, so as
of 1.4 interface values in memory always hold a pointer"

We need to understand that interface values from our code perspective are VALUELESS. There is nothing
concrete about an interface value  in and of itself.

We can see this in the next example.
*/
package main

import "fmt"

type notifier interface {
	notify()
}

type duration int

func (d duration) notify() {
	fmt.Println("Sending notification in", d)
}

func main() {

	// The above is all the same as the last example, except we are now using value semantics for our
	// interface implemenation. We create below a var 'n' with its zero default  value, a nil interface
	// value. [nil][nil] The var n is valuleless, and not until line 45 does the interface value have any
	// concrete data.
	// NOTE this is precisely what brought me over here. Note that its the assignment '=' that actually puts
	// a concrete value into the interface. This assignment can be through a func (arg assignment) or
	// with '=' as we did here below. We are basically convering the literal 42 to a duration and then say
	// it is not a notifier. And we can do that cuz duration has the method notify.
	// If I comment out that method I get an error that says, duration(42) doesn't implement the notifier interface
	// because its missing the method 'notify'.

	// The only thing that makes an interface concrete is the data that is stored inside of it. The method
	// set rules define what data (values or addresses) can be stored based on how the method set was
	// implemented (using value or pointer semantics). Integrity and semantics are what define the rules. How
	// all that is physically done is an implementation detail and shouldn't be part of our design.


	var n notifier
	n = duration(42)
	n.notify()
}

// To really bring this home, let's look at another example

