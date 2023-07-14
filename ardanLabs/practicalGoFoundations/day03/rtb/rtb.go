package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Real-time Bidding:
// We are bidding on ads. We have a limited time to bid on an ad. The bestBid function checks
// a url and get the best bid and ad for that url. Our algo team created the algorithm, but we
// need to create a function bidOn which should always give a bid an never err. We have to send
// the bid within the allotted time. So we need create the context with this timeout, etc and
// the innards of the bidOn function.

func main() {
	// We have 50 msec to return an answer
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

// If algo didn't finish in time, return a default bid
func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid ,1) // so by adding the buffer of one, when this goroutine is working and things 
	// timeout, once it comes back with its  result it won't be blocked. It'll pass its value to the buffer
	// and once it does then it'll end. The channel will have the value in it, but if there are no references
	// to the channel (for example from an open goroutine), then the gc will pick it up.
	go func() {
		// bid := bestBid(url) or
		ch <- bestBid(url) 
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}

}
// I was right there with this one. Just didn't know if we were going to incorporate chans, etc
// as soon as I knew that's what he wanted, it came together super easy. So context looks like its
// mainly used with Select and Channels

// There is a bug here though. If this times out then the goroutine will still be waiting to 
// deliver its bid. It'll be blocked since Select moved on with its context for timeout. And since
// the routine is waiting to deliver its bid then the gc can't collect it. This is considered a 
// goroutine leak which is to me the same as a memory leak. Leaking goroutines. There are two 
// solutions: Select on the goroutine and time it out, or simply using a buffered channel. Remember
// that buffered channels is like a bank in the channel, so it'll take as much data as it can hold
// dependent on the capacity of the buffer, without blocking. So if we have 1 bid on that routine
// we can buffer 1 and even if select doesn't receive it, it'll still get delivered to the channel
// and not block. o sea NOTE: Buffered channel has cap(ch) non-blocking send operations and 90+% of 
// the buffered channels we'll see in Go are only 1, though in the wild, most channels are unbuffered.


var defaultBid = Bid{
	AdURL: "http://adsЯus.com/default",
	Price: 3,
}

// Written by Algo team, time to completion varies
func bestBid(url string) Bid {
	// Simulate work
	d := 100 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsЯus.com/ad17",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int // In ¢
}
