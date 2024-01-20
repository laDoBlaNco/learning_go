package main

import "fmt"

// Engineer - stores the name and age of an engineer
type Engineer struct {
	Name string
	Age  int
}

func (e *Engineer) UpdateAge() { // pointer receiver using the e *Engineer syntax used to  update
	// or have access to the receiver which is the instance of our type or caller of the method
	e.Age++
}

func (e Engineer) UpdateName() { // This is a value receiver which uses e Engineer (no *) and as we
	// can see since we are using a value and not a pointer, we don't have access to the actual caller
	// as we are only working on a copy. so the method seems to work,but the changes are thrown away
	e.Name = "New Name"
	fmt.Println(e) // we can see  the change was done here, but again it was thrown away

}

func UpdateAge (e *Engineer){ // same as the method but the  receiver is an arg of our general func
	e.Age++
	
}

func main() {

	fmt.Println("Pointers in Go!")

	var name string
	name = "ladoblanco"
	fmt.Println(name)

	ptr := &name
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = "kelen"
	fmt.Println(name)

	println()

	lado := &Engineer{
		Name: "Ladoblanco",
		Age:  46,
	}
	fmt.Println(lado)
	lado.UpdateAge()
	fmt.Println(lado)
	lado.UpdateName()
	fmt.Println(lado)
	UpdateAge(lado)
	fmt.Println(lado) 
}
