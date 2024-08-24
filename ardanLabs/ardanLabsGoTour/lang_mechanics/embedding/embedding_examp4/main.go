package main

import "fmt"

type notifier interface{
	notify() 
}

type user struct{
	name string
	email string
}

func(u *user) notify(){
	fmt.Printf("Sending user email to %s <%s>\n",u.name,u.email) 
}

type admin struct{
	user // embedded type with value semantics
	level string
}

// here we have the same method implemented on our outer type
func (a *admin)notify(){
	fmt.Printf("Sending admin email to %s <%s>\n",a.name,a.email) 
}

func main(){
	
	ad := admin{
		user:user{
			name:"Kelen Delight",
			email:"whitesidekelen@gmail.com",
		},
		level:"super",
	}
	
	// NOTE here that since the outer type already has a notify method, the inner type's 
	// method isn't promoted. 
	sendNotification(&ad)
	
	// we can still get to the inner method directly
	ad.user.notify() // NOTE that this one is 'user email' as opposed to the other 'admin email'
	
	// but the inner type's method is NOT promoted to the outer type in this case
	ad.notify()  
}

func sendNotification(n notifier){
	n.notify() 
}

/*
	SUMMARY:
		* Embedding types alow us to share state or behavior between types
		* The inner type never loses its identity
		* This is not inheritance
		* Through promotion, inner type fields and methods can be accessed through the outer type
		* The outer type can override the inner type's behavior

*/
