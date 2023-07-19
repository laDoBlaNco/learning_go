// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go get honnef.co/go/structlayout/cmd/...

// We do though need to understand the memory allocations in this context
// that are happening here with a struct. Remembering that TYPES are everything and types
// tell the compiler what amount of memory to allocate as well as what's represented. So
// what happens when we have something complex like a struct with different types in it
// determined by the user?

// Let's look at the example above.
// 		- a bool is 1 byte
// 		- a int16 int is 2 bytes
// 		- a float32 is 4 bytes
// So maybe we would say 7 bytes. But the actual answer is 8 bytes, because there is a system
// of byte padding system in Go and it places a padding byte between flat (1 byte) and the counter
// (2 bytes) to stretch that to 4. This is done for ALIGNMENT. This idea of alignment is to allow 
// the hardware to read memory more efficiently by placing memory on specific alignment boundaries
// The compiler takes care of the alignment boundary mechanics so we don't have to worry about that. 
// Depending on the size of a particular field and its placement in the struct, Go will determine
// the padding needed to be sympathetic to the hardware. 

// What if we added another byte, after the counter, so now its flat/counter/flag2/pi (1/2/1/4)?


// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// For example,if we placed another byte flag in our previous example there would be
// new padding. The main algorithm ensures that each of fields start on a boundary/alignment 
// based its bytes. Though its sounds complicated to explain its not. Basically 1 byte
// items (bool) will start on 0 1 2 3 4 5 etc. 2 byte items (int16) will start on
// 0 2 4 6 8 etc. 4 byte items (float32) will start on 0 4 8 12 16 etc.

// So based on that our flag1 would be on 0, our counter would be on 2 our flag2 
// would need to be on 4, so then our pi would need to start on 8 and what's left
// (1,5,6,7) would all be padding bytes.

// If we need to  minimize the amount of padding for performance or size complexity
// we must layout the fields from highest to lowest allocation. This pushes any 
// necessary padding bytes down to the bottom of the struct and reduces the total number
// necessary. For example if we sort descending the order of our struct to
// Pi(4),Counter(2),Flag1(1),Fla2(1) then we have 8 bytes and can release 4 bytes of padding.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe" // not sure what we need this for
)


// Below we have different examples of padding.
// No byte padding, because everything is 1 byte, no padding needed.
type nbp struct {
	a bool // 	1 byte				sizeof 1
	b bool // 	1 byte				sizeof 2
	c bool // 	1 byte				sizeof 3 - Aligned on 1 byte
}
type npb struct{
	a bool
	b bool
	c bool
}

	var ex1 nbp
	var ex2 npb

// Single byte padding. Here like our original example needs 1 byte of padding between 0 and 2
type sbp struct {
	a bool //	1 byte				sizeof 1
	//			1 byte padding		sizeof 2
	b int16 // 	2 bytes				sizeof 4 - Aligned on 2 bytes
}

// Three byte padding. Here is like the second half of our previous example, we need 
// 3 bytes of padding between 0 and 4
type tbp struct {
	a bool //	1 byte				size 1
	//			3 bytes padding		size 4
	b int32 //	4 bytes				size 8 - Aligned on 4 bytes
}

// Seven byte padding - because we have 1 byte on 0 and the next element (int64) must be on 8.
// That means we need 7 bytes of padding between the 1 and 8 alignments.
type svnbp struct {
	a bool //	1 byte				size 1
	//			7 bytes padding		size 8
	b int64 //	8 bytes				size 16 - Aligned on 8 bytes
}

// No padding. These are sorted descending, so large to small and therefore no padding is necessary.
type np struct {
	a string // 16 bytes			size 16
	b string // 16 bytes			size 32
	c int32  //  4 bytes			size 36
	d int32  //  4 bytes			size 40 - Aligned on 8 bytes
}

// Eight byte padding on 64bit Arch. Word size is 8 bytes. Here a string is 2 words (16) our int32 
// is 4 bytes. The next string though needs to align on 8 bytes (its word size) which would start on 24,
// so we would need 4 bytes between 20 and 24. Our string (16) would take us then to 40 and the last
// int32 would be 4 more to 44. Now I don't understanding why we need padding to finish it off though
// to 48, but apparently we do. 
type ebp64 struct {
	a string //	16 bytes			size 16
	b int32  //	 4 bytes			size 20
	//  		 4 bytes padding	size 24
	c string //	16 bytes			size 40
	d int32  //	 4 bytes			size 44
	//  		 4 bytes padding	size 48 - Aligned on 8 bytes
}

func main() {

	var nbp nbp
	size := unsafe.Sizeof(nbp)
	fmt.Printf("nbp  : SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)

	// -------------------------------------------------------------------------

	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("sbp  : SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// -------------------------------------------------------------------------

	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("tbp  : SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// -------------------------------------------------------------------------

	var svnbp svnbp
	size = unsafe.Sizeof(svnbp)
	fmt.Printf("svnbp: SizeOf[%d][%p %p]\n", size, &svnbp.a, &svnbp.b)

	// -------------------------------------------------------------------------

	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("np   : SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)

	// -------------------------------------------------------------------------

	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	fmt.Printf("ebp64: SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)
	
	// Assigning Values
	// ex1=ex2 // not allowed, compiler error because they are from a named struct, even though they are the
	// same shape/structure
	
	// To make this work we would need to simply convert one to the other. Now we could also change 
	// of them to be an unamed literal type and then it would work without the need for conversion. 
}

// ASSIGNING VALUES;
// Even though 2 different named structs have the same exact shape, if they are are 
// named structs, we can't assign a value of one to the other. 

// NOTES: 
// 		- We can use the struct literal form to initialize a value from a struct type
// 		- The dot (.) operator allows us to access individual field values
// 		- We can create anony structs


