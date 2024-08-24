package main

import "fmt"

/*
	Now in this example, we'll do the same as before but adding in an interface and
	a polymorphic function that accepts any concrete value that implements the full method
	set of behavior defined by our notifier interface. In this case that's just 'notify'
	
	NOTE that because of embedding and promotion, values of type 'admin' now implement
	the notifier interface as well. So admin is aa 'notifier' automagically 
*/

type notifier interface{
	notify() // our method set of behavior
}

type user struct{
	name string
	email string
}

func(u *user)notify(){ // with this we implement our interface using pointer semantics
	fmt.Printf("Sending user email to %s <%s>\n",u.name,u.email) 
}

type admin struct{
	user // embedded type using value semantics
	level string
}

func main(){
	
	ad := admin{
		user:user{
			name:"Xavier Whiteside",
			email:"xavieroscar09@gmail.com",
		},
		level:"super",
	}
	
	sendNotification(&ad) 
}

// our polymorphic function 'sendNotification' accepts addres of values that implement the notifier
// interface and sends notifications
func sendNotification(n notifier){
	n.notify() 
}

// NOTE: what if the outter and inner types both have the same method implemented?
// let's see that in our last example.
