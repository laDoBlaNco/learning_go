package main

import (
	"errors"
	"fmt"
)

/*

	Create a custom error type called appError that contains three fields, err error,
	message string and code int. Implement the error interface providing your own message
	Implement a second method named temporary that returns false when the value of the
	code field is 9. Write a function called checkFlag that accepts a bool. if the value is
	false, return a pointer of your custom error type initialized as you like. If the value
	is true, return a default error. Write a main function to call the checkFlag function
	and check the error using the temporary interface.

*/

//  1. Declare a strut type named appError with tree fields, err of type error, message
//     of type string and code of type int.
type appError struct {
	err     error
	message string
	code    int
}

// 2. Declare method  for the appError struct type that implements the error interface
func (ae *appError) Error() string {
	return fmt.Sprintf("there was an app error %s: %s\nCode: %d\n",ae.err,ae.message,ae.code)
}

//  3. Declare a method for the appError type named Temporary that returns true when the value of
//     code field is not 9
func (ae *appError) Temporary() bool {
	return (ae.code != 9)
}

//  4. Declare the tempororary interface type wtih a method name Temporary that takes no params
//     and returns a bool
type temporary interface {
	Temporary() bool
}

// 5. Declare a function namec checkFlag that accepts a bool and returns an error interface value
func checkFlag(b bool) error {
	if !b {
		return &appError{errors.New("Flag False"), "The flag was false", 9}
	}

	// default error
	return errors.New("Flag True")
}

func main(){
	// 6. Call the checkFlag function to simulate an error of the concrete type
	if err:=checkFlag(false);err!=nil{
		// 7. Check the concrete type and handle appropriately
		switch e:=err.(type){
		case temporary:
			fmt.Println(err)
			if !e.Temporary(){
				fmt.Println("Critical Error!") 
			}
		default:
			fmt.Println(err)	
		}
	}
	
}
