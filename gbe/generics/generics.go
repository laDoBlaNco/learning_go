// Go by Example: Generics
package main

import (
	"fmt"
)

var p = fmt.Println

func main() {

	// GENERIC FUNCTION:
	m := map[int]string{1: "2", 2: "4", 4: "8"}
	m2 := map[string]float64{"Pi": 3.1492, "fun": 69, "not_as_much_fun": 66}

	p("Generic Function: MapKeys")
	p(MapKeys(m))
	p(MapKeys(m2))
	// Even though we leave it out due to inference, below is using MapKeys with
	// its type parameters included. At times we'll need to do this for ambiguity
	// Note that if you put the wrong type params in below, you'll get an error
	// that: can use m2 of type map[string]float64 as type map[int]string in
	// argument to MapKeys[int,string] -- very specific errors ;)
	p(MapKeys[string, float64](m2))
	p("-------------------")

	// GENERIC TYPES: (LINKED LIST IMPLEMENTATION EXAMPLE)
	lst := List[int]{} // note that we had to put the Type in on creation. Does
	// this mean that when creating a custom type with constraints, we need to
	// add in the type, or that's just in this case that it was ambiguious so we
	// needed to tell Go what it was.That is exactly what it is. When I try to
	// create an empty List{} struct I get an error that I can't use generic
	// type List[T any] without instantiation and this is because Go can't
	// infer the generic type if I don't try to put something in it. So if I
	// need an empty struct to start I need to use the type param block List[T]
	// the question now is, if I create a struct literal with values, DO I STILL
	// NEED THE TYPE PARAMETER BLOCK IF I CREATE AN INSTANTIATED STRUCT LITERAL?
	// WILL GO INFER AS IT DOES WITH GENERIC FUNCTIONS????
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	p("list:", lst.GetAll())

}

// Here's an example of a GENERIC FUNCTION, MapKeys will take a map of
// any type and return a slice of its keys. This function has two type
// params - K and V; K has to be comparable so it will have a comparable
// type constraint (interface) meaning we can compare it with == and !=.
// This is required for all map keys in Go. V has the type constaint of
// 'any' as its not restricted. 'any' being the same as interface{}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m)) // make a slice of K, 0 length and capacity of whatever
	// the length of the original map is.
	for k := range m { // range through all the keys in m
		r = append(r, k) // append all the keys to our result slice
	}
	return r // return our result slice.
} // example use in 'main'

// As an example of a GENERRIC TYPE we have a singly-linked list. Note the
// use of type parameters based on other customer types. This type List
// can use values of any type.
type element[T any] struct { // generic T of any, note that they element is not
	// exported since it shouldn't be accessed from outside the linked list.
	next *element[T] // Here we have an example of being self referential as
	// next points to an element itself with a general T as its parameter as
	// its type.
	val T
}

type List[T any] struct { // our List is another struct with both the head
	// and tail being an element[T] type
	head, tail *element[T] // this is still questionable to me [T] but I think
	// it has to do with that fact that element has a constraint [T], but I'll
	// do some more tutorials based on type parameters to be sure. ...From what
	// I can tell when creating the type with type constraints. The constraint
	// becomes part of the type name and therefore is part of the type. O sea
	// type List[T any] struct{}  is then referenced with List[T]
	// type element[T any] struct{} is now referenced as element[T]
}

// With this we can now define methods on these generic types just like we
// do on regular types, but we have to make sure we keep the type parameters
// consistently in place. The type is "List[T]" and not List. This points to
// what I note above that 'List' is a type and this 'List[T]' is another type.
// the second one being a type with constraints (I think I'm getting closer).
func (lst *List[T]) Push(v T) { // so again our pointer receiver is of type List[T]
	// with the name of the method being Push which takes a v of T type.
	if lst.tail == nil { // if the tail is nil, the list is empty so we then set
		lst.head = &element[T]{val: v} // head to a new referenced element[T] type
		// literal pushing in v to {val:v}
		lst.tail = lst.head // then we set tail to that same reference.
	} else {
		lst.tail.next = &element[T]{val: v} // if lst isn't empty we just set
		// lst.tail.next to a new ref'd element[T] literal of {val:v}
		lst.tail = lst.tail.next // and we set the tail to that 'next element'
	}
} // so with this we are creating a new node (element[T]) with each push and
// giving it the value from our v of T type arg. And our list grows, with
// each 'next' connected to a new different element reference.

// here's a GetAll method example. Again note the type sigs in the receiver
func (lst *List[T]) GetAll() []T { // pointer receiver of type List[T], no args returning
	// a slice of T
	var elems []T                             // create our nil slice (zero default)
	for e := lst.head; e != nil; e = e.next { // I NEED TO INVESTIGATE HOW THIS
		// WORKS WITH A LINKED LIST. NORMALL FOR IS USED AS COUNTER BUT HERE ITS
		// AS IF GO SEES MY LINKED LIST AS A GO COLLECTION????
		elems = append(elems, e.val)
	}
	return elems
} // usage of the methods above in 'main'
