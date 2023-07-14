package main

import "fmt"

func main() {
	var i any
	// go < 1.18:
	// var i interface{}

	// since all types can implement the empty interface we are basically bypassing
	// the go type system as everything takes everything.
	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	// RULE OF THUMB: Don't use any :) - When using any you have to think really hard
	// about your design. Does it really make sense or not. Its like using the
	// unsafe pack. Its possible but rarely necessary. For example 'printing' is
	// something so general that it makes sense using ...any (also in json and
	// marshaling) but again most of the time it doesn't make sense to use.

	// We can convert
	s := i.(string) // type assertion - extract the underlying type from an interface
	fmt.Println("s:", s)

	// n := i.(int) // and if I try to assert to an int, we panic, its a string not int
	// fmt.Println("n:", n) // we avoid the panic using it with ok

	// This is known as 'comma,ok' this is idiomatic go.
	n, ok := i.(int) // NOTE: the 'ok' isn't an error its a bool returned by assertion
	if ok {
		fmt.Println("n:", n)
	} else {
		fmt.Println("not an int")
	}

	// we can also do these checks with type switches
	switch i.(type) { // type switch - Note: the assertion using the word 'type'
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}

	fmt.Println(maxInts([]int{3, 1, 2}))
	fmt.Println(maxFloat64([]float64{3, 1, 2}))
	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]float64{3, 1, 2}))

}

// We can also now change interfaces to accept multiple types
type Number interface {
	int | float64
}

// with the above we could use - func max[T Number](nums []T) T{
// to simply give our constraint a type name. But that name only works inside
// the constraint. Trying to use func max[T Number](nums []Number) Number{ Fails

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// When we start to see that we have things in common between functiosn and it
// makes sense to unify them, then we can start to use generics (Example above)

// Mainly there ar only 2 sitations where you want to use generics:
// 1. When you see the same code with the type being the only difference
// 2. When you are writing generic based structure (linked list, binary trees) DSA
//    Very new to Go. The team has decided to not add anything related to generics
//    other than support for them in the STL. In 1.18 only 2 things happened:
//		1. The empty interface (interface{}) is now called 'any'
//  	2. There is a constraint called comparable.
//
// Go 1.19 did add an atomic type, but really the Go team is letting the community
// decide how they are going to use generics and from there they'll figure out how
// to best  add them to the STL

// There is a series of exp packages. They are maintained by the official Go team
// but they aren't in the STL because either not enough people use them or they are
// considered experiments (thus 'exp') 'golang.org/x/exp'. Here we find things like
// 'slices' which are tools to work with slices such as 'contains', 'max', etc.

func maxInts(nums []int) int {
	if len(nums) == 0 { // here we address getting an empty slice, probably better with
		return 0 // an error in practice
	}

	max := nums[0]               // here set max to the first slice item
	for _, n := range nums[1:] { // here we start to loop through the slice from 2nd item
		if n > max {
			max = n // if we find a bigger item turn it into the new max
		}
	}
	return max
}

func maxFloat64(nums []float64) float64 {
	if len(nums) == 0 { // here we address getting an empty slice, probably better with
		return 0 // an error in practice
	}

	max := nums[0]               // here set max to the first slice item
	for _, n := range nums[1:] { // here we start to loop through the slice from 2nd item
		if n > max {
			max = n // if we find a bigger item turn it into the new max
		}
	}
	return max
}
