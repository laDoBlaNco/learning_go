package main

import (
	"fmt"
	"log"
)

func main() {

	// the overall questions to answer:
	// If a struct and an embedded field both implemented teh same interface
	// 1. Would the compiler throw an error because we now have two implementations
	// 	  of the interface?
	// 2. If the compiler accepted the type declaration, how does the compiler
	//    determine which implementation to use for interface calls?

	// Once we learn the mechanics behind methods, interfaces, and embedded types,
	// the answers are obvious. Let's start with methods:

	// METHODS:
	// Go has functions and methods. In Go, a method is a function that is declared
	// with a receiver. A receiver is a VALUE or a POINTER of a named or struct
	// type. All methods for a given type belong to the type's method set.

	// Below we created a struct type User and a method that receives a value
	// of type User (like self or this in other langs).

	// Value of type User can be used to call the method with a value receiver:
	bill := User{"Bill", "bill@email.com"}
	bill.Notify() // when we say receive it means we need a User.Notify() to work

	// Pointer of type User can also be used to call a method with a val receiver
	// as I learned before.
	jill := &User{"Jill", "jill@email.com"}
	jill.Notify() // and it'll also just work as Go will take care of the syntax
	// Go adjusts and dereferences the pointer so the call can be made. When a
	// receiver is not a pointer, the method is operating against a copy of the
	// receiver value.
	// We can change the receiver below to a pointer and it'll work the same without
	// changes above.

	// INTERFACES:
	// Are  special and provide an incredible amount of flexibility and abstraction
	// in Go. They are the way of specifying that values jand pointers of a
	// particular type can behave in a specific way. Its a type that specifies a
	// method set and all the methods for an interface type are considered to be
	// the interface.

	// Below we declare an interface. Its a Go convention to name interfaces
	// with an -er suffix. This is why we have things like "Reader","Writer",
	// "Stringer", etc. We can specify as long a list as we want, but its good
	// to follow the Go team's lead as you will almost never find an interface
	// with more than two methods in its method list. Rob Pike said in one of his
	// talks that io.Writer, io.Reader, and interface{} are the 3 most important
	// interfaces in the entire Go eco-system and they have an average of 2/3s of
	// a method. O sea Writer has 1, Reader has 1 and interface{} is empty. So they
	// are more useful than any toher interface in all of Go.

	// Let's implement an interface...Go is unique  when it comes to how to implement
	// an interface we want our types to support because it just happens. Go doesn't
	// require any explicit state that our Type has to implement. If every method
	// that belongs to our interface's method set/job description is in our type,
	// then our type implements the interface, sencillo y period.
	
	// Calling our interface func:
	user:=&User{ // here we give the address (pointer) of User to pass interface compliance.
		Name:"janet jones",
		Email:"janet@email.com",
	}
	fmt.Println()
	fmt.Println("Called with our User type pointer:")
	SendNotification(user) // This answered my question below. I didn't have to 
	// set user to an instance of my interface. As long as my type User has the right
	// method it can be used by my func as the interface.
	
	// Just using SendNotification(user) gives us an error 'User doesn't implement
	// Notifier (Notify method has pointer receiver)' Go knows and tells us the 
	// specific problem. Why does the compiler not conisder our value to be of type
	// that implements the interface? The rules for determining interface compliance
	// are based on the receiver for those methods and how the interface call is 
	// being made. 
	// 1. The method set of the corresponding pointer type *T is the set of all
	//    methods with receiver *T or T. -- so if the interface var we are using to
	//    call an interface contains a pointer then methods with receivers based on
	//    both values and pointers will satisfy the interface. This doesn't apply
	//    for our example because we are passing a VALUE
	//
	// 2. The method set of any other type T consists of all methods with receiver
	//    type T. -- This is stating that if the interface var we are using to call
	//    an interface method contains a value, then only methods with value 
	//    receivers will satisfy the interface. This rule doesn't apply for our
	//    example cuz our receiver of Notify accepts a pointer. 
	// There are only two rules in the spec for interface compliance, so our derived
	// rule that applies to our example is:
	// 3. The method set of the corresponding type T DOES NOT consist of any methods
	//    with receiver type *T -- So in our case and why we are getting an error
	//    is that Notify method is using a pointer for the receiver and we are using
	//    a value to make the interface method call. To fix this we just need to 
	//    pass the address (pointer) of the User value to the function. So really
	//    for me in the end its simply that Go doesn't do the same adjustment for
	//    interfaces that it does for methods. Where  Go will adjust the dereferencing
	//    etc on a method, you need to do it manually on interfaces method calls 
	//    just as you would with any other Func. Remembering the difference I deduced
	//    earlier between methods and Funcs.
	
	// EMBEDDED TYPES:
	// So struct types have the ability to contain anony or embedded fields. This
	// is also called 'embedding a type'. When we embed a type into a struct, the
	// name of that type acts as the field name for what is then an embedded field.
	// We see this below with a new type creation and embedding our User into it.
	
	// Create a value of type Admin and pass its address to our func:
	admin:=&Admin{
		User:User{
			Name:"john smith",
			Email:"john@email.com",
		},
		Level:"super",
	}
	fmt.Println()
	fmt.Println("Called from embedded type of our Admin type:")
	SendNotification(admin) // And it works. We are call our function with a pointer
	// type Admin. Thanks to composition, the Admin type now implements the 
	// interface we need for our SendNotification func through the promotion of the
	// methods from the EMBEDDED User type. So now where are these fields in relation
	// to the struct? 
	
	// Effective Go says: "When we embed a type, the methods of that type become
	// methods of the outer type (Admin), but when they are invoked, the receiver of the
	// method is the inner type (User), not the outer one." -- this means that the name
	// of the embedded type acts as the field name and the embedded type exists
	// as an inner type, we can make calls through that inner type
	fmt.Println()
	fmt.Println("Here we invoke directly using the inner type name:")
	admin.User.Notify() // which also works. but... 
	
	// they are also promoted to the outer type:
	fmt.Println()
	fmt.Println("Here we use the outer type name as the methods are promoted:") 
	admin.Notify() 

	// Here are the rules for inner type method set promotion in Go:
	// Given a struct type S and a type name T, promoted methods are included in the
	// method set of the struct as follows:
	// 1. If S (Admin) contains an anony field T (User), the method sets of S and *S 
	//    both include promoted methods with receiver T.--
	//    This rule is stating that when we embed a type, methods for the embedded
	//    type with receivers that use a VALUE are promoted and available for calling
	//    by VALUES and POINTERS of the outer type.
	//
	// 2. The method set of *S (Admin) also includes promoted methods with receiver
	//    *T (User).--
	//    This rule tells us that when we embed a type, methods for the embedded
	//    type with receivers that use a pointer are promoted and available for 
	//    calling by POINTERS of the outer type.
	//
	// 3. If S (Admin) contains anony field *T (User), the method sets of S and *S
	//    both include promoted methods with receiver T or *T--
	//    This rule states that when we embed a pointer of the type, methods for the
	//    embedded type with the receivers that use both values and pointers are
	//    promoted and available for calling by values and pointers of the outer
	//    type.
	//
	// Again since there are only 3 rules stated, we've come up with a 4th:
	// 4. If S (Admin) contains  an anony field T (User), the method set of S
	//    DOES NOT include promoted methods with receiver *T. 
	//    This rule is stating that when we embed a type, methods for the embedded
	//    type with receivers that use a POINTER are not promoted for calling by
	//    VALUES of the outer type. This is consistent with the rules for interace
	//    compliance stated above. 
	//    
	
	
	
	
	// So now we an answer (or Bill can answer) our questions:
	// Below we create the method that displays an message on behalf of Admin.
	// Without any other changes in main we now see that Go is able to determine
	// which of the implementations to use. When we pass Admin to our SendNotification
	// func or use admin.Notify() we get the messags set for admin. And the others
	// when using User. The User's type implementation is no longer promoted to the
	// outer type. 
	
	// So, 
	// 1. Would the compiler throw an error because we have two implementations
	// of the same interface? -- No, cuz when we use an embedded type, the unqualified
	// type's name acts as the field name. This has the effect of fields and methods
	// of the embedded type having unique  name as an inner type of the struct. 
	// So we can have an inner and outer implemenation of the same interface with
	// each being unique and accessible.
	//
	// 2. If the compiler accepted the type declaration, how does it know which
	// implementation to use for the interface calls? -- If the outer type contains
	// an implementation that satisfies the interface, it'll be used, OTHERWISE, 
	// thanks to the method promotion, any inner type that implements jthe interface
	// and follows the rules for method set promotion can be used through the 
	// outer type.
	
	// CONCLUSION:
	// The way that methods, interfaces, and embedded tyeps work together is 
	// something very unique to Go. These features help us create  powerful constructs
	// to achieve the same ends as OOP code without all the complexity. Remember 
	// ITS NOT INHERITANCE, ITS COMPOSITION. With the features that I've learned
	// from this post, I'll be able to build abstracted and scalable frameworks
	// with minimal amounts of code and confusion, thus killing the  idea that 
	// Go is too verbose.
	
	// A final thought from Bill (as if all of this wasn't his thoughts):
	
	// The more I learn about the details of the language and the compiler, the
	// more I come to appreciate how orthogonal the language is. Small features
	// that work together and allow us to be creative and use the language in ways
	// not even the language designers thought or dreamed about. I recommend to take
	// the time to learn the language features so you can do more with less and be both
	// creative and productive at the same time.
}

// Here we declare a struct  type and a method
type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email to %s<%s>\n",
		u.Name,
		u.Email)
		
	return nil
}

// Here we declare an interface:
type Notifier interface {
	Notify() error // this is the job description or method list with which the
	// interface identifies what Types can be said to be "Notifiers"
}

// this calls the Notify method that is implemented by the value or pointer that will
// be passed into the function. This func can be used to execute the specific behavior
// for any value or pointer of a given type that implements the interface.
func SendNotification(notify Notifier) error {
	return notify.Notify()
}

// This is the missing piece for me on interfaces. I've built interfaces and methods
// but never focused on creating func that take in interfaces. This means that I would
// need to pass in an instance of my Notifier interface which in turn can be any
// type that has a notify method implemented. So my question is can I enter a type
// in this func without specifically making an interface instance?
// Looks like that's Bill's next step as we flesh out the method above.

type Admin struct{// Note that we don't add a type (int, string, etc) to the embedded
	User
	Level string
}
// This is NOT INHERITANCE but composition. There is no relationship between the
// User and the Admin type.

// To finalize and answer our questions we need to implement the Notifier interface
// for our Admin type, which mainly just means creating a method so Admin can 
// directly be considered a Notifier.
func (a *Admin)Notify()error{
	log.Printf("Admin: Sending Admin Email to %s<%s>\n",
	a.Name,
	a.Email) 
	
	return nil
}
// The only difference is that this method displays an message on behalf of Admin.
// This will help us determine which implementation gets called, Since we now have 
// two of the same, one direct and one embedded.
