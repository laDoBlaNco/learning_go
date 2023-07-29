// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// iota provides support that other languags provide with enumerators. These are successive 
// integer constants. Its possible the name comes from the integer function in APL. In APL, the
// function (represented by the ninth letter of the greek alphabet, iota) is used to create
// a zero-based array of consecutive, ascending integers of a specified length.



// Sample program to show how iota works.
package main

import "fmt"

func main() {

	// here we have a const using iota and it'll start a series of 0,1,2 as seen below for 
	// each var. Iota works within a constant block with the value 0 and each successive
	// constant declared in the block, iota increments by 1. 
	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1:", A1, B1, C1)
	
	// and here notice how iota doesn't necessarily need to be repeated. The successive nature
	// of the integer constants are assumed once applied. 
	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
	)

	fmt.Println("2:", A2, B2, C2)

	// and if we don't want to start at 0, or wanted to create a different pattern, we can
	// do so simply applying some math which will be applied to the increasing value of iota
	const (
		A3 = iota + 1 // 1 : Start at 0 + 1
		B3            // 2 : Increment by 1
		C3            // 3 : Increment by 1
	)

	fmt.Println("3:", A3, B3, C3)

	// below is a fancier example of this using some bit shifting math. We can use this to maybe
	// set flags or as we see below do bit operations, which again will increase accordingly 
	// to the values of iota. Below we are shift the bits of 1 by an increasing iota.
	const (
		Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
		Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
		Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
		Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
		LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}

// NOTES:
// 		- Constants ARE NOT VARIABLES
// 		- They exist only at compilation
// 		- Untyped constants can be implicitly converted where typed constants and variables 
// 		  can not.
// 		- Think of untyped constants as having a KIND, and not a TYPE
// 		- Learn about explicit and implicit conversions to complete understanding of untyped
// 		  constants.
// 		- See the power of constants and their use in the standard library.
