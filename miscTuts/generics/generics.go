// https://bitfieldconsulting.com/golang/generics
// Bitfield Consulting: Generics in Go
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ...for quick debugging
var p = fmt.Println

func main() {
	p("Generic PrintAnything (any):")
	//GENERIC PRINTANYTHING:
	PrintAnything("Hello!") // string
	PrintAnything(42)       // int
	PrintAnything(true)     // bool
	p()

	//STRINGER CONSTRAINT:
	p("Stringer Constraint:")
	// p(Join([]int{1,2,3})) // error "int doesn't implement String (missing method String)"
	p(Join([]Person{
		{"ladoblanco", 46},
		{"Odalis", 48},
		{"Kelen", 16},
		{"Xavier", 13},
	})) // this now works because Person has a String() method which was the constraint/interface

	p("Comparable constraint:")
	p(Equal(6, 9))
	p(Equal(6.9, 6.9))
	p()
	p(Max([]int{1, 6, 45, 69, 11}))
	p(Max([]string{"my", "name", "is", "ladoblanco"}))
	p()

}

// WHAT ARE GENERICS?
// In Go every value needs a type since Go is a typed language. When we write funcs
// we need to specify the type of their params in the func sig. This is inconvient
// when we are creating functions that take multiple types or can take 'any' type
// for that matter if the type isn't important in what our function is doing.
// For this we have generics and want to be able to declare generic functions where
// we don't need to know the specific type:
func PrintAnything[T any](thing T) { // T being our generic type and 'any' being the
	// constraints/limits of that type. So simply put, T will be whatever type thing is.
	p(thing)
} // example above in 'main'

// This new kind of parameter in [] is call type parameters. So generic funcs in Go
// take type parameters along with the normal params. They go after the name and
// before the normal args.

// CONSTRAINTS:
// So now let's talk about constraints. So PrintAnything was easy cuz fmt can do that
// by itself. But what if we want to create a strings.Join that takes any element e
// and returns a concat'd string?

/*
func Join[E any](things []E)(result string){//here again generic type is E and its
// constraints  is 'any' and our func will take a slice of E and return 'result' which
// is a string.

	for _,v:=range things{
		result+=v.String() // obviously this isn't going to work with types that
		// don't have a String() method, o sea not stringers.
	}
	return result

}
*/
// gopls gives me an error that my generic type E has no field String()
// Go needs to check in advance if E (whatever it is) has a string method. And since
// we don't yet know what E is, we can't check that. So this is where we need to
// 'constrain' our type E. Instead of 'any' it needs to use E types that have String()
// method implemented. For generics, Constraints are just Interfaces
type Stringer interface {
	String() string // so this interface is String() and return a string
}

// now we can use this as the constraint/interface rather than 'any' which in itself
// is just an alias for another inteface (interface{})

func Join[E Stringer](things []E) (result string) { //here again generic type is E and its
	// constraints  is 'Stringer' and our func will take a slice of E and return 'result' which
	// is a string.
	for _, v := range things {
		result += v.String() // now it works becasue it knows that Stringer types
		// have to have a String() method.
	}
	return result

}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v and I'm %d years old.\n", p.name, p.age)
}

// THE COMPARABLE CONSTRAINT
// As I learned in another tutorial, constraints based on method sets (Stringer)
// are useful, but what if we want to do something with our generic input that
// doesn't involve calling a method? For example if we want to write an Equal
// function that takes two params of Type T and return a bool
// Not gonna work:
/*
func Equal[T any](a,b T)bool{
	return a==b // this will give us: cannot compare a==b (incomparable types in type set)
}
*/
// This is a similar issue as we had with Join wtih the String(), but since
// we're not calling a method now, we can't use a constraint based on a method set
// we need to constraint T to only work with types that can work with == and !=
// These are known as 'comparable' types. Go gives us a built-in constraint for this
func Equal[T comparable](a, b T) bool {
	return a == b // now it'll work.
}

// and if we wanted to do something generic that didn't use a method and wasn't
// comparable? Like a Max func on a generic slice of E elements?
// not this way
/*
func Max[E any](input []E)(max E){
	for _,v:=range input{
		if v > max{// "cannot compare v u003e max (type parameter E is not comparable with u003e)"
			max=v
		}
	}
	return max
}
*/
// Since Go can't prove ahead of time that E will work with > (ordered) then Go won't
// approve this in the compiler. We can fix this by listing every possible allowed
// type constraint in an Ordered interface with a union of type approximations like
// ~int|~int8|~int16|... ~string...and so on.
// But this is a whole lot of typing, so Go has already builtin a constraint for
// this as well in a constraints package
func Max[E constraints.Ordered](input []E) (max E) {
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
} // now it works

// GENERIC TYPES
// so now we can write functions that can take any type, but what about creating
// types that can contain any type., like a slice of any type. We can do say with
// the same type params and constraints
type Bunch[E any] []E // this is a type based on []E (E being our generic type)
// we're saying that for any given element type E, a Bunch[E] is a slice of values
// of type E. NOTE: THIS IS THE TYPE NAME - Bunch[E] just like our arrays [5]int
// So Bunch[int] would be a slice of int, etc. The we create then in the normal
// way (kinda)
var x = Bunch[int]{1, 2, 3}

// or we can then write generic functions to take generic types
func PrintBunch[E any](b Bunch[E]) {
	for _, v := range b {
		p(v)
	}
}

// Methods as well:
func(b Bunch[E])Print(){
	p(b) 
}

// We can also apply constaints to generic types and get into some complex structures
type StringableBunch[E Stringer] []E // This works cuz I've already created the 
// Stinger constraint/interface above. Wo we are basically just adding that to our
// custom type based on generic E now with a constraint that it E must have a 
// method String().

