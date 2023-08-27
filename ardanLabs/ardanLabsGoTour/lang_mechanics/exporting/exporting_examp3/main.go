package main

import(
	"fmt"
	"github.com/ladoblanco/ardanLabsGoTour/exporting/examples" 
)

func main(){
	
	// Again we do the same, but this time we don't access our unexported type direclty,
	// we use a New func
	counter := examples.New(10) 
	
	fmt.Printf("Counter: %d\n",counter) 
}

// And it works. But again this doesn't mean that we should do this in practice nor does it
// mean that we're getting any real protections for it. This should be avoided, and if
// New will return a value, it should be of an exported type.
