// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// In this example we'll take a look at another pointer to share data, but first
// Let's take a dive into the Escape Analysis of Go and the HEAP. 
//
// Let's review this sample program to show the basic concept of using a pointer
// to share data, but in a struct
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name   string
	email  string
	logins int
}

func main() {

	// Declare and initialize a variable named bill of type user.
	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//** We don't need to include all the fields when specifying field
	// names with a struct literal, anything left blank will be given its zero default
	// similar to us just calling user{}

	// Pass the "address of" the bill value, again noted by the use of '&'
	display(&bill)

	// Pass the "address of" the logins field from within the bill value.
	increment(&bill.logins) // so this is our pointer to logins (*logins) which the
	// func increment calls for

	// Pass the "address of" the bill value.
	display(&bill) // and this is our pointer to a user (*user) which is what the 
	// func display calls for
}

// increment declares logins as a pointer variable whose value is
// always an address and points to values of type int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

// display declares u as user pointer variable whose value is always an address
// and points to values of type user.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
}
