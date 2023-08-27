package main

import "fmt"

/*
	From the Go spec:
	A short variable declaration mya redeclare variables provided they were
	originally declared earlier in the same block with the same type, and at least one
	of the non-blank variables is new.
	
	This isn't shadowing though as I thought. Here we are talking about redeclaring 'err' for
	example in 
	a,err := somefunc()
	b,err := someOtherfunc()
	Here err is being redeclared but since a and b are new, we are ok.
	
	Below we have a sample program to show some fo the mechancis behind the short declaration
	operator redeclaring

*/

// user is a struct type that declares user information
type user struct{
	id int
	name string
}

func main(){
	
	// Declare the error variable
	var err1 error
	
	// the short var declaration operator will declare u and redeclare err1 since u is new 
	// in this declaration
	// without the 'u' below we get the following error:
	// 'no new variables on left side of :=''
	u,err1:=getUser()
	if err1!=nil{
		return
	}
	
	fmt.Println(u) 
	
}

// getUser returns a pointer of type user.
func getUser()(*user,error){ // takes no args and returns a *user and an error
	return  &user{1432,"Betty"},nil
}
