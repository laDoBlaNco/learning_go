package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"time"
)

// ... for quick debugging
var p = fmt.Println

// A goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// suppose we have a function call f(s). here's how we'd call that in a
	// usual way, running it synchronously
	f("direct")

	// but to invoke this function in a goroutine, use ' go f(s)' This new
	// goroutine will execute concurrently with the calling one. (yes the 'main'
	// func is the principal goroutine).
	go f("goroutine")

	// you can also start a goroutine for an anony func call.
	go func(msg string) {
		fmt.Println(msg)
	}("going") // also note that anony funcs are the only time in Go you'll see
	// a func in a func.

	// Our two function calls are running asynchronously in separate goroutines now
	// Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")

}

// When we run this program, we see the output of the blocking call first, then the
// output of the two goroutines. The goroutines output may be interleaved, because
// goroutines are being run concurrently by the Go runtime.

// Next we'll look at the complement to goroutines in concurrent Go programs:
// channels.
