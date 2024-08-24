package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	We need to think of channels not a data structures, but as a mechanic for
	signaling. This goes in line with the idea that I send and receive from a
	channel, not read and write to it. If the problem in front of me can't be
	solved with signaling, if the word signaling isnot coming out of my mouth,
	then I need to question the use of channels.

	CHANNEL MECHANICS:

	The cost of having the guarantee at the signaling level is unknown latency.
	The sender won't know how long they need to wait for the receiver to accept
	the signal. Having to wait for the receiver creates blocking latency. The
	sender has to wait, for an unknown amount of time, until the receiver becomes
	available to receive the signal.

	Waiting for the receiver means mechanically, the receive operation happens
	before the send operation. With channels, the receive happens nanoseconds
	before, but it is in fact before. Its like a hand-off. One hand is in the
	air and that's the send, but that hand stays there until another hand
	comes up to recieve. The receive hand grabs the object passed or it receives
	first, then the sending hand let's it go or sends it.  This means the receiver
	takes a signal and then walks away, allowing the sender to now move on with
	a guarantee.

	What if the process can't wait for an unknown amount of time? What if that
	kind of latency won't work? Then the guarantee can't be at the signaling
	level, it needs to be outside of it. The mechanics behind this working is
	that the send now happens before the receive. The sender can perform the
	signal without needing the receiver to be available0 .So the sender gets to
	walk away and not wait. Eventually, you hope, the receiver shows up and
	takes the signal/data.

	This is reducing latency cost on the send, but its creating uncertainty about
	signals being received and therefore knowing if there are problems upstream
	with receivers. The can create the process to accept work that never gets
	started or finished. It could eventually cause massive back pressure ans systems
	to crash. I think the result would be goroutine leaks, memory leaks, and
	that would eventually crash the program.

	The second thing to focus on is, do I need to send data with the signal?
	If the signal requires the transmission of data, then the signaling is a
	1 to 1 between goroutines. If a new goroutine needs to receive the signal
	as well, a second signal must be sent.

	If data doesn't need to be transmitted with the signal, then the signal
	can be 1 to 1 or 1 to many between goroutines. Signaling without data is
	primarily used for cancellation or shutdowns. It's done by closing the
	channel.

	The third thing to focus on  is channel state. A channel can be in 1 of 3
	states.
	1. A channel can be nil state by constructing the channel to its zero default
	   state. Sends and receives against a channel in this state will be blocked.
	   This is good for situations where I want to implement short term stoppages
	   of work.
	2. A channel can be open by using the built-in function make. Sends and
	   receives against channels in this state will work under the following
	   conditions:
			UNBUFFERED CHANNELS:
			Guarantees at the signaling level with the receive happening before
			send. Sending and receiving goroutines need to come together in the
			same space and time for a signal to be processed.

			BUFFERED CHANNELS:
			Guarantees outside of the signaling level with the send happening
			before the receive this time. If the buffer is not full, sends can
			cmoplete else they block. If the buffer is not empty, receives
			can complete else they block.
	3. A channel can be in a closed state by using the built-in close func. I don't
	   need to close a channel to release memory, this is for changing the state.
	   Sending on a closed channel will cuase a panic, however receiving on a
	   closed channel will return immediately.

	With all of this information, I can focus on channel patterns now. The focus
	on signaling is important. The idea is, if I need a guarantee at the signaling
	level or not, based on latency concerns. If I need to transmit data with the
	signal or not, based on handling cancellations or not. I want to convert the
	syntax to these semantics.


	IN SUMMARY:
	Guarantee of Delivery:
	This is based on one question; "Do I need a guarantee that the signal sent by
	a particular goroutine has been received?"
		- Unbuffered channel = guaranteed
		- buffered channel = not guaranteed

	Signaling with or without Data:
	When I am going to signal with data, there are three channel configuaration
	options that I can choose depending on the type of guarantee I need:
		- Guarantee = unbuffered
		- No Guarantee = buffered > 1
		- Delayed Guarantee = buffered = 1

	Signaling without data serves the main purpose of cancellation. It allows one
	goroutine to signal another goroutine to cancel what they are doing and move on.
	Cancellation can be implemented using both unbuffered and buffered channels.
		- First choice = Context
		- Second choice = Unbuffered
		- Smelly code = Buffered

	State:
	The behavior of a channel is directly influenced by its current state. The
	state of a channel can be nil, open, or closed:
		- nil = both send and received blocked
		- open = both send and received allowed
		- close = sends will panic and receives are allowed.

*/

// This example demonstrates the wait for result channel pattern
func main() {
	waitForResult()
}

// waitForResult: In this pattern, the parent goroutine waits for the child
// goroutine to finish some work to signal the result
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("----------------------------------------------------------")

}

/*
	BUFFER BLOAT - 2011
	I should be careful about using large buffers with the idea of reducing
	latency.
		- Large buffers prevent timely notification of back pressure
		- They defeat the ability to reduce back pressure in a timely matter
		- They can increase latency not reduce it
		- Use buffered channels to provide a way of maintaining continuity
		- Don't use them just for performance
		- Use them to ohandle well defined bursts of data
		- Use them to deal with speed of light issues between handoffs.


*/
