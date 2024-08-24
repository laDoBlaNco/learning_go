package main

import(
	"fmt"
	"unicode/utf8" 
)

/*
	Let's take a look again at strings in Go.
	
	The interesting thing about strings in Go is that everything is based on UTF8. its all bytes
	but its assumed that we ar working with runes coded to UTF-8.  
	
	UTF-8 starts as bytes and become code points which are then looked at as characters, but again
	its all just bytes. 
	
	slices and strings are the core data structures in Go because we get so much mechanical 
	sympathy from them. Remember the connection to arrays and really the array is the most 
	important data structure in programming, hardware wise and with the cache lines, etc. And
	from what I'm learning almost every other Go data structure is implemented in some way with
	arrays. numbers are constant untyped literals, strings, arraays, slices, etc are all implemented
	with arrays.
*/



func main(){
	
	// Declare a string with both chinese and english chars
	s := "世界 means world" // 2 word data structure [*string][len 18] 3 bytes for each chinese
	// char and then 1 byte for each english letter. This is a string
	
	// utfmax is 4 -- up to 4 bytes per encoded rune (char)
	var buf [utf8.UTFMax]byte  // this is an array since we have the constant 4 in the []byte
	
	// iterate over the string.
	for i,r :=range s{ // using value semantics because this is a collection of strings
		// capture the number of bytes for this rune
		rl := utf8.RuneLen(r) // runelength
		
		// Calculate the slice offset for the bytes associated with this rune
		si := i + rl
		
		// copy of rune from the string to our buffer
		copy(buf[:],s[i:si]) // NOTE we have to convert our array to a slice to write to it.
		// copy only works with slices. but that's ok because "Every array in Go is just a slice
		// waiting to happen!" 
		
		// Display the details
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n",i,r,r,buf[:rl]) 
	}
	
}

func inspectSlice(slice []string) { // the arg its asks for a a slice VALUE not a pointer to a slice
	fmt.Printf("Length[%d]  Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { // and here we use the VALUE semantic version of the for/range
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
	// How do we know the form of the for/range to use:
	// What is this a collection of? Slice of strings.
	// What is a string? A built-in type
	// What do we use for built-in types? Value semantics
}

