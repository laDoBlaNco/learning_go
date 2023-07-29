// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// One of the more unique  things about Go is the implemenation of constants. The rules for constants in 
// this language specification are unique to Go. they provide flexibility Go needs to maek the code we write
// readable and intuitive while still maintaining type safety.

// Constants can be either TYPED or UNTYPED. When a constant is untyped, its considered to be a KIND. Constants
// of a KIND can be implicitly converted by the compiler. This all happens at compile time and not at runtime
// thus the type safety.

// Sample program to show how to declare constants and their
// implementation in Go.
package main

import "fmt"

func main() {

	// Constants live within the compiler. This means that they are given their values and created at
	// compile time.
	// They have a parallel type system in the sense that they can be both typed and untyped giving us the
	// flexibility of both implicit and explicit conversion when warranted. 
	// Compiler can perform implicit conversions of untyped constants.

	// Untyped Constants actually have a much larger precision as they are opened and not locked into any
	// box just yet as a Typed constant would be. So untyped numeric constants have a precision of 256 bits
	// and based on a 'kind' of number. 
	const ui = 12345    // kind: integer ,kind of integer not a type integer so not int16, 32, or 64
	const uf = 3.141592 // kind: floating-point, same here, not float32 or float64 specifically until
	// it needs to be one of them.

	// Typed Constants still use the constant type system but their precision
	// is restricted. o sea they are locked into a int32 or float64 box
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	// ./constants.go:XX: constant 1000 overflows uint8
	// const myUint8 uint8 = 1000
	// This doesn't work cuz the value 1000 is too large to fit in the uint8 type box.

	// Constant arithmetic supports different kinds.
	// Kind Promotion is used to determine kind in these scenarios and this is where we see some of the 
	// implicit conversion going on as we see in the few examples below. 

	// Variable answer will of type float64.
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)
	fmt.Println(zero)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math. So even if we are dealing with two untyped
	// kinds, some sort of conversion must happen in order to get the arithmetic done.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
	fmt.Println(two)
}
// Next let's look at the precision difference a little closer with another example.
