// In this program we are going to use the WithCancel function  to see how we can do
// timeouts as we discussed in the outset.

/*
	THE CHAIN OF FUNCTION CALLS BETWEEN THEM MUST PROPAGATE THE CONTEXT
	
	This is an important rule since a Context is request or task based. We want the Context
	and any changes made to it during the processing of the request or task to be propogated
	and respected. 
	
	Imagine a handler function named List that is executed when the user makes an http request
	for a certain endpoint. The handler accepts as its first param a Context, since its part
	of a reqwuest and will perform i/o. We would see the same Context value propogated down the 
	call stack. A new context value isn't created since this function request no changes to 
	it. If a new top-level Context value would be created by this function, any existing 
	Context information from a higher-level call associated with this request would be lost.
	This isn't what we want. 
	
	This is why everything accepts and returns a Context value so that it can be passed 
	throughout our program from one func to the other. that function as well respecting
	any timeout information set in the Context from any caller above.
	
	Because each function can add/modify the Context for their specific needs, and those
	changes should not affect any function that was called before it, the Context  uses
	value semantics. This means that any change to a Context value creates a new Context
	value that is t hen propatated forward to the next func.


*/

package main

import(
	"context"
	"fmt"
	"time"
)

func main(){
	
	// let's create a context that is cancellable only manually. The cancel function must be
	// called regardless of the outcome.
	ctx,cancel:=context.WithCancel(context.Background()) 
	defer cancel() // canceel will happen regardless
	
	// Lets ask the goroutine to do some work for us to make it i/o
	go func(){
		
		// wait for the work to finish. If it takes too long move on
		select{
			case <-time.After(100*time.Millisecond):fmt.Println("moving on")
			case <-ctx.Done():fmt.Println("work complete") 
		}
	}()
	
	// Simulate other work
	time.Sleep(50*time.Millisecond) 
	
	// report the work is done
	cancel() 
	
	// Just hold the program to see the output
	time.Sleep(time.Second) 
	
}

// in the next example we'll take a look at WithDeadline and how that works with timeouts
// as we saw here with WithCancel
