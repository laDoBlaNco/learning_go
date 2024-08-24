package main

/*
	In particular the idea of value and pointer semantics is everywhere in Go
	As statedin earlier posts, semantic consistency is critical for integrity
	and readability. It allows devs to maintain a strong mental model of a
	code base as it continues to grow. It also helps to minimize mistakes, side
	effects, and unexpected behavior.

	INTRO
	In this post, we will explore how the interface in Go provides both a value
	and pointer semantic form. Bill will teach us the associated language
	mechanics and show the depths of these semantics. Then we'll see how the compiler
	attempts to intercede for us when we are mixing semantics in a dangerous way.

	Let's start with this listing 1
*/

// LANGUAGE MECHANICS
// An interface can store its own copy of a value (value semantics) or a value can
// be shared with the interface by storing a copy of the value's address (pointer
// semantics). This is where the value/pointer semantics come in for interfaces.
// This piece of code shows us how these semantics work

import "fmt"

// first we declare an interface named 'printer' with a method set of 'print()'
type printer interface {
	print()
}

// then we have our concrete type named 'user' that implements the printer interface
// below in with the method using value semantics receiver
type user struct {
	name string
}

func (u user) print() {
	fmt.Println("UserName:", u.name)
}

// here in main we create a user value and initialize it with "Kevin". Then we
// create a slice of printers
func main() {
	u := user{"Kevin"}
	// in this slice of printers we have both semantics. The value semantics
	// storing a copy of the user value in index 0 of our slice and then
	// a pointer to store a copy of the address inside of index  1
	entities := []printer{
		u,  // value
		&u, // pointer (value's address)
	}
	// the user value field 'name' is changed directly on our user value
	u.name = "Kevin_CHG"
	// finally we iterate over our slice using the value semantic version of a
	// for/range loop displaying the name for each user stored inside the
	// respective interface value.
	for _, e := range entities {
		e.print()
	}
}

/*
	Its intersting to see how these two interface values actually exist. In
	our first copye (value semantics) of the original user which is in index 0
	we have a straight copy of our 2 word interface structure with a copy of
	the original value "Kevin". But in index 1 we have a pointer that refs
	our original value.

	so when we iterate over our slice we get the copy (unchanged) and the
	original value which was changed back on line 54.

	REMEMEBER: Choosing one semanatic over the other is a decision that is made
	at the time of declaring or using a type. We want to maintain as much consistency
	with semantics as possible. This example we just went through is showing how to
	apply different semantics with interfaces, but isn't the typical model we want
	to be using with a mixture of semantics.

	In the next example, we'll see how Go protects us from making big mistakes
	if we are mixing semantics.
*/
