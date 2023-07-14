package main

import "fmt"

// Trying to get clear in my head the diff between a value receiver and a pointer
// receiver.

// Functions declare the types of args, the return values and the function body

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person { //a function taks a string & int and returns *pointer
	return &Person{
		Name: name,
		Age:  age,
	}
}

// A method uses the same syntax but with the addition of a receiver before the
// func name
func (p *Person) isAdult() bool { // isAdult method on a *Person type
	return p.Age > 18
}

// Now the diff between a value receiver and a pointer receiver.

// Value receivers make a COPY of the type and pass it to the function. The function
// stack now holds an equal object but at a different location on memory. That means
// any changes doen on the passed obj will remain local to the method. The original
// obj will remain unchanged.

// Pointer receivers pass the address of the type to the function. The function stack
// has a reference to the original obj. So any modifications on the pass obj will
// modify the original object.

func ValueReceiver(p Person) {
	p.Name = "Kevin"
	fmt.Println("Inside ValueReceiver: ", p.Name)
}

func PointerReceiver(p *Person) {
	p.Age = 24
	fmt.Println("Inside PointerReceiver model: ", p.Age)
}

func main() {
	p := Person{"ladoblanco", 46}
	p1 := &Person{"Anthony", 45}

	ValueReceiver(p) // running the ValueReceiver func on p won't change p
	fmt.Println("Inside Main after value receiver: ", p.Name)
	PointerReceiver(p1) // running the PointerReceiver func on p1 will change it
	fmt.Println("Inside Main after pointer receiver: ", p1.Age)
}

// What to choose - If you want to change state, manipulate its value use a pointer
// receiver. If you DON'T NEED TO change state or manipulate a value, use a
// value receiver.

// ** pointer receivers don't make copies so they are also more efficient if the
// type is large.

// ** value receivers are concurrency safe since you are working on copies, pointer
// receivers are not.

