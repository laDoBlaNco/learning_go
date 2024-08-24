package main

import (
	"fmt"
	"log"
	"net/http"
)

/*

the benefit of worker pools is normally to reduce the cost of spinning up threads. In Go this
doesn't really apply since the cost of spinning up a goroutine is not signicant. But it is valuable
because we can limit the number of concurrent prrocesses we have running at any one time. This
is great for io apps or operations and you have limitations in bandwidth, etc.

let's demonstrate this with our crawler app from before.
*/

type Site struct {
	URL string
}
type Result struct {
	URL    string
	Status int
}

func crawl(wID int, jobs <-chan Site, results chan<- Result) {
	// the syntax <-chan and chan<- is specifying whether these are send or recieve channels
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wID)
		resp, err := http.Get(site.URL) // using our http pkg get get the url which returns
		// a response and possible error.
		if err != nil {
			log.Println(err.Error()) // if it errors we are only logging, not returning, so
			// the result will still be sent to the chan, regardless
		}
		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}

func main() {

	fmt.Println("Worker Pools In Go!")
	println()

	jobs := make(chan Site, 3) // we create the chans that will be used to communicate
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ { // here we create the 'w' workers in our worker pool that will each
		// take go with a goroutine and take with them our chan connections of jobs and results.
		go crawl(w, jobs, results) // since we are using multiple workers, we need to create this
		// crawl func to CONTINUALLY listen for new workers to accept and process
	} // Note that I've only got 3 workers but 4 urls. Take note that one of the workers has to
	// do two tasks. interesting concept of worker pools working till the work is done.

	// Now we can start to send urls
	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= len(urls); a++ {
		result := <-results
		log.Println(result)
	}

}
