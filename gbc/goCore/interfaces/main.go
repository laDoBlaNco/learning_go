package main

import(
	"fmt"
)
// ...for quick debugging
var pl = fmt.Println


// basically the allow  us to create contracts that say anything can implement said
// interface, must have certain specific methods. Meaning while we can have those
// actions implemented and do things in a specific way, the general action must 
// exist for all types that implement that interface

type Animal interface{
	AngrySound()
	HappySound() 
}

type Cat string
func(c Cat) Attack(){
	pl("Cat Attacks its Prey")
}
func(c Cat) Name() string{
	return string(c)
}
func (c Cat)AngrySound(){
	pl("Cat says Hisssss")
}
func (c Cat) HappySound(){
	pl("Cat says Purrrrr")
}


func main(){

	var kitty Animal
	kitty = Cat("Kitty") // note that this wasn't Cat{"Kitty"}. I've seen this before
	// and basically its not a type literal, it's type conversion. Same as string(123)
	// or int(123.45). Since its a defined type based on string (type Cat string) 
	// it will convert the same as string as in Cat("Kitty"). It wil automatically
	kitty.AngrySound()
	kitty.HappySound() // here we can use our methods on kitty
	
	pl()
	
	// but we can only use methods that are part of the interface Animal:
	// kitty.Attack() // this gives us an undefined error because kitty interface
	// doesn't know about it.
	
	// in order to get to the other methods we need to get to the underlying Cat type
	// and we do that using our type assertion
	var kitty2 Cat = kitty.(Cat) // taking our type kitty which implements the interface Animal
	// and with type assertion we assert that in kitty there's a Cat and assign that to
	// kitty2 (I had an error for missing the '.' Remember type assertion is 
	// kitty.(Cat) not just kitty(Cat) )
	kitty2.Attack() 
	pl("Cats Name:",kitty2.Name()) 
	

}
