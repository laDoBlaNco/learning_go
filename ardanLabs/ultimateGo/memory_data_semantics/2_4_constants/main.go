package main

/*
	Constants in go are very interesting from an implementation standpoint. In most langs constants
	are read-only variables, but in Go they aren't variables at all. Constants only exist at compile-time
	They are very type-safe.

	Also the fact that they have a parallel type system. They can be of a type and of a kind
*/

// Untyped Constants - 'kinds' can have a precision of 256 bits as opposed to the lower precision of
// typed numbers.
const ui = 12345    // kind: integer
const uf = 3.141592 // kind: floating-point
// literal values are also constants, or of a 'kind', which means that they can be implicitly converted
// by the compiler as well.

// Typed Constants still use the constant type system but there p is restricted
const ti int = 12345        // type:int - since its typed its bound to the laws of this type
const tf float64 = 3.141592 // type: float64

// This would error = ./constants.go:XX constant 1000 overflows uint8
// const myUint8 unint8 = 1000

// What makes constansts valuable is 'kind promotion'.

// variable answer will of type float64
var answer = 3 * 0.333 // since these are of kind the compiler will implicityly convert with no
// lost of integrity KindFloat(3) * KindFloat(0.333)

// constant third will be of kind floating point
const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

// Constant zero will be of kind integer
const zero = 1 / 3 // KindInt(1) / KindInt(3) both are int Kind so answer is zero

// This is an example of constant arithmetic between type and untyped constants. Kind is always promoted
// to type.
const one int8 = 1
const two = 2 * one // int8(2) * int8(1)

// So constants can be of a KIND or a TYPE, if they are type they follow the rules for type and constants
// only exist at compile-time. When of a kind the compiler can do implicit conversions at compile-time, but
// if they are typed, then we need to do it explicitly.

const (
	// Max integer value on 64bit arch
	maxInt = 9223372036854775807

	// Much larger value than int64 since its a 'kind'
	bigger = 9223372036854775808543522345

	// This will not compile - as a type it overflows
	// biggerInt int64 = 9223372036854775808543522345
)

// iota is another interesting part to constants. Its always used with a block syntax and it starts with
// 0 unless you calculate something else
const (
	a1 = iota // 0 : start at 0
	b1 = iota // 1 : Increment to 1
	c1 = iota // 2 : increment to 2
)

// normally iota would look more like this
const (
	a2 = iota
	b2
	c2
)

// and as I mentioned, we can calculate the starting point
const (
	a3 = iota + 1
	b3
	c3
)

// Something we can do with this iota and the Time package uses it a lot, is to create something like
// a bitmap
const (
	Ldate = 1 << iota // 1 : shift to the left 0, this shift will then move to the left on each constant
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
)

func main() {
	// numbers are super long because of the 256 bit precision
	println(ui, uf, ti, tf)
	println(answer)
	println(third)
	println(zero)
	println(two)
	println(maxInt)
	println("1: ", a1, b1, c1)
	println("2: ", a2, b2, c2)
	println("3: ", a3, b3, c3)
	println("Log: ", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)

}

// - Constansts are very powerfull in Go
// - Constants are compile time only
// - They aren't variables at all
// - they have a parallel type system
// - The literal values we've been dealing with are also constants of a 'kind'
