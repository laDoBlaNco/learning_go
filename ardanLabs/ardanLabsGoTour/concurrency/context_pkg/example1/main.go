package main

import(
	"context"
	"fmt"
)

// the context package defines the type Context, which carries deadlines, cancellation signals
// andother request-scoped values across api boundaries and between processes

// Context values are for request-scoped data that passes through programs in a dist system
// What are the semantics?

/*

	The Go programming language has the builtin keyword 'go' to create goroutines, but has no
	keywords or direct support for terminating goroutines. In a real world service, the ability
	to time-out and terminate goroutines is critical for maintaining the health and operation  
	of a service. No request or  task can be allowed to run forever so identifying and managing
	latency is a responsibility every programmer has. 
	
	A solution provided by the Go team to solve this problem is the Context package. It was written
	and intro'd by Sameer Ajmani back in 2014 at the Gotham go conf. Theres also a blog post atÑ
	Talk Video: https://vimeo.com/115309491
	
	Slide Deck: https://talks.golang.org/2014/gotham-context.slide#1
	
	Blog Post: https://blog.golang.org/context	
	
	From here weve seen an evolved set of semantics around contexts. 
	
	INCOMING REQUESTS TO A SERVER SHOULD CREATE A CONTEXT
	The time to create a Context is always as early as possible in the processing of a request or 
	task. Working with Context early in the development cycle will force us to design APIs to 
	take a Context as the first parameter. Even if we aren't 100% sure a function needs a 
	Context, it's easier to remove the Context from a few funcs than try to add one later. 
	
	Its an idiom in Go to tuse the variable name 'ctx' for all Context values. Since a Context
	is aan interface, no pointer semantics should be used. 
	
	type Context interface{
		Deadline() (deadline time.Time,ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{} 
	}
	
	Every function that accepts a Context should get its own copy of the interface value
	
	OUTGOING CALLS TO SERVERS SHOULD ACCEPT A CONTEXT
	The idea behind this semantic is that higher level calls need to tell lower level calls
	how long they are willing to wait. A great example of this is with the http package and
	the version 1.7 changes made to the Do method to respect timeouts on a request. We'll see
	an example of this in Example 5 where we limit the program with a timeout of 50milliseconds
	Basically we'll see the Do method respecting the timeout of 50 milliseconds set inside the
	Context within the Request value. It'll be a higher level function telling the Do method
	(lower level function) how long we're willing to wait around for the operation to be
	completion. 
	
	DO NOT STORE CONTEXTS INSIDE A STRUCT TYPE
	Instead, pass a Context to each function that needs it. Essentially, any func that is
	performing I/O should accept a Context value as it's first param and respect any timeout
	guidelines configured by the caller. In the case of a Request, Go has backwards compatibility
	instead of changing the API's, the mechanic shown in the last section was implemented. 
	
	Before moving on let's see some basic use of Context and how to store and retrieve values
	from it.

*/

// TraceID is will rep a trace id
type TraceID string

// TraceIDKey is the type of value to use for the key. The key is type spectific and only 
// values of the same type will match.
type TraceIDKey int

func main(){
	
	// first let's create a traceID for this request
	traceID := TraceID("f47ac10b-58cc-0372-8567-0e02b2c3d479") 
	
	// declare a key with the value of zero of type userKey
	const traceIDKey TraceIDKey = 0
	
	// Now let's store the traceID value inside the context with a value of zero for the key type
	ctx := context.WithValue(context.Background(),traceIDKey,traceID) 
	
	// we can then retrieve the traceID value from the Context value bag
	if uuid,ok :=ctx.Value(traceIDKey).(TraceID); ok{
		fmt.Println("TraceID:",uuid) 
	}
	
	// Retrieve that traceID value from the Context value bag not using the proper key type
	if _,ok := ctx.Value(0).(TraceID); !ok{
		fmt.Println("TraceID not found!") 
	}
}

// so here we see how we can use type assertion to determine what is  being held by the 
// Context.
