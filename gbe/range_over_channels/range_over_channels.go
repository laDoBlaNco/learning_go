package main

import (
	"fmt"
)

var p = fmt.Println

// In a previous example we saw how 'for' and 'range' provide iteration over basic data
// structures. But we can also use this syntax to iterate over values received from a
// channel.

func main() {

	// Here let's iterate over 2 values in the QUEUE channel
	queue := make(chan string, 2) // make a channel for strings with a buff of 2
	queue <- "one"                // then we put 'one' in our channel
	queue <- "two"                // since we have a buff of 2 we can fit another string in there
	close(queue)                  // by closing our queue we are tell Go there's no more and we can 'range'

	// This range iterates over each element as its received from queue. Because we closed
	// the channel above, the iteration terminates after receiving the 2 elems
	for elem := range queue {
		fmt.Println(elem)
	}

	// This example shows that its possible to close a non-empty channel but still have
	// the remaining values be received.

	// Now let's check out 'Timers'

}
