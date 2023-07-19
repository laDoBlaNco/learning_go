// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type User struct{ 
	Name string
	Email string
	Age int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u:=User{
		Name:"ladoblanco",
		Email:"whitesidkevin@gmail.com",
		Age:46,
	}

	// Display the field values.
	fmt.Printf("Username is:\t%v\n",u.Name)
	fmt.Printf("User email is:\t%v\n",u.Email)
	fmt.Printf("User age is:\t%v\n",u.Age)

	// Declare a variable using an anonymous struct.
	u2:=struct{
		Name string
		Email string
		Age int
	}{
		Name:"ladoblanco",
		Email:"whitesidekevin@gmail.com",
		Age:46,
	}

	fmt.Println()
	// Display the field values.
	fmt.Printf("Username is:\t%v\n",u2.Name)
	fmt.Printf("User email is:\t%v\n",u2.Email)
	fmt.Printf("User age is:\t%v\n",u2.Age)

}
