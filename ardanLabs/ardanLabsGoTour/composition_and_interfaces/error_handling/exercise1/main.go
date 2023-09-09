package main

import "fmt"

/*
	Error Handling Design -
	Integrity matters and its a big part of the engineering process. At the heart
	of integrity is error handling. When it comes to Go, error handling is not an
	exception to be handled later or somewhere else in the code. Its a part of the
	main path and needs to be a main focus. This was a deliberate design decision by
	Pike, Griesemer, and Thompson to make errors values and a first class citizen in Go.

	Developers have the responsibility to return enough context about any error so a
	user can make an informed decision about how to proceed. Handling an error is about
	three things:
		1. logging the error
		2. not propagating the error any further
		3. determining if the Goroutine/program needs to be terminated.

	Again, in Go, errors are just values, so they can be anything you need them to
	be. They can maintain any state or behavior.
*/

// ERROR HANDLING BASICS -
// The error interface is built into the language. This is why it appears to be an
// unexported identifier.
type error interface {
	Error() string
}

// Any concrete value that implements this interface (o sea has an Error() method)
// can be used as an error value
// One important aspect of Go is that error handling is done in a decoupled state
// through this interface. A key reason for this is because error handling is an
// aspect of our application that is more susceptible to change and improvement.
// This interface is th etype Go applications must use as the return type for error
// handling. This is why funcs that return errors always use 'error' Its the interface.

// =====================================================================================
// This is the most commonly used concrete error value in Go programs. It's declared
// in the errors package from the standard library. Notice how the type is UNEXPORTED
// and it has one UNEXPORTED field which is a string. We can also see how pointer
// semantics are used to implement the error interface. This means that only addresses
// to values of this type can be shared and stored inside the interface. The method
// just returns the error string.
type errorString struct { // our concrete type that will implement the error interface
	s string
}

func (e *errorString) Error() string { // implementing the error interface
	return e.s
}

// Its important to note that the implementation of the Error method serves the purpose
// of implementing the interface and as well LOGGING. If any user needs to parse the
// string returned from this method, then we've already failed to provide the user with
// the right amount of context to make an informed decision.

// ======================================================================================
// The New function is how an error using the concrete type errorString is constructed.
// Notice how the function returns the error using the error interface. Also notice
// how pointer semantics are being used
func New(text string) error {
	return &errorString{text}
}

func main() {

	// Again, context is everything with errors. Each error must provide enough
	// context to allow the caller to make an informed decision about the state
	// of the goroutine/application. In this example, the webCall function returns
	// an error with the message Bad Request. In the main function here the call
	// is made to webCall and then a check is made to see if an error as occurred
	// with the call.
	if err := webCall(); err != nil {
		// Here we have the check for the error above. The key and something we see
		// throughout Go is to check is err!=nil. What this condition is asking is,
		// is there a concrete value stored inside the err interface value. When the
		// interface value is storing a concrete value, there is an error. In this
		// case, the context is literally just the fact that a concrete value exists,
		// its not important what the concrete value is.
		fmt.Println(err)
		return
	}
	fmt.Println("Life is good")

}

func webCall() error {
	return New("Bad Request")
}

// What about when it is important to know what error value exists??? This is where we
// use error variables, which we'll see in the next example.
