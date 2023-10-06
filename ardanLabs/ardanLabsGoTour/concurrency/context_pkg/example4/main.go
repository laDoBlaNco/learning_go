/*

	In this program we will show the semantic nature of the WithTimeout function. The call to
	WithTimeout returns a new Context value and a cancel function. Since the function call
	requires a parent Context, the code uses the Background function to create a top-level
	empty Context. This is what the Background function is for that we've been putting in
	()s.

	Moving forward the Context value created by the WithTimeout function is used. If any
	future functions in the call chain need their own specific timeout or deadline, they
	should also use the appropriate With function and this new Context value as the
	parent.

	Its critically important that any cancel function returned from a With function is
	executed before that function returns. This is why the idiom uses the defer keyword right
	after the With call, as we can see. Not doing this call will cause memory leaks in our
	program
*/

package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	UserID string
}

func main() {

	duration := 150 * time.Millisecond

	// now this Context will both manually cancell and signal a cancel at the specificed duration
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// our channel to signal
	ch := make(chan data, 1)

	// our goroutine again
	go func() {

		time.Sleep(50 * time.Millisecond)

		ch <- data{"123"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
