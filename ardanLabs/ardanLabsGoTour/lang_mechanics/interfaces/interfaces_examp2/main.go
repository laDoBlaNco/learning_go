package main

import "fmt"

/*
	"POLYMORPHISM means that a piece of code changes its behavior depending on the CONCRETE data
	its operating on."" This was said by Tom Kurtz, the inventor of BASIC. This is the ideal
	definition of polymorphishm and what we will use moving forward.

	Look at our new function below 'retrieve'

*/

// first we have our interface 'reader' again
type reader interface {
	read(b []byte) (int, error)
}

// then we have our concrete type 'file'
type file struct {
	name string
}

// and we have our 'read' function to implement the reader interface, again NOTE that this is
// a specific implementation of what we want our read behavior to be with this 'file' type
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channe></rss>"
	copy(b, s)
	return len(s), nil
}

// we also have our concrete type 'pipe'
type pipe struct {
	name string
}

// and we have our 'read' function to implement the reader interface with its own specific
// implementation behavior for this 'pipe' type
func (pipe) read(b []byte) (int, error) {
	s := `{name: "ladoblanco", title: "go dev"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	// Let's create our two concrete values again for both file and type
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// now we call our ONE function that works on both types
	retrieve(f)
	retrieve(p)
}

// Now take a look at the type of data this function accepts. Its wants a 'value' of type reader
// This is of course impossible cuz we just said that reader is an interface and interfaces
// are valueless types. So it can't be asking for a reader value as they don't exist. So what is it
// really asking me for?
func retrieve(r reader) error { // give me a 'reader' value???
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
// If the function is not asking for a reader value then what is it asking for??? IT IS ASKING FOR
// THE ONLY THINK IT CAN ASK FOR, CONCRETE DATA.  The function retrieve is a POLYMORPHIC function
// because it's asking for concrete data not based on what the data is (its value or type) but rather
// based on what the data can do (interface type) 

// We constructed two concrete values, one of type file and one of type pipe. Then we passed a copy
// of each value to the polymorphic function. This is because each of these values implement
// the full method set of behavior defined by the reader interface. 

// When the concrete value 'file' is passed into 'retrieve', the value is stored inside of a 
// two word internal type or structure (just like string, slices, refrence types, etc)
// This two word struture has its first word points to a special data structure called an iTable
// and its second word points to the value being stored. In this case, its a copy of the file
// value since value semantics are at play

// An iTable serves 2  purposes
// 		- It describes the type of value being stored. In our case, its a file value
// 		- It provides function pointers to the concrete implementations of the method set for
// 		  the type of value being stored. O sea it determines what the value is in that 2nd 
// 		  word pointer and depending on what the value is, the iTable will point to the method
// 		  that goes with that value
// So for example when the read call is made against the interface value, an iTable lookup is made
// to find the concrete implementation of the read method associated with the type. Then the method
// call is made against the value being stored in the second word.

// We can say that 'retrieve' is a polymorphic function because the concrete value pipe can be
// passed into retrieve and now the call to read against the interface value changes its behavior.
// This time that call to read is reading a network instead of a file.


// So now we need to look at some of the rules that govern these interfaces and their method sets.

