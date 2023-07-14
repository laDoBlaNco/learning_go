package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// siteTime("https://google.com")
	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://no-such-site.biz",
	}
	// as in the last file when I send a goroutine to do something, its like sending out a
	// little worker off into the wild with a task. If I don't wait for them to get back or I
	// finish what I'm doing an walk away, then I'll never get anything. This is where WaitGroups
	// come into play. We could use a channel with an empty {} or something just to see if something
	// got done, but again, Waitgroups are a better solution.
	var wg sync.WaitGroup // its basically a fancy counter
	wg.Add(len(urls))     // we can create the full count right at the beginning
	for _, url := range urls {
		// wg.Add(1) // or add 1 for every goroutine spun.
		url := url
		go func() {
			defer wg.Done() // this then reduces the counter by 1, o sea, checks it off the WaitGroup list
			siteTime(url)
		}()
	}
	// and finally we set up the Wait which will not be  called until everything on its list is
	// complete.
	wg.Wait()
} // a similar pkg is 'errgroup' which does basically the same but focuses on returning errors.

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}
