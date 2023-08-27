// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

const ( // remember if its in a constant block, there's no need for 'const' on the declaration...stupid
// Declare a constant named server of kind string and assign a value.
	server = "Server1"

// Declare a constant named port of type integer and assign a value.
	port int = 3333
)

func main() {

	// Display the value of both server and port.
	fmt.Println("Server:",server)
	fmt.Println("Port:",port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	result := 3333 / 8.0

	// Display the value of the variable.
	fmt.Println(result)
}
