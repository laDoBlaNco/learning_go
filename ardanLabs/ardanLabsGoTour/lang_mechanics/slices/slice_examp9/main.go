package main

import(
	"encoding/binary"
	"fmt"
)

func main(){
	
	// LINEAR TRAVERSAL EFFICIENCY
	// let's talk a littl about linear traversal efficiency, which goes right along with 
	// our use of contiguous blocks of memory and our mechanical sympathies. 
	// The beauty of a slice is its ability to allow for performing linear traversals that are
	// mechanically sympathetic while sharing data using value semantics to minimize heap 
	// allocations. 
	
	x := []byte{0x0A, 0x15, 0x0e, 0x28, 0x05, 0x96, 0x0b, 0xd0, 0x0}
	
	a := x[0]
	b := binary.LittleEndian.Uint16(x[1:3]) 
	c := binary.LittleEndian.Uint16(x[3:5]) 
	d := binary.LittleEndian.Uint32(x[5:9]) 
	
	fmt.Println(a,b,c,d) 
	
	// This code is performing a linear traversal by creating slice values that read different
	// sections of the byte slice from beginning to end. So rather than creating a whole new
	// backing array or allocation separate blocks of memory, we have one contiguous block of 
	// memory and each slice is just a WINDOW into a specific section of that backing array.
	// All data in this code stays on the stack and there are no extra copies or allocations
	// being made. 
}

// NOTE:
// 	1. Slices are like dynamic arrays with special and built-in functionality
// 	2. There is a difference between a slice's length and capacity and they each serve a purpose
// 	3. Slices allow for multiple 'views' of the same underlying array
// 	4. Slices can grow through the use of the built-in function append
