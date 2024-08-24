package main

import(
	"fmt"
	"runtime"
)

// This sample shows how to recover from panics, though I don't know why we would
// want to do this.

func main(){
	
	// call the testPanic function to run the test
	if err := testPanic();err!=nil{
		fmt.Println("Error:",err) 
	}
	
}

// testPanic simulates a function that encounters a panic to test our catchPanic function
func testPanic()(err error){
	
	// schedule the catchPanic function to be called when the testPanic function returns
	defer catchPanic(&err) 
	
	fmt.Println("Start test")
	
	// Mimic a traditional error from a function
	err = mimicError("1") 
	
	// Trying to dereference a nil pointer will cause the runtim to panic
	// NOTE that panics are what we call runtime errors in other langs
	var p *int // this is a zero default (nil) pointer
	*p = 10 // and we are trying to access (dereference) a nil pointer
	
	fmt.Println("End test")
	return err 
}

// catchPanic catches panics and processes the error
func catchPanic(err *error){
	
	// Check if a panic occured
	if r:=recover();r!=nil{
		fmt.Println("PANIC Deffered")
		
		// Capture the stack trace
		buf := make([]byte,10000)
		runtime.Stack(buf,false)
		fmt.Println("Stack Trace:",string(buf))
		
		// if the caller wants the error back provide it
		if err!=nil{
			*err = fmt.Errorf("%v",r) 
		}
	}
}

// mimicError is a function that simulates an error for testing the code
func mimicError(key string)error{
	return fmt.Errorf("Mimic Error: %s",key)  
}
