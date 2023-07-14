package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	// strings are slices of bytes
	str := "A word"
	// Escape Sequences: \n \t \" \\

	replacer := strings.NewReplacer("A", "Another")
	str2 := replacer.Replace(str)
	pl(str2)
	pl("Length:", len(str))
	pl("Length:", len(str2))
	pl("Contains Another:", strings.Contains(str2, "Another")) // case sensitive
	pl("o index:", strings.Index(str2, "o"))
	pl("Replace:", strings.Replace(str2, "o", "😀", -1)) // the -1 tells Go replace ALL matches
	str3 := "\nSome Words \n"
	str3 = strings.TrimSpace(str3)
	pl(str3)
	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower:", strings.ToLower(str2))
	pl("Upper:", strings.ToUpper(str2))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
	pl()
	pl()

	// in Go chars are calls runes which are just the unicode (code points) repping
	// chars

	chars := "abcdefg💪🏾😁hijk" // interesting, adding color to runes adds byte size
	pl("Rune Count:", utf8.RuneCountInString(chars))
	pl("Byte Count:", len(chars))
	for i, v := range chars {
		fmt.Printf("%d : %U : %c : %d bytes = %d rune\n",
			i,
			v,
			v,
			len(string(v)), // len doesn't work on individual runes or bytes so
			// we have to convert rune to str
			utf8.RuneCountInString(string(v)))
	}
}
