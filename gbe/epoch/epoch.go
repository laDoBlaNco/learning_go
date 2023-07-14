// Go by Example: Epoch
// A common requirement in programs is getting the number of seconds, milliseconds,
// or nanoseconds since the Unix epoch. Here's how to do it in Go

package main

import (
	"fmt"
	"time"
)

func main() {
	// Use time.Now with Unix, UnixMilli or UnixNano, respectively
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// You can also convert integer seconds or nanoseconds since the epoch into the
	// corresponding time.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
