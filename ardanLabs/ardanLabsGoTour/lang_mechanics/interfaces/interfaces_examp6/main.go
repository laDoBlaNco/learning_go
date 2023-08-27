package main

import(
	"fmt"
	"log"
)

// This sample program is to show the syntax of type assertions which is something we
// use a lot with interfaces. Since interfaces are valueless, we at times need to know
// the underlying type is being held in the interface.

type user struct{
	id int
	name string
}

// finder interface reps the ability to find users
type finder interface{
	find(id int)(*user,error) 
}

// a service for dealing with users
type userSVC struct{
	host string
}

// find implements the finder interface using pointer semantics
func(*userSVC) find(id int)(*user,error){
	return &user{id:id,name:"Anna Walker"},nil
}

func main(){
	svc := userSVC{
		host:"localhost:3434",
	}
	
	if err:=run(&svc);err!=nil{
		log.Fatal(err) 
	}
}

// run performs the find operation against the concrete data that is passed into the call
func run(f finder)error{
	u,err := f.find(1234) 
	if err!=nil{
		return err
	}
	fmt.Printf("Found user %+v\n",u) 
	
	// Ideally the finder abstraction would encompass all of the behavior we care about
	// But what if for some reason we really need to get to the concrete value stored
	// inside the interface?
	// Can we access the "host" field from the concrete userSVC type pointer that is stored
	// inside this interface var? No, not directly.
	// All we know is the data has a method named 'find'. 
	// log.Println("queried",f.host) 
	// f.host undefined (type finder has no field or method host
	
 	// we can though use a type assertion to get a copy of the userSVC pointer that is stored
 	// inside the interface.
 	svc := f.(*userSVC) 
 	log.Println("queried",svc.host) 
 	
 	return nil
}
