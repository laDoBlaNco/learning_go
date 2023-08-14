package main

import(
	"fmt"
	"unicode/utf8"
)

/*
	STRINGS AND SLICES;
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid utf-8 sequences. 
	Those sequences represent  Unicode code points, called runes
	No guarantee is made in Go that characters in strings are normalized
	
	---------------------------------------------------------------------------------
	
	Multiple runes can represent different chars:
	
	The lower case grave-accented letter à is a char, and its also a code point
	(u+00e0), but it has other representations as well.
	
	We can use the 'combining' grave accent code point, u+0300, and attach it to
	the lower case letter a u+0061, to create the same char. 
	
	In general, a char may be represented by a number of different sequences of
	code points (runes), and therefore different seqs of UTF-8 bytes.
	
*/

func main(){
	
	// Declare a string with both chinese and english chars.
	s := "世界 means world"
	
	// UTFMax is 4 -- up to 4 bytes per encoded rune (or 32bit (int32))
	var buf [utf8.UTFMax]byte // here we create a buffer array putting a number in []byte
	
	// iterate over the string
	for  i,r := range s{
		
		// capture the number of bytes for this rune
		rl := utf8.RuneLen(r)
		
		// Calc the slice offset for the bytes associated with this rune.
		si := i + rl 
		
		// copy of rune from the string to our buffer
		copy(buf[:],s[i:si]) // we use [:] because we can't use 'copy' with an array only a slice
		// and since "every array is just a slice  waiting to happen", we are good.
		
		// Display the details
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n",i,r,r,buf[:rl])  
	}
}
