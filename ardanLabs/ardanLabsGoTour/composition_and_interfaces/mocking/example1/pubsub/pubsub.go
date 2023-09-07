// Package pubsub simulates a package that provides publication/subscription type services
package pubsub

import "fmt" 

// INTERFACE OWNERSHIP
// One thing that is different about Go from other languages is that idea of convention over
// configuration. This really shows itself with how Go handles interface compliance. Because
// the compiler can perform static code analysis to determine if a concrete value implements
// an interface, the dev declaring the concrete type doesn't have to worry about providing 
// interfaces as well. 

// PubSub provides access to a queue system
type PubSub struct{
	host string
	// more fields here
}

// factory function creates a pubsub value for use.
func New(host string)*PubSub{
	ps := PubSub{
		host:host,
	}
	
	return &ps 
}

// Publish sends the data for the specified key
func(ps *PubSub)Publish(key string,v interface{})error{
	fmt.Println("Running Publish") 
	return nil 
}

// Subscribe sets up a request to receive messages for the specified key
func(ps *PubSub)Subscribe(key string)error{
	fmt.Println("Running Subscribe") 
	return nil
}

// As seen above we just implement the API that provides a concrete implementation for publish
// and subscribe. There are no interfaces being provided here because the API does not need one
// This is a single concrete implemenation. 

// What if the application dev wanting to use this new API needs an interface because they have
// the need mock this implemenation during tests? In Go, that dev can declare the interface
// and the compiler can identify the compliance


