package main

import "fmt"

// Basic sends and receives on channels are blocking. However, we can use select with a
// default clause implement non-blocking sends, receives, and even non-blocking multi-way
// selects.

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is available on messages then select
	// will take the <-messages case with that value. If not it will immediately take
	// the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similiarly. Here msg cannot be sent to the messages
	// channel, because the channel has no buffer and there is no receiver. Therefore
	// the default case is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multple cases above the default clause to implement a multi-way non-
	// blocking select as well. Here we attempt non-blocking receives on both messages
	// and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
