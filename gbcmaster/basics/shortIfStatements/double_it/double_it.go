// scoping with control structures and short statements. The vars in this
// structure below can only be used in the branches of the if statement.
// They don't exist outside of that statement.

// Also we see below how shadowing can hurt us. there are two 'n' variables
// below. We either need to assign n,err separately or use '=' instead of ':=' or

package main

import (
	"fmt"
	_ "log"
	"os"
	_ "os/exec"
	"strconv"
)

// ... for quick debugging
var p = fmt.Println

func main() {

	var (
		n   int
		err error
	)

	if a := os.Args[1:]; len(a) != 1 {
		// this branch can only see the 'a' var
		fmt.Println("Give me a number.")
		// below we use '=' instead of ':=' so as not to shadow n,err from main
		// asi working with the same var throughout main.
	} else if n, err = strconv.Atoi(a[0]); err != nil {
		// this branch can see the 'a', 'n', and 'err' vars
		fmt.Printf("Cannot convert %q\n", a[0])
	} else {
		// this last branch can also see all vars already declared in statement
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[0], n)
	}
	fmt.Printf("n is %d. 👻👻👻 - you've been shadowed ;)", n)
}

// Note since there is nothing outside of this if statement in main, there's no
// need to 'return' from any of the branches, etc.
