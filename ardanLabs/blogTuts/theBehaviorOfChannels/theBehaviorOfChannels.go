package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Many make the mistake of thinking about channels as structures, but that lead of understanding
// will only lead us to a lot of bad and complicated concurrent code. The real secret of channels
// is the signaling. A channel allows communication from one goroutine to another signaling about
// a particular event. Its at the core of everything we need to do with channels.

// We need to master 3 attritbutes:
// 1. Guarantee Of Delivery
// 2. State
// 3. With or Without Data

func main() {
	// These 3 attributes are the basis of a design philosophy around signaling.

	// GUARANTEE OF DELIVERY:
	// "Do I need a GUARANTEE that the signal sent by a particular goroutine has been received?"
	// Depending on how we answre that question, we'll know if we need to use a buffered or non-buffered
	// channel. Unbuffered = Guaranteed that was was sent was received, while Buffered = Not Guaranteed.
	// And knowing if you do or don't need a guarantee is crucial when writing concurrent software.

	// STATE:
	// "What is the current state of the channel?"
	// Behavior of a channel is directly influenced by its current state:
	// 1. nil (zero default)
	// 2. open
	// 3. closed

	// ** nil channel - a channel is in nil state when it is declared to its zero default
	var c chan string
	//  a channel can be placed in a nil state as well by explicitly setting it
	c = nil

	// ** open channel - a channel is in an open state when it's made using the built-in make func
	c = make(chan string)

	// ** closed channel - a channel is in a closed state when it's closed using the built-in close func
	close(c)

	// The importance of state is that it determines how send and receive operations behave:
	// NOTE: Signals are 'sent' and 'received' through a channel. DON'T say 'read' and 'write'
	// as that will give us the wrong mental model and channels don't perform i/o
	// State behavior:
	// 1. nil = send -> blocked & receive -> blocked
	// 2. open = send -> allowed & receive -> allowed
	// 3. closed = send -> panic & receive -> allowed

	// Understanding state and guarantee of delivery will allow us to start to make design choices
	// and understand and quickly spot bugs just by reading the code as we know how channels
	// will behave.

	// WITH AND WITHOUT DATA:
	// "Do I need to signal with or without data?"
	// Signaling WITH DATA is done by a send on a channel. Its normally used when you want to
	// start a new task or report back a result
	ch := make(chan string)
	go func() { ch <- "paper" }()
	<-ch

	// a signal WITHOUT DATA is done by closing a channel. NOTE we are signaling from one goroutine
	// to another that we are done for business.
	// 'close(ch)'
	// Normally this is when we want to tell a goroutine to stop what its doing, or a goroutine
	// reports back they are done with no result or the goroutine reports that it has comopleted
	// processing and shut down.

	// NOTE: one of the benefits of signaling without data is that a single gr can signal many grs
	// at once. When we are sending data its a 1:1 exchange.

	// SIGNALING WITH DATA:
	// If we are signaling with data then we can have one of the following:
	// Unbuffered = Guarantee that a signal being sent has been received. This is because  the
	// send can't complete until the Receive happens (or else its blocked)

	// Buffered>1 = No Guarantee that a signal being sent has been received because the signal
	// happens the receive of the signal completes.

	// and Buffered=1 = Delayed Guarantee that the previous signal that was sent has been
	// received because the receive of teh first signal, happens before the send of the second
	// signal completes.
	// NOTE: The buffer size should never be just a random number. It must always be calculated for
	// some well defined constraint.

	// SIGNALING WITHOUT DATA:
	// Mainly this is used for cancellation to tell a goroutine that we are done here and can move
	// on. We can use both buffered and unbuffered channels, but sending nothing on a buffered
	// channel, is just bad code.
	// Preferences should be, in order, :
	// 1. context.Context
	// 2. Unbuffered
	// 3. Buffered - code smelly

	// The built-in close function is used to signal without data. We can still receive data on a
	// closed channel. Any receive on a closed channel will not block and the receive operation
	// will always return.

	// As mentioned above the preferred way to do this is with context.Context and that's really
	// just an unbuffered channel with an automatic close(ch) underneath. If we do want to use our
	// channel for cancelling only, the idiomatic way is with a chan struct{} (empty struct).
	// It is the zero-space idiomatic way to indicate channel is only used for signaling.

	// SCENARIOS;
	// Signal with data -- Guarantee - UnBuffered
	waitForTask()
	waitForResult()

	// Signal with data -- No Guarantee - Buffered > 1
	fanOut()
	selectDrop()

	// Signal with data -- Delayed Guarantee - Buffered = 1
	waitForTasks()
	
	// Signal without data -- Context
	withTimeout() 

}

// Let's create some scenario functions below to be run in MAIN

// SIGNAL WITH DATA - GUARANTEE - Unbuffered Channels
// Bill K likes to think of goroutines as people when reading and writing about channels, so
// we'll do the same, as I've done the same before as well and its easier to contemplate what's
// going on.

// When we need to know that a signal beign sent has been received, two scenarios come into play
// 1. Wait For Task - you send waiting for a task to be run based on your sent and received signal
// 2. Wait For Result - you wait on the result of a worker sending you the result of a task

func waitForTask() {
	// Scenario 1 - Wait for task. Here we are a hiring manager and we want our new employee
	// to perform a task but they need to wait until we are ready. As  if we need to hand
	// them a piece of paper before they can start
	ch := make(chan string)

	go func() {
		p := <-ch // here our goroutine (worker) is sent out and told to wait for our signal
		// so this 'blocks' the worker routine until it receives a signal from main

		// Employee performs work here
		fmt.Println(p) // the main function may end before this has a chance to print.
		// Employee is done and free to go.
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper" // here we send the paper signal from our main routine. Since we are using
	// an unbuffered channel, we get a guarantee  that the worker has received the 'paper' once
	// our send operation completes. The receive happens before the send.
}

func waitForResult() {
	// Scenario 2 - Wait for result - This time we want our new employee to perform a task
	// immediately when they are hired, and we need to wait for the result from their work.
	// We need to wait because we need the paper from them before we can continue.
	ch := make(chan string) // unbuffered channel

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		ch <- "paper" // task to send signal is complete here as long as main receives it won't
		// be blocked.

		// worker is done and free to go
	}()

	p := <-ch // 'paper' is received so the goroutine work can go away and we have our result to
	// print
	fmt.Println(p)

	// Again with the use of an unbuffered channel the receipt is guaranteed as nothing is sent
	// until it is received, either from main -> gr or gr -> main.
	// In this case with the unbuffere, the receive happens again before the send and the worker
	// is guaranteed that we have received the result. Once this guarantee is had, then the worker
	// is done and free to go. In this case, we don't know how long its going to take to receive
	// the result. We just wait for it.

	// So while the BENEFIT of the unbuffered channel is the guarantee, the COST is unknown latency
	// Its something we need to live with if we want the guarantee.

}

// SIGNAL WITH DATA - NO GUARANTEE - Buffered Channels > 1
// Now when I don't need to know that a signal being sent has been received, the following 2
// scenarios come into play:
// 1. Fan Out - you throw a well defined number of workers at a problem and they work it
//    concurrently.
// 2. Drop - this pattern allows me to throw work away when our workers are at capacity

// Buffered channels have defined space that can be used to store the data being sent. We decide
// how much space we need by answering the following:
// 1. Do i have a well defined amount of work to complete? How much work is there?
// 2. If my worker can't keep up, can I discard new work? How much outstanding work puts me at
//    capacity?
// 3. What level of risk am I willing to accept if my program terminates unexpectedly? Anything
//    waiting in the buffer will be lost!!
// If these question don't make sense for our situation, than using a buffer any larger than 1
// is probably wrong.

func fanOut() {
	// Here you have a worker for every task and you know exactly how many report backs you will
	// receive. They all take turns placing there result in your box, but you don't need to know
	// the order that they do it in.

	// This time I'm a manager with a team of employees. I have an individual task I want each
	// worker to perform. As each individual finishes their task, they provide me with the result
	// int eh box on my desk.
	emps := 20
	ch := make(chan string, emps) // emps here is the capacity of the buffer

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		emps--
	}
	// This time the channel is created with 20 bufferes or a buffere with capacity of 20. Then
	// all 20 workers immediately go to work (with each gr). Then the workers send their results
	// and nothing is blocked since our channel can hold up to 20 pieces of data without a receipt
	// We then work through a for (while) loop and block the main routine till we received all
	// 20 results and print each one. Note how we decrement emps-- to ensure we aren't still asking
	// for stuff after the workers are done. We could and then we would start to get junk data
	// a bunch of zero defaults

}

func selectDrop() {
	// The benefit of this scenario is to continue to accept work and never apply pressure
	// or latency in the acceptance of that work. The key here is though that we need to know
	// when we truly are at capacity so that we don't under or over commit on what we can get
	// done.

	// So here we are a manager and hire a single worker. I have an individual task and I don't
	// care to know when they are done. All that's important is whether I can or can't place
	// new work in the box. If I can't perform the send, then I know the box is full and the
	// worker is at capacity. At this time work would need to be discarded so things can keep
	// moving.
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		// use for a select case
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
		}
	}

	close(ch)

	// Here we create a buffered channel size 5. A single worker is then put to work. A for/range
	// is used for the channel receive. We then process everyone that's received.

	// We try to send 20 'paper's, this time  we use a 'select/case' statement. If the send
	// isnt received then the 'paper' is dropped. In the end we then close the channel. This
	// signals without data (as we've seen above) that there is no more work and the worker
	// is free to go once they've completed everything they were assigned.
}

// The BENEFIT of not having to worry about this guarantee and that's reduced or no latency
// in the communication. In Fan Out there is a buffer space for each employee that will be
// sending a report. In Drop the buffer is measured for capacity and if capacity is reached
// work is dropped so things can keep moving. The COST is that the lack of guarantee , so make
// sure this doesn't pose any issues to the logic of my system.

// SIGNAL WITH DATA - DELAYED GUARANTEE - Buffered Channel of 1
// Now we have the case of more than 1 task. When you need the guarantee that a task has been
// received before sending a new signal. WAIT FOR TASKS NOTE the task vs tasks here.

func waitForTasks() {
	// So this time we have a new worker and they have more than one task to handle, so we are
	// going to feed them tasks, one after the other. The must though finish each individual task
	// before starting the next. Giving them more than one at a time would cause even more latency
	// issues if they try to asynchronously handle 2 at a time.

	// The benefit of buffer of 1 is that everything runs as expected and we don't wait for each
	// other as nothing is blocked, when I put more work in, the buffer is empty and when the worker
	// needs a new taks he takes it from the buffer, etc. The best part being that if we try to send
	// more work and we can't, then we know the buffer is still full, o sea the worker is having issues
	// with its current task and hasn't moved on to the next one. This is where the delayed guarantee
	// comes into play. When the buffer is empty and we can send that is the guarantee that the worker
	// has received the previous and working it, but when we can't send, its guaranteed that they
	// haven't.

	ch := make(chan string, 1) // buffer of exactly 1

	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
	}

	close(ch)
}

// SIGNAL WITHOUT DATA - CONTEXT
// This last scenario shows how to cancel a running goroutine with the context pkg. The reason
// this works is by leveraging an unbuffered channel that is closed to perform a signal without
// data.

func withTimeout() {
	// here once again we are a manager with a single worker. This time we aren't willing to
	// wait around for some uknown moment in time for the worker to finish. We have a defined
	// deadline and if the worker doesn't finish, then we cancel him then and there.

	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // should always follow your creation of a context

	ch := make(chan string, 1) // why are we buffering here? The use of the buffered channel
	// is an important piece to this algo. If the worker doesn't finish in time, we move on 
	// without them WITHOUT GIVING THEM ANY NOTICE. So from their perspective they still need to
	// send the result and they are blind to the fact if you receive it or not. If we use 
	// unbuffered then the worker would block forever trying to send a result with noone to
	// receive it. The creates a GOROUTINE LEAK. So why a buffer of 1? The buffered channel 
	// prevents this memory leak.

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)
	case <-ctx.Done():
		fmt.Println("moving on")
	}
}

// IN SUMMARY;
// LANGUAGE MECHANICS;
// 1. Use channels to orchestrate and coordinate goroutines
// 		- Focus on the signaling attributes and not the sharing of data
// 		- Signaling with data or without data
// 		- Question their use for synchromizing access to shared state
// 			- There are cases where channels can be simpler for this but initially question it
// 2. Unbuffered Channels:
// 		- Receive happens before the send
// 		- Benefit: 100% guarantee the signal has been received
// 		- Cost: Unknown latency on when the signal will be received
// 3. Buffered channels:
// 		- Send happens before the receive
// 		- Benefit: Reduce blocking latency between signaling
// 		- Cost: No guarantee when the signal has been received
// 			- The larger the buffer, the less the guarantee
// 			- Buffer of 1 can give you one delayed send of guarantee
// 4. Closing channels:
// 		- Close happens before the receive (like buffered)
// 		- Signaling without data
// 		- Perfect for signaling cancellations and deadlines (timeouts)
// 5. nil channels:
// 		- Send and Receive are blocked
// 		- Turn off signaling
// 		- Perfect for rate limiting or short term stoppages

// DESIGN PHILOSOPY:
// 1. If any given Send on a channel CAN cause the sending goroutine to block:
// 		- Not allowed to use a Buffered channel larger than 1
// 			- Buffers larger than 1 must hav reason/measurements
// 		- Must know what happens when the sending goroutine blocks
// 2. If any given Send on a channel WON'T cause the sending goroutine to block:
// 		- You have the exact number of buffers for each send
// 			- Fan Out pattern, for example. 
// 		- You have the buffer measured for max capacity
// 			- Drop pattern
// 3. Less is more with buffers
// 		- Don't think about performance when thinking about buffers
// 		- Buffers can help to reduce blocking latency between signaling
// 			- Reducing blocking latency towards zero does not necessarily mean better throughput
// 			- If a buffer of one is giving you GOOD ENOUGH throughput then keep it.
// 			- Question buffers that are larger than one and measure for size
// 			- Find the smallest buffer  possible that provides GOOD ENOUGH throughput
