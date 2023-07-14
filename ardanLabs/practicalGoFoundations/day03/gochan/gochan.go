package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine") // we start goroutines with 'go'
	fmt.Println("main")         // we only see main because the main goroutine closes before my goroutine can end

	for i := 0; i < 3; i++ {
		// Fix 2; Use a loop body variable - this is was TutorialEdge always does
		i := i // here 'i' shadows 'i' from the for loop
		go func() {
			fmt.Println(i) // this is the shadow 'i'
		}()

		/* Fix 1: Use a parameter
		go func(n int){
			fmt.Println(n)
		}(i) // now 'i' is passed as a copy to n
		*/
		/* BUG - as I was explaining below. All goroutines use the 'i' for the for loop. This is a closure
		bug because the goroutine is a closure and it remembers its environment. But when it finishes its work
		and comes back its 'i' is no longer what it started with, its now 3.
		go func(){ // this has the original i from the loop
			fmt.Println(i) // i remember this warning, from TutorialEdge. When using concurrency you want to
			// never use the changing var directly in the loop as its changing and so your value in your routine
			// isn't stable.
		}()
		*/

	}

	time.Sleep(10 * time.Millisecond) // this allows us to stay open till the goroutine ends. We can't
	// control when the routine ends. You send it out and it does its work and comes back.

	shadowExample()

	// we can do two things on a goroutine channel
	ch := make(chan string)
	go func() {
		ch <- "hi" // We can SEND through or to
	}()
	msg := <-ch // We can RECEIVE from
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	// when using a range with a channel you need to tell go when it ends, or you will deadlock
	// again. So we need to 'close' the channel
	for msg := range ch {
		fmt.Println("got:", msg)
	}
	/* for/range is just doing this but with syntatic sugar
	for{
		msg,ok:=<-ch
		if !ok{
			break
		}
		fmt.Println("got:",msg)
	}
	*/

	msg = <-ch
	fmt.Printf("closed: %#v\n", msg) // if we try to receive from a closed channel we get its
	msg, ok := <-ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok) // if we try to receive from a closed channel we get its

	// zero default value.
	// ch <- "hi" // panic: send on closed channel

	// Channels are typed.

	// Let's try running my exercise:
	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))

}

// we'll practic now with concurrency and channels implementing the worst sorting algorithm
// there is

/*
For every value 'n' in values, spin a goroutine that will:
	- sleep 'n' milliseconds
	- send 'n' over a channel

In the function body, collect values from the channel to a slice and return it.

*/

func sleepSort(values []int) []int {
	c := make(chan int)

	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			// I had to convert 'v' to a time.Duration in order to multiply it to a time.Duration
			c <- n
		}()
	}

	// Another use of the for/range loop that I don't remember seeing before, its the same as:
	// for i:=0;i<len(value);i++{}
	// because are we not using the v:=range then we don't need to close the channel since we are
	// running this n times. 
	var res []int
	for range values { // basically just taking out the var creation
		n := <-c
		res = append(res, n)
	}
	return res
}

/*
	Channel Semantics:
	1. send & receive will block until opposite operation (*)
		- So we must have a way receive meaning someone else needs to be receiving
		The same routine can't send and receive in this sense. We can fix then with a
		goroutine.
		- Buffered channel has cap(ch) non-blocking send operations
	2. Channels are not queues. There's no sitting in the queue and waiting while you
	   move on. If no one receives then you are stuck until someone receives.
	3. Receive from a closed channel will return the zero default value without blocking.
		- So to know that you are getting an actual value vs a zero default, use comma,ok
	4. Sending to a closed channel will panic, so we need to know who the owner of the chan
		is since its only the owner that can close the channel.
	5. Closing a closed channel will also panic.
	6. Send/recieve to a nil channel will block forever, there is a reason for this and why
		its allowed, so we'll need to investigate that.
*/

func test() {
}

// So the main piece with goroutines is using 'go' to spin it up and beware of the closure bug.

func shadowExample() {
	n := 7
	{ // NOTE how I'm creating new scope just by putting {} and nothing else.
		n := 2
		//n=2 // this is assignement of the outer n. If we  comment out n:=2 (never shadow) then this would
		// simply change or reassign the outer n
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}

// a working with locks and mutexes has always been a difficult problem. Go can do it, but
// Go prefers CSP (Communicating Sequential Processes). Sending messages between Goroutines
// with Channels.
