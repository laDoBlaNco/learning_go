package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println(div(1, 0)) // this will give a panic, which gophers don't like
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))
}

func safeDiv(a, b int) (q int, err error) { // this simple adjustment is like the Try Block
	// in other languages. We just return 2 values, one of which is the possible err
	// we can do this with defer anonymous func using builtin recover func.
	// The Recover func returns a nil val if there is no panic, and if there is a panic
	// its returns the current thing that's panicing. It returns 'any' (empty interface)
	// not an error. Now if this is run as is, we still get returned error nil
	// we can change that by simply using 'named' return values
	// q & err are now local vars in safeDiv
	// (just like a & b)
	defer func() {
		// e's type is 'any' (or interface{}), *not* error)
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	// panic("ouch!")

	return a / b, nil

	/* Miki don't like this kind of programming, though we'll see it used in practice
	   Because it makes it more difficult to see where everything is coming from.

	q = a / b // If we don't set the named returned vars, they will still be
	// returned, but with their default values.
	return // a naked return

	*/
}

func div(a, b int) int {
	return a / b
}

// the main usage for named return values is in conjunction with 'recover()'. We can
// use them in other areas, but this is the practical use.
// When we do this we'll typically use the 'naked retur' since we already have the
// return values named. When using the named return values they will always be
// returned. But if we don't set them to something in the function we will receive
// their zero default vals.

// Note: error is just an interface - type error interface{Error() string}

// Miki (instructor) doesn't use recover() much as its better to have the program
// crash and figure out how to address then to have it recover with a log and then
// continue in an unstable state that you aren't aware of. But both ways are seen
// everyday Go code.

// Go Proverb to remember: Don't panic.
