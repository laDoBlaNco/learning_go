package main

/*
	Interface pollution comes from the fact that people are designing software with interfaces
	instead of discovering them as Rob Pike tells us to.
	
	We should design a concrete solution to the problem first. Then we can discover where the 
	program needs to be polymorphic in our refactoring, IF AT ALL
	
	Here are some things some devs say:
	
	"I'm using interfaces because we have to use interfaces" -- NO we don't have to use interfaces
	We use interfaces when its PRACTICAL to use them. There is a cost to using interfaces: a level 
	of indirection and allocation when we store concrete values inside of them. Unless the cost
	of the allocation is worth what we're gaining by decoupling, then we shouldn't be using 
	interfaces. 
	
	
	"I need to be able to test my code so I need to use interfaces" -- NO. We need to design our API
	for the user first, not for our tests. If the API is not testable, then we should question if its
	even usable. There are different layers of API's as well. The lower level unexported API's can 
	and should focus on testability. The higher level exported API's need to focus on usability.
	
	Functions that accept raw data in and return raw data out are the most testable. Separate the
	data transformation from where the data comes from and where it is going. This is a refactoring
	exercise we need to do during the ENGINEERING coding cycle, not the PROGRAMMER coding cycle.
	
	Here is an example of interface pollution by improperly using an interface when one is not
	needed
*/

// A Server interface defining a contract for TCP servers. The problem here is that we don't need
// a contract, we need an implementation first. There will only be one implementation as well, 
// especially since we are the one's implementing it. We do not need someone else to implement it
// for us, so no need for an interface.
type Server interface{ // interface or "contract" 
	Start() error
	Stop() error
	Wait() error
}

// Plus this interface is based on a noun and not a verb. Interfaces are all about behavior
// not state. Concrete types are nouns since they represent the concrete problem. Interfaces
// describe the behavior and Server is not behavior.

// Here is the server for our Server implementation
type server struct{
	host string
	// more fields of course after this one
}

// NewServer (factory function) returns an interface value of the type Server with a server
// implementation
func NewServer(host string)Server{
	// SMELL - Storing an unexported type pointer in the interface
	return &server{host} 
}

// Start allows the server to begin to accept requests
func(s *server)Start()error{
	// implemenation stuff goes here
	return nil
}

// Stop shuts down our server
func(s *server)Stop() error{
	// implemenation stuff goes here
	return nil
}

// Wait prevents the server from accepting new connections
func(s *server)Wait()error{
	return nil 
}

func main(){
	
	// create a new Server
	srv := NewServer("localhost") 
	
	// use our API
	srv.Start() 
	srv.Stop() 
	srv.Wait() 
}



/*
	NOTE: Here are some ways to identify interface pollution:
		- A package declares an interface that matches the entire API of its own concrete type
		- The interfaces are exported but the concrete types implementing the interface are
		  unexported.
		- The factory function for the concrete type returns the interface value with the 
		  unexported concrete type value inside. 
		- The interface can be removed and nothing changes for the user of the API
		- The interface is not decoupling the API from change. 
		
		
	Guidelines around interface pollution:
	Use an interface:
		- When users of the API need to provide an implementation detail
		- When APIs have multiple implemenations that need to be maintained
		- When parts of the APIs that can change have been idenetified and require decoupling
		
	Question an interface:
		- When it's only purpose is for writing testable APIs (write usuable APIs first)
		- When it's not providing support for API to decouple from change
		- When it's not clear how the interface makes the code better
*/

// let's remove the pollution in the next example
