package main

import "fmt"

// Defining a struct - type <structName> struct{field type field type}
type person struct {
	name string
	age  int
	pet  string
} // note that there are no commas in the field names.

var fred person // note we used our struct as the 'type' of person -

func main() {

	fmt.Println(fred) // fred has default 0 values.

	bob := person{} // note this first way of creating the instance then adding vals
	bob.name = "William"
	bob.age = 32
	bob.pet = "cat"
	fmt.Println(bob) // agani default 0s

	// when creating an instance with values you do use both ':'s and ','s
	sam := person{name: "Samantha", age: 20, pet: "dog"} // here we add vals with instance
	fmt.Println(sam)

	sam.age = 21
	fmt.Println(sam.age)

	// Anonymous structs
	var user struct { // note there's no struct name and we use long var dec
		name string
		age  int
		pet  string
	}
	user.name = "joe" // then we just add in the values
	user.age = 50
	user.pet = "fish"

	pet := struct { // here we use short dec for the var
		name string
		kind string
	}{ // interestingly here we just pegue the actual struct literal with values
		name: "Fido", // we can only do this because its anonymous. if we were
		kind: "dog",  // creating a type struct, we wouldn't want actual values as
	} // it needs to be general to be used by other vars

	fmt.Println(user)
	fmt.Println(pet)
}

// Structs are basically when we need to design and create types that aren't available
// in the built-in library. A container that holds a collection of fields. You use
// themt o group a related set of data together.

// People as have I make comparisons of Gos structs and OOP classes. They aren't the
// same, though Go does borrow from those concepts which are helpful with structs.

// We also can work with anonymous structs. Basically its declaring a var directly
// to struct type without actually giving the struct type a name. You can also
// embed structs like this:
// type Base struct{
// 	b int
// }
// type Container struct{
//  Base
//  c string
// }
// The inner struct is then accessible from the outer: co ;= Container{}
// co.b = 1

// Gerald says something interesting. In reality we are just creating a new type
// based on the original type 'struct' just the same as we create a new myInt based
// on int. The magic is in the 'type' keyword then, not the 'struct'
