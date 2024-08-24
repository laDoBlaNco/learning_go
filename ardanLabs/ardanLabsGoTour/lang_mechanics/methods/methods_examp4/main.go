package main

import "fmt"

/*
	KNOW THE BEHAVIOR OF THE CODE:

	Knowing the data semantics at play allows us to know the behavior of our code and that
	is vital. If we know the behavior the code then we know the cost. If we know the cost,
	NOW WE ARE ENGINEERING!

*/

// let's look at another type and method set below:
type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My name is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "is age", d.age)
}

func main() {
	// here we see once again the construction of our same data type and method sets.
	// we again assign it to d and then assign the d.method to f1 and since methods in Go
	// are just values and belong to the set of internal types, this is just an assignmenet
	// like any other type.
	d := data{
		name: "Kevin",
	}

	f1 := d.displayName

	// after the assignment we can call the method indirectly through the use of the f1 variable
	// This displays the name'Kevin'. Then we try to change the data so the name shows Odalis,
	// and call the method once again with f1. But we don't see the change. Why?
	// It of course has to do with the data semantics at play. The displayName method is using
	// a value receiver so value semantics are at play. This means that f1 var maintains and
	// operates on its own copy of d. so calling the method through the f1 var, will always use
	// the copy and that copy is protected against change. This is what you  want with value
	// semantics, the protection against external mutation.

	f1()
	d.name = "Odalis"
	f1()

	fmt.Println()

	// Now let's do the same thing with the setAge method.
	d2 := data{
		name: "Kevin",
	}

	f2 := d2.setAge
	f2(46)
	d2.name = "Odalis"
	f2(48)
	
	// This time the setAge method is assigned to the var f2 and the method is indirectly 
	// executed through the f2 var. We make the calls and this time we see the change. Since
	// setAge is using pointer receiver it doesn't operate on a copy of d2, but is operating
	// directly on the d2 variable. So f2 has shared access
	
	// If we don't know the data semantics at play, then we won't know the behavior of the
	// code. These data semantics are real and affect the behavior. 
	
	fmt.Println()
	fmt.Println("Using Func Types:") 
	d3 := data{
		name:"Kevin",
	}
	
	// Let's use fireEvent1 handler that accepts any function or method with the right sig
	fireEvent1(event) // here we accept our 'event' function and the fireEvent1 uses it to 
	// call on our "anonymous" message.
	fireEvent1(d3.event) // Here its used to call our method with the d.name and the message. 
	
	// Now using the fireEvent2 handler that accepts any func or meth of type 'handler' OR any
	// literal func or meth with the right sig
	fireEvent2(event)
	fireEvent2(d3.event) // same behavior as above, but this looks for the func sig as well as the
	// type 'handler' 
	
	// let's crate a var of type handler for the global and meth based event functions
	h1 := handler(event)
	h2 := handler(d3.event) 
	
	// Now let's use the fireEvent2 handler that accepts values of type handler
	fireEvent2(h1)
	fireEvent2(h2) 
	
	// and now the fireEven1 handler that only looks at the func sig
	fireEvent1(h1) 
	fireEvent1(h2)

	// so basically we can use func types either explicitly or through a custom type to
	// manipulate or determine the behavior based on a func signature in addition to the data
	// that is being used. 

}  

// In this sample let's also talk a little bit about Function types. let's start by adding another
// func to display global events
func event(message string){
	fmt.Println(message) 
}

// now let's add 'event' method to our data type as well
func (d *data) event(message string){ // using pointer semantics here
	fmt.Println(d.name,message) // we are just getting info though, Why are we using pointers sem???
}

// fireEvent1 uses an anony func type
func fireEvent1(f func(string)){ // so f is of type func and we see the type is a func sig
	f("anonymous") 
}

// handler reps a func for handling events
type handler func(string) // its a type of type func with the same func sig

// fireEvent2 uses a function type as it uses the type of type func we created above
func fireEvent2(h handler){
	h("handler") // so this is used the same as our  fireEvent1 except instead of using the
	// func type directly in the args we use a custom type based on our func type
}


/*
	NOTE:
		- Methods are functions that declare a receiver variable
		- Receivers bind a method to a type  and can use value or pointer semantics
		- Value semantics mean a copy of the value is passed across program boundaries
		- Pointer semantics mean a copy of the values ADDRESS is passed across program boundaries.
		- Stick to a single semantic for a given type and be consistent. 
		- Methods are only valid when it is practical or reasonable for a piece of data to
		  expose a capability.

*/
