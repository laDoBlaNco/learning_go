package main

import(
	"fmt"
	"github.com/ladoblanco/ardanLabsGoTour/exporting/examples"
)

/*
	In this last example, even though the user type is unexported, it has two exported
	fields. This means that when the user type is embedded in the exported Manager type,
	the user fields that are exported are promoted and accessible. Its common to have types
	that are unexported with exported fields because the reflection package can only operate
	on exported fields. Marshallers won't work otherwise.
	
	The example does though create a bad situation where code outside of the package can 
	construct a Manager, but since the embedded type user is unexported, the fields for those
	can't be initialized. This creates a partial construction problem that will lead to
	bugs later. We need to be consistent with exporting and unexporting.
*/


func main(){
	
	u := examples.Manager{
		Title: "Dev Manager",
	}
	
	// we can now set the unexported fields off our new var, but they weren't initialized since 
	// user2 is unexported.
	u.Name = "Kevin"
	u.ID = 46
	
	fmt.Printf("User: %#v\n",u) 
}

/*
	IN SUMMARY:
		- Code in Go is compiled into packages and then linked together
		- Identifiers are exported (or remain unexported) based on the letter-case
		- We import packages to  access exported identifiers
		- Any package can use a value fo an unexported type, but this is annoying to use. 


*/
