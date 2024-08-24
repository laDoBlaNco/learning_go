package main

import (
	"context"
	"fmt"
	"time"
)

// before we talked about the fact that when you send something out on a goroutine, you lose
// control of it in the sense that it becomes its own being, out in the wild doing what it was
// told to do. The best you can do is wait for it to get done before you close shop. But when you
// have multiple grs out there and you can't control the order the come back to you in, how do you
// keep control of the results?

// SELECT/CASE

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Millisecond)  
	defer cancel() 

	// Select can handle various routines, each case is a channel operation, it can be a send
	// or receive. Remeber that select will take whatever case comes back first. If you want it
	// to do something else, you put it in a for loop to keep receiving. The other channels are
	// just being blocked, waiting on someone else to receive.
	select {
	case val := <-ch1:
		fmt.Println("ch1:", val)
	case val := <-ch2:
		fmt.Println("ch2:", val)
	// one thing we use select for a lot is timeouts and cancellation
	// case <-time.After(5 * time.Millisecond):
	case <-ctx.Done():
		fmt.Println("timeout")
	}

	// select {} // This would block forever for example. Blocking without consuming any CPU
	// gRPC framework is something that uses context a lot. So if I start using that with my
	// microservices, I'll get more experience with contexts
	
	// Context is an interface and we can ask a lot from the context. Err(), Done(), Deadline(),
	// All of these are things that we can ask from context or about the context. The context
	// gets passed around all the time. Because its being passed around it can also be used as
	// a key/value store as well with Value(key any) any. 
}
