package main

import(
	"fmt"
	_"log"
	_"os"
	_"os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go by Example: Channel Buffering
// By default channels are unbuffered, meaning that they will only accept sends
// (chan<-) if there is a corresponding receive (<- chan) ready to receive
// the sent value. Buffered channels accept a limited number of values without
// a corresponding receiver for those values. 

func main(){
	
	// first let's make a channel for strings buffereing up to two values
	// without the need of a receiver
	messages := make(chan string,2) 
	
	// Because this channel is buffered, we can send these values into the channel
	// without a corresponding concurrent receive. 
	messages <- "buffered"
	messages <- "channel" 
	
	// later we can receive these two values as usual without the need of a separate
	// routine. 
	p(<-messages)
	p(<-messages) 
	
}

// Next lets take a look at Channel Syncing
