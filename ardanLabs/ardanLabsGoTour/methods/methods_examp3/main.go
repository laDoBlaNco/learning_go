package main

import "fmt"

/*
	METHODS ARE JUST FUNCTIONS:

	So as anthonygg said in one of his posts, methods are really just functions that
	provide some syntatic sugar to provide the ability for data to exhibit behaviors

*/

// here we have a type and two methods declared for it.
type data struct {
	name string
	age  int
}

// The displayName method is using value semantics
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// and setAge is using pointer semantics
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

// NOTE: DON'T ACTUALLY IMPLEMENT SETTERS AND GETTERS in Go. These are not apis with purpose
// and in these cases it's better to make those fields exported.

func main() {
	// let's create a value of type data and call the methods we have created
	d := data{
		name: "Kevin",
	}

	fmt.Println(d) // since we only set the name, the age will show as its zero default 0
	fmt.Println()

	d.displayName()
	d.setAge(46) 
	
	// since methods are really just functions with syntatic sugar, we an also call the
	// methods as normal functions, but with a different syntax.
	// This is what Go is doing underneath the hood actually:
	fmt.Println()
	fmt.Println("What the compiler is doing:") 
	data.displayName(d) 
	(*data).setAge(&d,49) 
	
	// The receiver is really just a parameter and its the first parameter (like self or this)
	// When you call a method, the compiler converts that to a function call underneath. 
	// NOTE: Don't execute methods like this though, even though we might see this in the wild
	// with some tooling messaging.
	
	fmt.Println()
	fmt.Println("Calling Value Receiver Methods with Variable:") 
	// We can also create variables to hold our functions/methods. The function/method is passed
	// by value so our var gets its own copy of the func/meth
	f1 := d.displayName // note we don't use '()' since this isn't a call, its an assignment
	
	// then we can call it now with ()
	f1() 
	
	// let's try to change value of d
	d.name = "Joan" 
	
	// Now let's use our f1 again, but note that we don't see any changes
	f1() 
	d.displayName() // as we can see here, we don't see the change using f1 because it is its own 
	// copy of the original meth. When we use the original meth then we do see the change 
	
	fmt.Println()
	fmt.Println("Call Pointer Receiver Method with Variable:") 
	f2 := d.setAge
	f2(80) // With this we see the changes  due to the pointer semantics. f2  might be getting its
	// own copy of the func, but that func is still a pointer to the original data type
	
}
