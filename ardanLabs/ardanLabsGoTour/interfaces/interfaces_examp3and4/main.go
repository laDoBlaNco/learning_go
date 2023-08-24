package main

import "fmt"

/*
	METHOD SET RULES:
	Our rules really come into effect when we are talking about the use of pointer vs value
	semantics as we spoke back when we created methods for the first time.
	
	Implementing an interface using pointer semantics applies some constraints on interface
	compliance

*/

// First we create our notifier interface with only one method set of behavior
type notifier interface{
	notify() 
}

// here we have our concrete type that will implement our notifier interface
type user struct{
	name string
	email string
}

// NOTE here that we set up the method, but we use pointer semantics for the method. Its still a
// notify method but its using pointer semantics. This is the point in which we call ourselves
// implementing an interface using pointer semantics. NOTE we are talking here about the IMPLEMENTATION
// and not the declaration of our interface.
func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email) 
}

// so here we create our function which requires a 'notifier' or in other words a create value that
// has the notifier behavior.
func sendNotification(n notifier){
	n.notify() 
}

func main(){
	u := &user{"Ladoblanco","ladoblanco@email.com"} // (I added the & after writing below)
	sendNotification(u) // but this fails
	// interfaces_examp3/main.go:42:19: cannot use u (variable of type user) as notifier value 
	// in argument to sendNotification: user does not implement notifier (method notify has 
	// pointer receiver)
	
	// this might be confusing as we see that it does in fact implement the interface, so what is
	// happening here???
	
	// This happens because there is a special set of rules in the specification about method
	// sets. These rules define what methods are attached to values and pointers of a type. 
	// They are in place to maintain the highest level of integrity in our programs:
	// 		- For any VALUE of type T, only those methods implemented with a VALUE receiver
	// 		  for that type belong to the method set of that value. 
	//		- For any ADDRESS of type T (pointer semantics), ALL methods implemented for that
	// 		  type belong to the method set for that value.
	// This is so simple, Not sure why I didn't understand this before. So in my own words, if 
	// the construct var uses value semantics then only methods with value receivers will 
	// be considered part of its method set, or said to be implementing the interface. But if 
	// we use a pointer (address of type T) then the method can be value or pointer receiver
	// and it'll work.
	
	// In Bill's words, when working with an address (pointer), all methods implemented are 
	// attached and available to be called. When working with a value, only those methods 
	// implemented with value receivers are attached and available to be called. Pointer receivers
	// don't belong to the method set of a value, period.
	
	// We saw this in a previous lesson about methods, were we were able to call a method 
	// against a concrete piece of data regardless of the data semantics declared by the
	// reciever. This is because the compiler can adjust to make the call. In this case, a
	// value is being stored inside an interface and the methods must exist. No adjustments
	// can be made. But why can't methods implemented with pointer receivers be attached to 
	// values of type T? What is the integrity issue?
	
	// ONE REASON is because we can't guarantee that EVERY value of a type T is going to be
	// addressable. If a value doesn't have an address, it can't be shared and thus can't have
	// a valid pointer. Example below: 
	
	// duration(42).notify() 
	// interfaces_examp3/main.go:78:15: cannot call pointer method notify on duration  
	// because it can't take the address of duration(42). In this exaample the value of 42
	// is a constant of KIND int. Even though the value is converted into a value of type
	// duration, it's not being stored inside a variable. This means the 'value' is never on
	// the stack or the heap. There isn't an address. Constanst only live at compile time.
	
	// but if I do the following will it work??
	test := duration(42) // Yessssss it works. So that's when something isn't addressable
	// when we don't put it in a var, we just convert the constant for immediate use. Here we 
	// ;ass the constant to duration and convert it, meaning its on the stack and has an address.
	test.notify() 
	
	// THE SECOND REASON and the bigger of the two, is that the compiler is telling me that
	// I'm not allowed to sue value semantics if I've already chosen to use pointer semantics. 
	// In other words I'm being forced to share the value with the interface since its not safe
	// to make a copy of a value that a pointer points to. If I chose to implement the method
	// with pointer semantics, then I'm stating that a value of this type isn't safe to be
	// copied. 
	
	// So the answer to our misterious error is simply to use pointer semantics, not so change
	// the method to use value semantics. 
	
}

type duration int

func (d *duration)notify(){
	fmt.Println("Sending notification in",*d) 
}
