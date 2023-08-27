/*
	Embedding types provide the final piece of sharign and reusing state and behavior
	between types. Through the use of INNER TYPE PROMOTION, an inner type's fields 
	and methods can be directly accessed by references of the outer type.
	
	EMBEDDING MECHANICS
	The first example we are looking at doesn't show embedding, just the declaration
	of two struct types working together as a field from one  type to the other. 
	NOTE: this is NOT considered embedding
*/

package main

import "fmt"

// first we have our user
type user struct{
	name string
	email string
}

// adding a notify method to our user type
func(u *user)notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email) 
}

// now let's create a admin
type admin struct{
	person user // we put a person type user in our admin struct THIS IS NOT CONSIDERED EMBEDDING
	level string
}

func main(){
	// Let's create out admin
	ad := admin{
		person:user{
			name:"Kevin Whiteside",
			email:"whitesidekevin@gmail.com",
		},
		level:"super",
	}
	
	// then we access the fields methods for our person
	ad.person.notify() 
}

// So the key here is that when we see an field that has a custom struct type as its value
// this IS NOT EMBEDDING, but simply composing the types together to use the methods and fields
// with each other.

// Now in the next example we'll do the same but with actual embedding. 
