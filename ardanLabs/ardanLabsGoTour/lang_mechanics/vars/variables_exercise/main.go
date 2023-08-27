// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var s string
	var i int
	var b bool

	// Display the value of those variables.
	fmt.Printf("The value of 's':\t%q\n",s)
	fmt.Printf("The value of 'i':\t%v\n",i)
	fmt.Printf("The value of 'b':\t%v\n",b)
	fmt.Println()
	// Declare variables and initialize.
	// Using the short variable declaration operator.
	ss:="Kevin"
	ii:=42
	bb:=true

	// Display the value of those variables.
	fmt.Printf("The value of 'ss':\t%v\n",ss)
	fmt.Printf("The value of 'ii':\t%v\n",ii)
	fmt.Printf("The value of 'bb':\t%v\n",bb)
	fmt.Println()

	// Perform a type conversion.
	f:=float32(3.14)
	// Display the value of that variable.
	fmt.Printf("The value of 'f':\t%v\n",f) 
}
