// sample program to show type assertions using the comma,ok idiom
package main

import (
	"fmt"
	"log"
)

// first we have our user
type user struct{
	id int
	name string
}

// then we have a finder interface with a method list of 'find'
type finder interface{
	find(id int)(*user,error) // it'll take an in and return a pointer to a user
}

// we have now a struct type that will implement our interface
type userSVC struct{
	host string
}

// we have the find method implementing the finder interace on userSVC with pointer semantics
func(*userSVC)find(id int)(*user,error){
	return &user{id:id,name:"Kevin Whiteside"},nil 
}

// let's mock our service with mockSVC
type mockSVC struct{}

// Let's implement finder on our  mockSVC as well, also with pointer semantics
func(*mockSVC)find(id int)(*user,error){
	return &user{id:id, name:"Odalis Whiteside"},nil
}

func main(){
	// create our var for mockSVC
	// var svc mockSVC
	var svc2 userSVC // I added this to test a change in the result and it worked. type assertion
	// was able to tell that I changed the  underlying type and thus the resulting method that 
	// was run changed. 
	
	if err := run(&svc2);err!=nil{
		log.Fatal(err) 
	}
}

func run(f finder)error{ // our run func is going to take a finder and return a possible error
	// and as long as our type implements a find method, its a finder
	u,err := f.find(1234) // and to be a finder you have to return a *user, error 
	if err!=nil{
		return err
	}
	fmt.Printf("Found user %+v\n",u) 
	
	// If the concrete type value stored inside the interface value is of the type userSVC
	// then "ok" will be true and "svc' will be a copy of the pointer stored inside the
	// interface 
	if svc,ok :=f.(*userSVC);ok{
		log.Println("queried",svc.host) 
	}  
	
	return nil 
} 


