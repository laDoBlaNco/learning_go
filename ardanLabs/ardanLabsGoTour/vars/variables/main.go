// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// The playground is actually a 64-bit env with 32-bit pointers
// The os/arch combo is named nacl/amd64p32 - This is important as it'll impact some of the concurrency
// stuff. Also when we thinking about the architecture and the main Goroutine treads, we need to keep in
// mind that the playground is a single threaded environment. 

// Variabls are at the heart of the Go language and provide us the ability to read and write to memory. In Go
// access to memory is TYPE SAFE. This means  the compiler takes types SERIOUSLY and will not allow us to use
// variables otuside of the scope of how they are intended (declared. 


// Sample program to show how to declare variables.
package main

import "fmt"



func main() {
	// BUILT-IN TYPES:
	// Types provide INTEGRITY and READABILITY by asking 2 questions of all values:
	// 		- What is the amount of memory to allocate? (e.g. 1,2,4,8 bytes)
	// 		- Waht that memory represent? (e.g. int, uint, bool, ...)
	// Types can be specific to a precision such as int32 or int64:
	// 		- uint8 for example, represents an unsigned integer with 1 byte of allocation
	// 		- int32 for example, represents a signed integer with 4 bytes of allocation
	// When we declare  a type using a non-precision based type (uint, int) the size of the value is based on
	// the arch being used to build the program:
	// 		- 32 bit arch: int represents a signed int at 4 bytes of memory allocation
	// 		- 64 bit arch: int represents a signed int at 8 bytes of memory allocation
	
	// WORD SIZE
	// The 'word size' represents the amount of memory allocation required to store integers and pointers for
	// a given arch. For example:
	// 		- 32 bit arch: word size is 4 bytes of memory allocation
	// 		- 64 bit arch: word size is 8 bytes of memory allocation
	// The reason why this matters is because Go has internal data structures such as maps, channels, slices, 
	// interfaces, and functions that store integers and pointers. The size of these data structures will be 
	// based on the arch being used to build the program. 
	// The amount of memory allocated for a value of type int, a pointer, or a word will always be the same on
	// the same arch. 
	
	// ZERO VALUE CONCEPT
	// Every single value you construct in Go is initialized at least to its zero value default state, unless 
	// you specify the initialization value at construction. The zero value is the setting of every bit in 
	// every byte to zero. 
	// This is done for data integrity as it ensures that we don't get any garbage data when we declare vars
	// without an initial value, and its not free. It takes time to push electrons throug the machine to reset
	// those bits, but you should always take integrity over performance. 
	
	// Zero Values:
	// Type Initialized Value
	// Boolean false
	// Integer 0
	// Floating Point 0
	// Complex 0i
	// String "" (empty string)
	// Pointer nil


	// The keword 'var' is used to declare variables that are set to their zero value. Its a good practice
	// to use 'var' whenever we want to declare without initialization (zero defaults) and the short
	// declaration when we are initializing the value.
	var a int // 0
	var b string // ""
	var c float64 // 0.0
	var d bool // false

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)
	
	// Let's take a side step and talk about strings in Go. Strings use utf-8 char set, but are really just
	// a collection (slice) of bytes in the end. So anything we do with slices we can do with strings in
	// reality. 
	// A string is a two-word internal data structure in Go. Remember that words are for integers (everything 
	// we do in programming is manipulating integers (or floating point if we work with GPUs)) and pointers.
	// 		- The first word represents a pointer to an underlying array (fixed-size) of bytes
	// 		- The second word represents the length (integer) or the number of bytes in the underlying array.
	// 		- If the string is set to its zero default, then the first word is nil and the second is 0.

	// Declare variables and initialize.
	// Using the short variable declaration operator. This declaration takes advantage of Gos type inference
	// system. Remember that its best to use this := declaration when we are setting the intial value
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	
	// CONVERSION OVER CASTING - Go does conversion and not casting. Instead of telling the compiler to map
	// a set of bytes to a different representation, the bytes need to be copied to a new memory location for
	// the new representation. This ensures once again program integrity and there is no change for data loss.
	// Again its costs some space complexity, but its integrity over performance, and the cost isn't much at
	// all.
	
	// Go does have a package  in the standard library called 'unsafe', if we need to perform casting. But we
	// should really avoid that and be honest with ourselves about WHY WE WANT TO USE IT. Performing a
	// conversion provides the highest level of integrity for these types of operations. 
	// Specify type and perform a conversion. Conversion is explicit for all 'Typed Values' 
	// Below we see the simple Go conversion syntax. 
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

/*

FINAL NOTES:
	- The purpose of all programs is to transform data from one form to the other
	- Code primarily allocates, reads and writes to memory
	- Understanding type is crucial to writing good code and understanding code.
	- If I don't understand the data, I don't understand the problem.
	- Therefore, I understand the problem better, by understanding the data.
	- When vars are being declared to their zero default, use 'var'
	- When vars are being declared and initialized, use the := var operator.

*/

