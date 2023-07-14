package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"unicode/utf8"
)

// ... for quick debugging
var p = fmt.Println

// a Go string is a read-only slice of bytes. The language and the stdlib
// treat strings specially - as containers of text encoded in UTF-8. In other
// langs, strings are made of chars. In Go, the concept of a char is called a
// 'rune' - it's an integer that represents a Unicode code point. More info here:
// https://go.dev/blog/strings for a good introduction.

func main() {

	// s is a string assigned a literal value representing the word "hello"
	// in the Thai language. Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	// Since STRINGS ARE EQUIVALENT to []byte, this will produce the  length
	// of the RAW BYTES stored within NOT THE ACTUAL CHAR COUNT
	p("Length:", len(s))

	// Indexing into a string produces the raw byte values at each index. This loop
	// generates the hex values of all the bytes that constitute the code point
	// s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	p()

	// To count how many actual runes are in a string, we can use the utf8 package
	// Note that the  run-time of RuneCountInString depends on the size of the
	// string, because it has to decode each UTF-8 run sequentially. Some Thai chars
	// are rep'd by multiple UTF-8 code points, so the result of this count may be
	// surprising.
	p("Rune count:", utf8.RuneCountInString(s))
	p()

	// A range loop handles strings specially and decodes  each rune along with
	// its offset in the string. (This i didn't realize. So Range has a special
	// response when it encounters chars (runes) and acts intuitively
	var myRuneCount int
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		myRuneCount++
	}
	// I figure using the fact that 'range' reacts differently to runes, I don't
	// want to use the unicode package, I can use that reaction to count runes
	// in a string but simply counting the iterations of the loop. And I get the
	// same result as if I used utf8.RuneCountInString(s)
	p("myRuneCount:", myRuneCount)
	p()

	// We can achieve the same iteration by using the utf8.DecodeRuneInString
	// function explicitly. so in other words what I did using range to get the
	// length, we can use utf8 package to make a normal for loop react to  runes
	// just like a range does.
	p("Using DecodeRuneInString:")
	// basically setting to vars (index,width) checking the width to advance to the
	// next index.
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// This will demonstrate for us passing a rune value to a function
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals (just as char literals in
	// some other langs). We can compare a rune value to a rune literal directly.
	if r == 't'{
		p("found tee") 
	}else if r=='ส'{
		p("found so sua") 
	}
}
// Super interesting stuff. 
