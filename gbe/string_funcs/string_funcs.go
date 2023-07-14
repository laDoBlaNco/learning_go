package main

// The std strings package provides many useful string-related funcs. Here
// are some examples to give you a sense of the package

import (
	"fmt"
	"strings"
)

// we alias fmt.Println to a shorter name as we'll use it a lot below
var p = fmt.Println

func main() {

	// Here's a sample of the funcs available in strings. Since these
	// are functions from the package, not methods on the string object
	// itself, we need to pass the string in question as the first arg
	// to the function. You can find more functions in the 'strings'
	// pkg docs.

	p("Contains:	", strings.Contains("test", "es"))
	p("Count:		", strings.Count("test", "t"))
	p("HasPrefix:	",strings.HasPrefix("test","te")) 
	p("HasSuffix:	",strings.HasSuffix("test","st"))
	p("Index:		",strings.Index("test","e")) 
	p("Join:		",strings.Join([]string{"a","b"}, "-")) 
	p("Repeat:		",strings.Repeat("a",5)) 
	p("Replace:		",strings.Replace("foo","o","0",-1)) 
	p("Replace:		",strings.Replace("foo","o","0",1)) 
	p("Split:		",strings.Split("a-b-c-d-e","-")) 
	p("ToLower:		",strings.ToLower("TEST")) 
	p("ToUpper:		",strings.ToUpper("test")) 
}
