package main

import "fmt"

/*
	Its important to remember that in Go the concepts of subtyping or sub-classing really
	don't exist and these design patterns should be avoided
	
	The following is an anti-pattern that we shouldn't follow or implement. We use type
	hierarchies with an OOP patten. This is not something we want to do in Go. Go doesn't
	have the concept of sub-typing. All types are their own and the concepts of 'base'
	and 'derived' types do no exist in Go. This pattern doesn't provide a good design
	principle in a Go program.
*/

// Using Animal to hold our base fields for animals. Here the Animal type tries to be the
// base type for all animals and define the data that is common to all animals. 
type Animal struct{
	Name string
	IsMammal bool
}

// We also try to give some common behavior  to all animals in our methods 
// Speak provides generic behavior for all animals and how they speak
func(a *Animal)Speak(){
	fmt.Printf("UGH! My name is %s, it is %t I am a mammal\n",a.Name,a.IsMammal)
	// While most animals have the ability to speak in one way or the other, we
	// have no idea at this point what sound this animal makes 
}

// Now the problem starts when I start to try to use my base class. here we  attempt to 
// use Animal embedded to make a Dog everything an Animal is plus more. On the surface this 
// will seem to work, but there will be problems. 
type Dog struct{
	Animal
	PackFactor int
}

func (d *Dog)Speak(){
	fmt.Println("Woof!","My name is",d.Name,", it is",d.IsMammal,"I am a mammal with a pack factor of",d.PackFactor)
}

// Then we do the same  with a Cat
type Cat struct{
	Animal
	ClimbFactor int
}

func(c *Cat)Speak(){
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n",
		c.Name,
		c.IsMammal,
		c.ClimbFactor,
	)
}

// Now eveything seems to work out fine and it looks like embedding is providing the
// same functionality as inheritance does in other languages. Then we try to go ahead 
// and group dogs and cats byt he fact they are Animals
func main(){
	
	// create a list (slice) of Animals that know how to speak
	animals := []Animal{
		Dog{
			Animal: Animal{
				Name:"Rocky",
				IsMammal:true,
			},
			PackFactor:5,
		},
		Cat{
			Animal:Animal{
				Name:"Tom",
				IsMammal:true,
			},
			ClimbFactor:4,
		},
	}
	
	// let's have our animals speak
	for _,animal := range animals{
		animal.Speak() 
	}
}

// we get the following errors:
// groupingExam1/main.go:65:3: cannot use Dog{…} (value of type Dog) as Animal value in array 
// or slice literal
// groupingExam1/main.go:72:3: cannot use Cat{…} (value of type Cat) as Animal value in array 
// or slice literal

/*
	When we try to do this, the compiler complains that a Dog and Cat are not an Animal and this
	is true. Embedding isn't the same as inheritance and this is the pattern that we need to
	stay away from. A dog is a dog and a cat is a cat and an animal an animal. I can't pass
	Dogs and cats around as if they are animals because they aren't.
	
	This kind of mechanic is also not very flexible. It requires configuration by the dev
	and unless I have access to the code, we can't make configuration changes over time.
	
	If this is not how we can construct a collection of  Dog's and Cats, then how can we do
	this in Go???
	Its not about grouping through common DNA, its about grouping through common behavior. 
	BEHAVIOR IS KEY.
	
	let's see this in the next example.
*/


