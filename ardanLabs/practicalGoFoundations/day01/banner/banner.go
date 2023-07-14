package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6) // all args are passed by value in Go
	banner("G☺", 6)

	s := "G😀"
	fmt.Println("len: ", len(s))
	s = "G☺"
	fmt.Println("len: ", len(s))
	// code points = rune ~= unicode character
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r) // difference here being we are pulling
		} // the value/character in the range with is the rune while below its converted
	} // to the byte.
	// rune (int32)

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// byte (uint8)

	x, y := 1, "1"
	// Using Printf gems
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // when debugging/logging is better to use %#v

	fmt.Printf("%20s!\n", s)
	fmt.Println(
		isPalindrome("g"),    // -> true
		isPalindrome("go"),   //-> false
		isPalindrome("gog"),  //-> true
		isPalindrome("g😀😀g"), //-> true as my reverse function solutions is unicode aware
	) // using the []rune(s)
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	return s == reverse(s)
}
func reverse(s string) string {
	rns := []rune(s)                                    // convert to byte string, or runs to handle unicode
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { //create a var for first and last and stop when ==
		// swap the letters of the string, first with last, etc
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2 // := short var dec, type inference
	for i := 0; i < padding; i++ {
		fmt.Print(" ") // so we could've just used PrintF("%20s!",s) to move the text
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

// there are no JS style methods on strings. "kevin".substring blah blah blah
// We have the package strings which has all of that in there.

// Text is basic but its important. We are always going to come down to formatting text
// in any application.
