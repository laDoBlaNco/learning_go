package main

import (
	"errors"
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// These are the new features of go 1.20. Thought I'd run through these and then
// get back to the normal flow of my learning.

func main() {

	// Wrapping multiple errors in a single error with the Join func
	err1 := errors.New("err1")             // normal new error text message
	err2 := errors.New("err2")             // another new error text message
	err3 := fmt.Errorf("err3: [%w]", err2) // here using fmt.Error for formatted msg
	err := errors.Join(err1, err3)         // here we are using errors.Join to create one err
	// from the previous. Note we are joining errors.New and fmt.Errorf results

	if errors.Is(err, err1) && errors.Is(err, err2) && errors.Is(err, err3) {
		p("one for all and all for one")
	} // errors.Join puts them all together into one error so Go sees them as
	// the same now.
	p(errors.Unwrap(err3)) // we can also take it back by 'unwrapping' the joined err

	// There are also some changes to the optimization and cli which I won't put
	// here but they are using "skip=" to  skip certain tests and Profile-Guided
	// Optimizations.

	// Memory arenas was removed from the release.

}
