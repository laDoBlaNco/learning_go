package main

import (
	"fmt"
	"unsafe"
)

// Sample program to explore how interface assignments work when values
// are stored inside the interface. Bill talks a lot about values inside of
// intefaces. It also refers to the underlying type in the 2 word structure

// Let's create a notifier inteface to provide support to notifying events.
type notifier interface {
	notify()
}

// our user type
type user struct {
	name string
}

// this is the method that implements our notifier interface
func (u user) notify() {
	fmt.Println("Alert", u.name)
}

func inspect(n *notifier, u *user) {
	word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
	value := (**user)(unsafe.Pointer(word))
	fmt.Printf("Addr User: %p Word Value: %p Ptr Value: %v\n", u, *value, **value)
}

func main() {
	// create a notifier interface and concrete type value
	var n1 notifier
	u := user{"Kevin"}

	// Store a copy of the user value inside the notifier interface value
	n1 = u

	// We see the interface has its own copy.
	inspect(&n1, &u)

	// make a copy of the interface value
	n2 := n1

	// we see the interface is shareing the same value store in the n1
	// interface value.
	inspect(&n2, &u)

	// Store a copy of the user address value inside the notifier interface
	// value
	n1 = &u

	// We see the interface is sharing the u variables value directly
	// There is no copy.
	inspect(&n1, &u) 
}
