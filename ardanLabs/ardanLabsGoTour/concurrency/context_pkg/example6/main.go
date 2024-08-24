// In this last example we see a Context being cancelled and how all contexts derived from it
// are then also cancelled. 


/*
	The use of value semantics for the Context API means each new Context value is 
	given everything the parent Context has plus any new changes. This means if a parent
	Context is cancelled, all children by that parent Context are also cancelled. 
*/

package main

import(
	"context"
	"fmt"
	"sync"
)

// Need a key type
type myKey int

// Need a key value
const key myKey = 0

func main(){
	
	// create a context that can be cancelled
	ctx,cancel := context.WithCancel(context.Background()) 
	defer cancel() 
	
	// we will also use WaitGroup for the orchestration of our grs
	var wg sync.WaitGroup
	wg.Add(10) 
	
	// Create ten goroutines that will derive a Context from the one created above. Each
	// gr places their unique id inside their own Context . The call to WithValue is 
	// passed the main function's Context value as its parent, Then  each gr waits until
	// their Context is cancelled
	for i:=0;i<10;i++{
		go func(id int){
			defer wg.Done() 
			
			// Derive a enw context for this gr from the Context owned by the main func
			ctx:=context.WithValue(ctx,key,id) 
			
			// we then wait until the context is cancelled
			<-ctx.Done()
			fmt.Println("Cancelled:",id) 			
		}(i)
	}
	
	// here main we then cancel the original Context and in turn any derived as well
	cancel()
	wg.Wait() 
}

/*
	Once the cancel() function is called, all 10 contexts become unblocked and print that
	they've been cancelled. So its just one call to cancel to cancel them all. 
	
	This shows that the same Context may be passed to functions running in different goroutines
	A context is safe for simultaneous use by multiple grs
	
	We can't pass a nil Context, even if a function permits it. We should pass a TODO Context
	if we are unsure about which Context to use. One of  the nices things in the Context pkg
	is the TODO func. Since we are always just drafting code, its great to be able to actually
	code in todos in order to move things along and come back to polish later.  sometimes we are
	unsure of what Context to use or even where it will come from. Or maybe we aren't the ones 
	responsible for creating that level Context, so we can't use the Background func. Well we can
	use the todo Context.
	
	USE CONTEXT VALUES ONLY FOR REQUEST-SCOPED DATA
	
	We shouldn't use the Context for passing optional params to functions. this might be the
	most important semantic of all with Contexts. We can't use the Context value to pass data
	into a function with that data is required by the func to execute its code. In other words,
	a func should be able to execute it logic with an empty Context value. In cases where a func
	requires informaton to be in the Context, if that information is missing, the program should
	fail and signal the application to shutdown. 

	

*/
