// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how constants do have a parallel type system. Again this means that they are 
// actually two types in one and can be converted to the one that bets fits the need of the scenario.
package main

import "fmt"

const (
	// Max integer value on 64 bit architecture.
	maxInt = 9223372036854775807

	// Much larger value than int64 and the reason is can be larger is because its untyped
	bigger = 9223372036854775808543522345

	// As  you see if we try to put that bigger untyped 256 precision number in our int64 type box, it
	// Will NOT compile
	// biggerInt int64 = 9223372036854775808543522345
)

func main() {
	fmt.Println(maxInt)
	// fmt.Println(bigger) // but I can't print it cuz fmt.Println uses int64 box in its algo. But it does
	// compile, so the precision is real.
	fmt.Println("Will Compile")
}
