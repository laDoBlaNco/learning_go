package main

import (
	"fmt"
)

// This example is not necessarily specific for slice mechanics, but because variadic
// functions use slice mechanics to work,we will put it here. In a nutshell, when we
// have a func that will take an arbiratry amount of args, we can tell Go how to
// treat the args inside the function. We may want to do 1 of 2 things:
// 		- give the args as we would in any other function separated by commas or
// 		- provide the args in a slice and tell Go to treat them as a slice of args inside
// 		  the body of the func

// Let's start with a struct type that declares user information
type user struct {
	id   int
	name string
}

func main() {

	// Declare and initialize a value of type user
	u1 := user{
		id:   1432,
		name: "Betty",
	}

	// Declare and initialize another value of type user
	u2 := user{
		id:   4367,
		name: "Janet",
	}

	// display both user values with a custom func
	display(u1, u2)

	// Create a slice of user values
	u3 := []user{
		{24, "Bill"},
		{32, "Joan"},
	}

	// display all the user value from the slice
	display(u3...) // note the '...' after our slice in the same display func

	// Here we use another custom func to make a change and show the backing array being shared
	change(u3...) // again note the '...' on our slice value
	fmt.Println("********************************************")
	for _, u := range u3 {
		fmt.Printf("%+v\n", u)
	}
}

// Now let's create our display function that can accept and display multple users
func display(users ...user) { // here is the key. NOTE the '...' in front of the type. Here we
	// are telling go that users will be a arbitrary amount of data of the user type and inside it should
	// be used as a slice. This way when we put display(u1,u2,u3) inside Go create a slice of
	// []users{u1,u2,u3} and when we put display(u3...) we are in effet telling go that u3 is already
	// a slice so just pass it in and use it or we could think of Go as breaking it down and turning
	// it into display(u3,u4)
	fmt.Println("********************************************")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// or custom function 'change' shows how the backing array is shared, also note the '...user'
func change(users ...user) {
	users[1] = user{99, "Some Backing Array"}
}

// Everything is the same mechanics we've already studied about regarding working with slices
// The only thing we need to remember is how its designed within variadic functions using '...'
// 		1. Before the type means that this is a variadic function and all args after a certain
// 		   point should be grouped together in a slice to be used in the func body (we say up to
// 		   a certain point because we can have normal args/types prior to the variadic_arg ...type)
// 		2. After the arg means that the arg is already a slice of the needed type and Go can pass it 
// 		   as is, or if we want to think of it as unload that slice in the arg list and proceed
// 		   as stated above in #1.


