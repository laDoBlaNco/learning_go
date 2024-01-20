package main

import "fmt"

// Methods are just associated funcs, though we borrow the name 'methods' from other oop langs
// this will give us functionality to our structs.
type Engineer struct {
	Name    string
	Age     int
	Project Project
}
type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

// we can create a method or associated func to use some of the data in our struct
func (e Engineer) Print() { // this is using a value receiver (e Engineer). We an also use ptr
	// receiver which we'll discuss later, but are using when we need access to mutate the state of
	// our struct type. With the value receiver we only have access to read the data from our instance
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
	// One way to think of the receiever is that its the connection to the struct that'll be
	// receiving your request for data or access to change data. so here 'e' represents the
	// struct Engineer and when we use this on our instance 'engineer.Print(), e will be our
	// instance of Engineer (engineer)

}
func (e *Engineer) UpdateAge(age int) { // 100% on the same page as Elliot here ;) This is the
	// use of a pointer receiver.
	e.Age = age
}

// Now let's create an ass func that  returns a value. typically if you need to change state you
// need to use pointer recvr and if you just need data, its ok to use a value recvr. Most Go devs
// try to be consistent with everything in there project and others just err on the side of caution
// and always use a *receiver.
func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	println()
	fmt.Println("Structs In Go!")
	println()

	engineer := Engineer{
		Name: "Ladoblanco",
		Age:  46,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	} // the use of the field name is optional, but adds
	// more clarity. If you don't use it, it must be in the right order. Also any fields you haven't
	// specified will be zero defaulted.
	// fmt.Printf("%+v\n", engineer)

	// fmt.Println(engineer.Name)
	// fmt.Println(engineer.Project.Name) // note the .sytnax to get to the specific field data

	// Let's use our new method:
	// Note that even though we changed our recvr's to pointer, we didn't have to change how we
	// call the methods below. Go handles that for us, but if we wanted to we could
	// do it manually as noted below as well.
	engineer.Print()
	(&engineer).UpdateAge(23) // technically since this uses a pointer receiver its telling go that
	// the instance that is receiving our request should be a pointer. Though Go adjusts for this
	// under the hood, we could manually change it to a pointer as I've done here.
	println()
	engineer.Print()

	fmt.Println("Project Priority:", engineer.GetProjectPriority())

}
