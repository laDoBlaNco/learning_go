// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type User struct{
	Name string
	Email string
	Title string
}

// Create a function that changes the value of one of the user fields.
func titleChange(u *User,newTitle string) {

	// Use the pointer to change the value that the
	// pointer points to.
	u.Title = newTitle
}

func main() {

	// Create a variable of type user and initialize each field.
	user1 := User{
		Name:"Kevin",
		Email:"whitesidekevin@gmail.com",
		Title:"Director_of_Operations",
	}

	// Display the value of the variable.
	fmt.Println(user1)

	// Share the variable with the function you declared above.
	titleChange(&user1,"Fulltime_Go_Dev")

	// Display the value of the variable.
	fmt.Println(user1)
}
