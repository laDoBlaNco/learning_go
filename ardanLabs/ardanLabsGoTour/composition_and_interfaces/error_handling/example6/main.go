package main

import (
	"errors"
	"fmt"
)

/*
	Handling Errors:
	Handling errors is more of a macro level engineering conversation. In some worlds, error
	handling means the error stops with the function handling the error, error is logged with
	full context, and the error is checked for its severity. Based on the severity and ability
	to recover, a decision to recover, move on, or shutdown is made.

	One problem is that not all functions can handle and error. One reason could be because not
	all functions are allowed to log. So what happens then when an error is being passed back
	up the call stack and can't be handled by the function receiving it?? An error needs to
	be wrapped in context so the function that eventually handles it, can properly do so.

*/

type AppError struct {
	State int
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", ae.State)
}
 
func IsAppError(err error) bool {
	var ae *AppError
	return errors.As(err, &ae)
}

func GetAppError(err error) *AppError{
	var ae *AppError
	if !errors.As(err,&ae){
		return nil 
	}
	return ae 
}

func main(){
	if err:=firstCall(10);err!=nil{
		
		// check if the error is an AppError
		if IsAppError(err){
			ae:=GetAppError(err) 
			fmt.Printf("Is AppError, State: %d\n",ae.State) 
		}
		
		fmt.Print("\n***********************************\n\n") 
		
		// display the error using the implementation of 
		// the error interface.
		fmt.Printf("%v\n",err) 
	}
}

func firstCall(i int)error{
	if err:=secondCall(i);err!=nil{
		return fmt.Errorf("secondCall(%d) : %w",i,err) 
	}
	return nil
}

func secondCall(i int)error{
	return &AppError{99} 
}

/*

	NOTE:
	- Use the default error value for static and simple  formatted messages
	- Create and return error variables to help the caller identify specific errors
	- Create custom error types when the context of the error is more complex and we need
	  information over the state
	- Error values in Go aren't special, they are just values like any other, and so we
	  have the entire language at our disposal. 
	  
	Quotes:
	"Systems cannot be developed assuming that human beings will be able to write millions of 
	code without making mistakes, and debugging along is not an efficient way to develop reliable
	systems." - Al Aho (inventor of AWK) 
	
*/
