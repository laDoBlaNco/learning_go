package main

import "fmt"

// some people like to think of interfaces as 'contracts', and another great description I heard
// from one teacher was to think of them as 'job descriptions'. I like it better because I
// think its fits much better. Contracts are very strong, black and white, but job descriptions
// are a bit looser. You MUST fulfill certain criteria, but if you come with more skills, you can
// still do the job. So I can have ass funcs that fulfill the interface in addition to 50 others
// that dont', and I'm still said to be implementing the interface.

// Typically interfaces are implemented at the package level depending on the dependency. We'll see
// that in some of the future courses.

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Employee Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{"Ladoblanco"} // the reason we needed to fuss with the & here is because
	// we are using a func in addition to the method and the func is requesting an Employee
	// interface and that interface a method, and that method we are using is a pointer receiver
	// So for our Engineer to fulfill that interface it also has to be a pointer.

	// The key to remembering about this piece with the pointers is that it really only comes up
	// with the interfaces. Otherwise their use is pretty intuitive. But again, Interfaces need
	// have their contracts or job descriptions fulfilled, and if they are requiring a method
	// that hast a pointer receiver, they are fulfilled with a pointer, if its a value receiver,
	// than either will work.
	manager := &Manager{"Odalis"}
	PrintDetails(engineer)
	PrintDetails(manager)
}
