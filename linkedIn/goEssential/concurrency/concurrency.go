package main

import (
	"fmt"
	"net/http"
	// "sync"
	"context"
	"log"
	"strconv"
	"time"
)

var (
	urlTemplate = "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	colors      = []string{"green", "yellow"}
)

func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))

}

type result struct {
	url  string
	size int
	err  error
}

func sizeWorker(url string, ch chan result) {
	fmt.Println(url)
	res := result{url: url}
	res.size, res.err = downloadSize(url)
	ch <- res
}

func main() {

	start := time.Now()
	size := 0
	ch := make(chan result)
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplate, color, month)
			fmt.Println(url)
			go sizeWorker(url, ch)
		}
	}

	for i := 0; i < len(colors)*12; i++ {
		r := <-ch
		if r.err != nil {
			log.Fatal(r.err)
		}
		size += r.size
	}

	duration := time.Since(start) // this seems like a Go idiomatic timer
	fmt.Println(size, duration)

	// using select to working with multiple channels at the same time
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select { // much like the switch statement, The case chosen depends on which ch has data
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	// Using Select for timeouts
	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(20 * time.Millisecond): // behind the scenes time.After is creating another gr
		fmt.Println("timeout") // and then after sleeping sending a val to the channel. That's why
		// that <-time.After works.
	}

	// So it seems like Context is used in place of in or timeout situations more formally
	// The context is the constraints  or what the world looks like and how we should
	// respond

	// Here we had time to get the actual bid
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "https://http.cat/418"
	bid := findBid(ctx, url)
	fmt.Println(bid)

	// here with only 10 milliseconds we timedout and had to use the default bid
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url = "https://http.cat/404"
	bid = findBid(ctx, url)
	fmt.Println(bid)

	/*
		//CHALLENGE - Using channels instead of sync wait groups
		urls := []string{
			"https://go.dev",
			"https://api.github.com",
			"https://httpbin.org/ip",
		}

		// create response channel
		ch := make(chan string)
		for _, url := range urls {
			go returnType(url, ch) // instead of an anony func we simply 'go' our func
		}


		for range urls {// This is the first time I see this with no vars. Apparently you
			out:=<-ch   // can do for range if you aren't using any of the vars k or v, i or val, etc
			fmt.Println(out)
		}
	*/
	/*
		start := time.Now()
		siteSerial(urls)
		fmt.Println(time.Since(start))

		start = time.Now()
		sitesConcurrent(urls)
		fmt.Println(time.Since(start))
	*/
	/*
		// Working with channels
		ch := make(chan int) // using the allocation func to make a chan
		// ch <- 353            // sending to the channel

		go func() {
			// send number to the channel
			ch <- 353
		}()

		val := <-ch // receiving from channel
		fmt.Printf("got %d\n", val)

		// Here we send several values through
		const count = 3
		go func() { // here we see with channels, no need for waitgroup. It blocks to
			for i := 0; i < count; i++ { // ensure we send/receive, send/receive,etc
				fmt.Printf("sending %d\n", i)
				ch <- i
				time.Sleep(time.Second)
			}
		}()

		for i := 0; i < count; i++ {
			val := <-ch
			fmt.Printf("received %d\n", val)
		}

		// Also we can close a channel. The only reason for this would be to signal
		// to the other end that nothing else is coming through. We would need to do this
		// if we wanted to use a for range loop and need to know if and where it ends.
		go func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("sending %d again\n", i)
				ch <- i
				time.Sleep(time.Second)
			}
			close(ch) // we close the channel when loop is finished
		}()

		// now we can do the same as above but with range, and it just works!
		for i := range ch { // range of the actual channel
			fmt.Printf("received %d again\n", i)
		}

		// Finally we have buffered channels meaning we can load channels up to a certain
		// point before they block. They are still receiving one piece of data at a time
		ch2 := make(chan int, 1) // buffered channel, I can send 1 value without blocking regardless
		ch2 <- 19
		val2 := <-ch2
		fmt.Println(val2) // this is like our first example, except here it doesn't block
		// so even though we are only using the main routine, we can send and receive
		// since it doesn't block.💪🏽
	*/
}

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype) // Note I don't have to return. I
	//put them in the channels from inside my function
}

/*
func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup // I think we're using wg since we aren't using channels yet
	for _, url := range urls {
		wg.Add(1)
		go func(url string) { // my first go routine spun with anony func
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
*/

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: "https://adsʁus.com/19",
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "https://adsʁus.com/default",
	Price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered to avoid gr leak - a memory leak of leaving a gr hanging out there
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done(): // ctx.Done also sends a channel and if received it means we didn't get bid
		return defaultBid
	}
}
