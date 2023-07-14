package main

import "fmt"

/*

The Summary

The following points are explain in much detail below, but here is the summary that I retreived
after reading through 'https://gronskiy.com/posts/2020-04-golang-pointer-vs-value-methods/ for the
second or third time.

--	With ordinary variables, Go allows "everything to be called on everything": pointer.pointerMethod
	, pointer.valueMethod(), value.valueMethod(), and value.pointerMethod(). Mechanisms engaged are
	selectors and automatic dereferencing.
--	With interfaces, it IS PROHIBITED to assign value to an interface which has pointer methods
--	The above behavior is formally regulated by the notions of method sets and addressability
	Value type does not belong to pointer method sets because there value inside an interface
	is not addressable
--	In turn, reason for such behavior is Go's implementation of "interface by value" - interface
	always holds a copy, hence calling pointer method on a copy does not make much sense for the
	purpose of modifying the original caller.
--	In addition, there's a reasoning around type consistency when taking a pointer of the
	interace value.
--	It seems that the decision to keep 'interface by copy' is a manifestation of a more general
	paradigm 'no references, everything is a copy, deal with it' 

*/

type Person struct {
	Name string
	Age  int
}

func main() {

	val := Person{"Kevin", 45}
	pointer := &val

	fmt.Printf("Value created on \t%#v with address %p\n", val, &val)
	fmt.Printf("Pointer created on \t%#v wtih address %p\n", *pointer, pointer)

	val.valueMethod()
	pointer.pointerMethod()

	val.pointerMethod()   // implicit conversion happening here based on receiver
	pointer.valueMethod() // same here

	// so now using these methods on interfaces and they work as one would expect. Pointer method
	// with a pointer and a value method with a value both on an interface
	callValueMethodOnInterface(val)
	callPointerMethodOnInterface(pointer)
	// So now we start to  experiment and with the interface we do the same as we did with the
	// type methods. We start using the ValueMethod with a pointer and it works as expected.
	// Go does what it needs to under the hood.
	callValueMethodOnInterface(pointer)

	// Now we try vice versa calling a pointer method on an interface with a value...
	// callPointerMethodOnInterface(val) // and it explodes. Of the 4 options (2 for options/switches
	// for the receiver type and 2 for the calling interface), this is the only one that fails.
	// So this confirms that on a value receive method type, we can call both val and pointer VIA
	// THE INTERFACE. But with a pointer receiver method type, we can only use a pointer VIA THE
	// INTERFACE. the common demonitator is the interface. So the only time I really need to
	// worry about it is when we are using an interface and a pointer receiver. All other combos
	// Go will take care of it for me without me having to think about it.

	// But why?
	// The method set of a Pointer Type include all methods declared with receiver *T and T. So all
	// method sets are included. But the method set of the Value Type only includes methods declared
	// with the value receiver. So in the end YOU CAN'T USE A POINTER RECEIVER ON A VALUE TYPE VIA
	// AN INTERFACE.
}

/*
	In Go, methods are not semantically bound to the enclosing structures as we see
	in other OOP languages. As such, one declares them using the contructr of a
	receiver, which is the object on which the given method will be called. Loosely
	speaking this is analog of 'self' in python or 'this' in c++ and JS.

	a pointer receiver: in this case the method should be called on the pointer and
	can modify it.
	a value receiver: in this case the method is only called on a COPY of the obj
	which calls it.
*/

// Pointer type receiver
func (receiver *Person) pointerMethod() {
	fmt.Printf("Pointer method called on \t%#v with address %p\n\n", *receiver, receiver)
}

// Value type receiver
func (receiver Person) valueMethod() {
	fmt.Printf("Value method called on \t%#v with address %p\n\n", receiver, &receiver)
}

// Note the diff in the format args used above. since the first receiver is a pointer
// then the var receiver had to be deref'd to get the value and nothing to get the
// address.

// on the second version which is a valueMethod I didn't have to do anything to get
// the value but I had to &receiver to get the address.

// This shows that the obj is being handed to the func as a pointer or a value.

// THE CONFUSING PART:
// The part that gets confusing for us newcomers is the syntatic-sugary part of Go
// to automatically  deref pointers or take value addresses. For exaample my calls
// in main to val.valueMethod() and pointer.pointerMethod(). The receivers are
// respectively val = value and pointer = pointer. So works as you would expect.
// even if I flipped them:
// pointer.valueMethod() -- will implicitly convert to (*pointer).valueMethod()
// val.pointerMethod() -- will implicitly convert to (&value).pointerMethod()
// based on the method receiver.

// So with T reciever you can call both T and *T
// and with *T receiver you can call both T and *T
// ALL FOUR COMBOS WORK?!?! Yes they do. Note from Effective Go under the section
// Pointers vs Values:

/*
	This rule arises because pointer methods can modify the receiver; invoking
	them on a value would cause the method to receive a copy of the value, so
	any modifications would be discarded. The language therefore disallows this
	mistake. There is a handy exception, though. WHEN THE VALUE IS ADDRESSABLE,
	the language takes care of the common case of invoking a pointer method on a
	value by inserting the address operator automatically. In our example,
	the variable b is addressable, so we can call its Write method with just
	b.Write. The compiler will rewrite that to (&b).Write for us.

Another note about these rules from the language spec:

	As with selectors, a reference to a NON-INTERFACE method with a value receiver
	using a pointer will automatically dereference that pointer: pt.Mv is equivalent
	to (*pt).Mv.

	As with method calls, a reference to a NON-INTERFACE method with a pointer
	receiver using an addressable value will automatically take the address of
	that value: t.Mp is equivalent to (&t).Mp.

	— From Specs#Method_values

*/

// Now let's dive into INTERFACES:
// In Go an interface is a separate type that reps a set of methods or a set of Types. They are
// semantically decoupled (just like methods) from types which want to implement
// them. As we learned before there's no explicit way to say that a said type is
// implementing a certain interface. It just does implicitly if it has the right
// methods implemented to do so. Some call this Structural typing which is different
// from Nominal typing (C++) or Duck typing (Python/JS)

// Ian Taylor said that: 'Interfaces in Go are similar to ideas in several other programming
// languages: pure abstract virtual base classes in C++; typeclasses in Haskell; duck typing
// in Python; etc. That said, I'm not aware of any other language which combines interface
// values, static type checking, dynamic runtime conversion, and no requirement of explicitly
// declaring that a type satifies an interface. The result in Go is powerful, flexible
// efficient, and easy to write.

func callValueMethodOnInterface(v ValueMethodCaller) {
	v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
	p.pointerMethod()
}

type ValueMethodCaller interface {
	valueMethod()
}

type PointerMethodCaller interface {
	pointerMethod()
}
