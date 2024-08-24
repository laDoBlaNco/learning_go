/*
So as mentioned, to really bring this home, when two interface values
are compared, it's the concrete data inside of them and not the interface
values themselves that are compared.
*/
package main

import "fmt"

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func New(text string) error {
	return errorString{text}
}

var ErrBadRequest = New("Bad Request")

func main() {
	err := webCall()
	if err == ErrBadRequest {
		fmt.Println("Interface values MATCH")
	}
}

func webCall() error {
	return New("Bad Request")
}

/*
	From lines 10-20 we copied the implemenation of the default error type in Go
	from the errors pkg with one change. Our implemenation of the New function
	on line 18 is using value semantics instead of pointer semantics. So its
	storing an errorString value inside the error interface value being returned
	and not its address.

	Then on line 22, an error interface variable is declared for the "Bad Request"
	error. Skipping down to line 31, the webCall func is returning a new error
	interface value with the same message "Bad Request". Then in the main
	func on line 24, the webCall func is called and the returned error interface
	value is compared with the error interface variable. And its the VALUES that
	are inside the interface that are being compared

	Even if they are different interface values, if they have the same concrete
	data they will be considered equal. For example different error interfaces
	with the same concrete value inside. The data inside the interface is what is
	being compared, not the interface itself. When using pointer semantics, addresses
	are being compared. When using value semantics, values are being compared.
	Interface values are valueless and it's always about the concrete data stored
	inside of them.

	CONCLUSION
	This post represented another example of how value and pointer semantics play
	a significant role in writing code in Go. The interface can store its own copy
	of a value (value semantics) or a copy of an address (pointer semantics). We
	wanted to show how the method set rules are providing a level of integrity
	checking by not allowing a change in semantic from pointer to value. This
	promotes the idea that it is not safe to make a copy of the value that is pointed
	to by a pointer. This mix of semantics must be taken seriously.

	As we continue to write code in Go, we need to look at the semantics we are
	using for any given type. During code reviews, looking for consistency in
	semantics for data of a given type and question code that violates that
	consistency. There are exceptions to every rule, but they must be made as a
	concious choice. 

*/
