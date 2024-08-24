package main

import "fmt"

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
	fmt.Printf("%+v\n", engineer)

	fmt.Println(engineer.Name)
	fmt.Println(engineer.Project.Name) // note the .sytnax to get to the specific field data
}
