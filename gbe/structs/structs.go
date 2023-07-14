package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go's structs are typed collections of fields. They're useful for grouping
// data together to form records.

// This 'person' struct type has a name and age field
type person struct {
	name string
	age  int
}

// newPerson func will construct a new person struct with the given name
func newPerson(name string) *person { // note we are returning a *ptr to a person and
	// not the actual person. This means we are going to create heap memory

	// You can safely return a pointer to a local variable as a local variable
	// will survive  the scope of the function (on the heap as far as  I understand)
	p := person{name: name}
	p.age = 46
	return &p
}

func main() {
	
	// This syntax creates a new struct literal
	p(person{"Bob",20}) 
	p()
	 
	// You can assign the fields by name when initializing a struct
	p(person{name:"Alice",age:30}) 
	p() 
	
	// Omitted fields will simply be zero-valued
	p(person{name:"Fred"}) 
	p() 
	
	// An & prefix gives us the pointer (memory address) to the struct
	p(&person{name:"Ann",age:40}) 
	p() 
	
	// Its idiomatic to encapsulate new struct creation though in 'new' funcs
	// (constructors in other langs) 
	p(newPerson("Jon")) 
	p() 
	
	// Access struct fields with a dot
	s:=person{name:"Sean",age:50}
	p(s.name) 
	p() 
	
	// you can also use dots with struct pointers - the pointers are automatically
	// dereferenced
	sp := &s
	p(sp.age)
	// same as:
	p((*sp).age) 
	p() 
	
	// structs are also mutable
	sp.age = 51 // note we are using a pointer
	p(sp.age) 
	
}
