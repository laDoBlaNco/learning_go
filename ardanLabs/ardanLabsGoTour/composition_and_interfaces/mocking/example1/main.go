package main

import(
	"fmt" 

	"github.com/ladoblanco/ardanLabsGoTour/composition_and_interfaces/mocking/example1/pubsub" 
)

// The best way to take advantate of embedding is through the compositional design pattern.
// The idea is to compose larger types  from smaller types and focus on the composition of 
// behavior. 

// Here the dev can implement their own publisher interface and pubsub implementation for 
// testing. NOTE the key here is that this application dev doesn't use any concrete implementation
// directly, but decouples themselves through their own interface.

// publisher is an interface to allow this package to mock the pubsub package support
type publisher interface{
	Publish(key string,v interface{})error
	Subscribe(key string)error 
}

// mock is a concrete type to help support the mocking of the pubsub package
type mock struct{} 

// Publish implements the pub interface for the mock
func(m *mock)Publish(key string,v interface{})error{
	fmt.Println("Running mock.Publish") 
	return nil
}
// Subscribe implements the publisher interface for the mock
func(m *mock)Subscribe(key string)error{
	fmt.Println("Running mock.Subscribe") 
	return nil 
}

func main(){
	
	// Let's create a slice of publisher interface values. Assign the address of the pubsub.PubSub
	// value and the address of the mock value
	pubs :=[]publisher{
		pubsub.New("localhost"), 
		&mock{},
	}
	
	// Now let's range over the interface value using our value semantic form to see how the
	// publisher interface provides the level of decoupling the user needs. The pubsub package
	// didn't need to provide the interface type. 
	for _,p :=range pubs{
		p.Publish("key","value") 
		p.Subscribe("key") 
	}
}


