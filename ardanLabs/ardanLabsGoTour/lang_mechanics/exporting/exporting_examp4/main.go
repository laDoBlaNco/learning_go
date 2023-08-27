package main

import(
	"fmt"
	
	"github.com/ladoblanco/ardanLabsGoTour/exporting/examples" 
)

// here we see some of the same but with exported and unexported fields of a struct.
// When it comes to  fields in a struct, the first letter again declares if the field is
// accesible outside of the package it's declared in. In this case, Name and ID are accessible
// but password is not. Its an idiom in Go to separate exported and unexported fields in this 
// manner below if this is a reasonable or practical thing to do. normally all fields would 
// be one right after the other, as we've seen everywhere else. 

func main(){
	
	// using our exported User type from our examples package
	u := examples.User{
		Name: "Kevin Whiteside",
		ID: 46,
		
		password: "xxxxxx",
		// exporting_examp4/main.go:23:3: unknown field password in struct literal of type examples.User
		// the error is due to password not being exported, so I don't have access to it
	}
	
	fmt.Printf("User: %#v\n",u) 
}

