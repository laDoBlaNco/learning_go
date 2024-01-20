package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Waitgroups allow us to block the execution of the application until a series of goroutines
// have finished their execution. this is necessary when we have a lot of work to do and we
// split it amongst various routines and we need to make sure that all the routines complete
// and return before the application finishes.

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://tutorialedge.net",
}

func fetch(url string, wg *sync.WaitGroup) { // and make sure our fetch accepts the pointer
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	wg.Done() // will decrement the counter by 1 taking us back down to 0 as we  complete.
	// its like putting stuff on and then checking them off a list on a clipboard as we
	// manage our goroutines.
}

func crawl() {
	var wg sync.WaitGroup // we create our wg var of type sync.WaitGroup

	for _, url := range urls {
		wg.Add(1) // here we update the wg everytime we add a new url to our goroutine
		// this increments the internal wg counter which will continue to block until that counter
		// returns to 0. So we need a way to block and we do that below with wg.Wait in th emain
		// routine
		go fetch(url, &wg) // we then add a reference (pointer) to that wg to our fetch func
	}

	wg.Wait() // once the counter is at zero, this will stop blocking the completion of our app
	fmt.Println("finished crawling urls")
}

// Note that without the waitgroup, this func will terminate before the goroutines have
// a chance to complete their work.
func main() {
	crawl()
}
