// In this example of using composition and interfaces, we'll see something that we WANT TO DO 
// in Go. We will group common types by their behavior and not by their state. This patten 
// does provide a good design principle in a Go program.

package main

import "fmt"

// If we use an interface, we can define common method set of behavior that we want to group
// different types of data against. Speaker provides this common behavior below for all
// concrete types to follow if they want to be a part of this group. This is a contract
// for those concrete types to follow.
type Speaker interface{
	Speak() 
}

type Dog struct{
	Name string
	IsMammal bool
	PackFactor int
}
func(d *Dog)Speak(){
	fmt.Printf(
		"Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

type Cat struct{
	Name string
	IsMammal bool
	ClimbFactor int
}
func(c* Cat)Speak(){
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n",
		c.Name,
		c.IsMammal,
		c.ClimbFactor, 
	)	
}

func main(){
	// Now we can create our list
	speakers :=[]Speaker{
		&Dog{
			Name:"Rocky",
			IsMammal:true,
			PackFactor:5,
		},
		&Cat{
			Name:"Patches",
			IsMammal:true,
			ClimbFactor:4,
		},
	}
	
	// And have them speak now
	for _,spkr := range speakers{
		spkr.Speak() 
	}
}

/*
	Guidelines around declaring types:
		- Declare types that represent something new or unique
		- Don't create aliases just for readability
		- Validate that a value of any type is created or used on its own
		- Embed tyeps not becasue you need the state, but becasue we need the behavior
		- If we are not thinking about behavior, we're locking ourselves into a design that
		  we can't grow in the future without cascading code changes.
		- Question types thar are aliases or abstractions for an existing type.
		- Question types whose sole purpose is to share a common set of states.
		
		
	DON'T DESIGN WITH INTERFACES
	Unfortunately too many devs attempt to solve problems in the abstract first. They focus
	on interfaces right away and this leads to interface pollution. As a dev, we exist in two
	modes: a programmer and then an engineer. 
	
	When we are a programmer, we are focused on getting a piece of code to work. Trying to solve
	the problem and break down walls. Prove that our initial ideas work. That is all we care 
	about. This programming should be done in the concrete and is NEVER PRODUCTION READY.
	
	Once we have our prototype of code that solves the problem, then we need to switch over to
	being engineers. We need to focus on how we write the code at a micro-level for data semantics
	and readability, then at a macro-level for mental models and maintainability. We also need to
	focus on errors and failure states.
	
	This work is done in a cycle of refactoring. Refactorring for readability, efficiency, abstraction,
	and for testability. Abstracting is only one fo a few refactors that need to be performed. This 
	works best when we start with a piece of concrete code and then discover the interfaces that are
	needed. Don't apply abstractions unless they are absolutely necessary.
	
	Remember that EVERY problem we solve with code is a DATA PROBLEM requiring us to write data
	transformations. If we don't understand the data, then we don't understand the problem. If we don't
	understand the problem, then we can't write any code. Starting with concrete solutions that are based
	on the concrete data structures is critical. As Rob Pike said, 
	
	"Data dominates. If you've chosen the right data structures and organized things well, the
	algorithms will almost always be self-evident." - Rob Pike
	
	So then when is abstraction necessary? When we see a place in the code where the data could
	change later and we want to minimize the cascading code effects of changing code later that
	would result. We might use abstraction to help make code testable, but we should try to avoid
	this if at all possible. The best testable functions are functions that take raw data in and
	send raw data out. it shouldn't matter where the data is coming from or going.
	
	In the end, start with a concrete solution to every problem. Even if the bulk of that is
	just programming. Then discover the interfaces that are absolutely requireed for the code
	today. 
	
	"Dont design interfaces, discover them." - Rob Pike

*/
