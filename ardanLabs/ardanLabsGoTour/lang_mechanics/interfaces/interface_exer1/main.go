/*
Declare an interface named speaker with a method named speak.
Declare a struct named english that represents a person who speaks english
Declare a struct named spanish that represents a person who speaks spanish
Implement the speaker interface for each struc using a value receiver and these
literal strings "Hello World" and "你好世界"
Declare a variable of type speaker and assign the address of a value of type english and
call the method. Do it again for a value of type chinese.
Add a new function named sayHello that accepts a value of type speaker.
Implement that function to call the speak method on the interface value, then create new values
of each type and use the function.
*/

package main

import "fmt"

// declare the speaker interface with a single method call speak
type speaker interface {
	speak()
}

// Declare an empty struct type named english
type english struct{}

// Declare a method named 'speak' for struct type english with a value receiver to implement the speaker interface
// Use "Hello World"
func (english) speak() {
	fmt.Println("Hello World")
}

// Declare an empty struct type named chinese
type chinese struct{}

// Declare a method named 'speak' for struct type chinese with a value receiver to implement the speaker interface
// Use "你好世界".
func (*chinese) speak() {
	fmt.Println("你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker) {
	// call the speak method from the speaker parameter.
	s.speak()
}

func main() {
	// Declare a avar of the interface speaker typ set it its zero default
	var s1 speaker

	// Declare a var of type english
	var e english

	// Assign the english value to the speaker var
	s1 = e

	// Call the speak method against the speaker var
	s1.speak()

	// Do all the above for chinese
	var s2 speaker
	var c chinese
	s2 = &c
	s2.speak()

	// call the sayHello function with the new values and pointers of english and chinese
	sayHello(s1)
	sayHello(s2)

}
