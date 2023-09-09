package main

import (
	"fmt"
	"reflect"
)

/*
	So what if an error variable isn't enough context for our user?? What if some special
	state needs to be checked, like with networking errors?? In these cases, we can use
	custom concrete error types. Some tutorials and books start with these but we should
	really only use them if the context need calls for it. If existance and error variables
	aren't enough.
*/

// First we have a custom concrete error type implemented in the json package. Notice
// the name as a suffix of Error in the name of the type. Also notice the use of pointer
// semantics for the implementation of the error interface. Once again the implementation
// is for logging and shoudl display information about all the fields being captured.
type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}

// Then we have a second custom concrete error type found in the json package. The
// implementation of this Error method is a bit more complex, but once again just
// for logging and using pointer semantics
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer) " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

// user is a type for use in the Unmarshal call.
type user struct {
	Name int
}

func main() {

	// The context of the error here is more about the type of error stored inside
	// the error interface. There needs to be a way to determine that
	var u user

	// A generic type assertion within the scope of the switch statement is how we test
	// what type of value is being stored inside the err interface value. Type is the 
	// context here and now we can test and take action with access to all the states
	// of the error. 
	
	err := Unmarshal([]byte(`{"name":"kevin"}`), u)
	if err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n",
				e.Value,
				e.Type,
			)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		return
	}
	fmt.Println("Name:", u.Name)
}

// Here is a portion of the Unmarshal function. Notice how it constructs the concrete
// error values in the return, passing them back to the caller through the error
// interface. Pointer semantic construction is being used because pointer semantics
// were used in the declaration of the Error method.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
// There is one problem here though. WE are no longer decoupled from the concrete
// error value. This means if the concrete value is changed, then our code can 
// break. The beautiful part of using an interface for error handling is being 
// decoupled from breaking changes. 

// Let's take a look at this in the next example.
