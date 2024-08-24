package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

func main() {

	city := "Limbo"

	switch city {
	case "Paris", "Lyon": // case clause with case condition
		fmt.Println("France")
		break // you can use breaks as in other langs, but they are necessary
	default:
		fmt.Println("Where?") // this can go anywhere and there can only be 1
	case "Tokyo":
		fmt.Println("Japan")
	}

	// behind the scenes Go changes the switch statements into an if statements
	// this allows Go to evaluate any type of runtime expression in the switch
	// statements while in langs like C you can't
	if city == "Paris" || city == "Lyon" {
		fmt.Println("France")
	} else if city == "Tokyo" {
		fmt.Println("Japan")
	} else {
		fmt.Println("Where?")
	}

	p()

	i := 0
	switch { // bool expression 'true' not needed. 'switch true {...} '
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}

	p()

	switch i = 1000; true { // here we moved the  i assignment after switch and a ';'
	case i > 100: // if we remove the 'true' it still works.
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Println("number")
	}

}

// Rulez:
// 1. Values in case conditions must be unique among other case conditions
// 2. The types of the switch and case condition expressions must be comparable

// each case creates its own block. So anything in that block is only visible
// in it self. So if you create vars in a case block, the other cases won't know
// about it.

// the 'default' clause is the 'else' statement of switch. so if no case matches
// the input the default gets run. Note, default doesn't have to be the last one
// it can be anywhere and still works the same.

// with multiple case conditions we don't need multiple clauses with the same
// result, we can add more with commas

//Using bool expressions in a switch statement. switch true{...} is the same
// as switch{...}. the type of the case  condition as a bool expression. This way
// the types of the condition and the conditions still match.

// in Go there's no automatically falling through since Go puts in the 'breaks'
// itself. So we use fallthrough keyword. In the example able we can remove
// repeated strings by using the next clause statements. This way we still have
// 3 options (big positive number, positive number, number). Those these are hardly
// used in production, they are useful in some cases.

// finally we can do the  'short switch statement' which is just like the short
// if statement in that we are assigning a var that will only be visible in the
// switch as part of the structure.
