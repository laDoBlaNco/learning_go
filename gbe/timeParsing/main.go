package main

import (
	"fmt"
	"time"
)

// ...for quick debugging
var p = fmt.Println

func main() {

	// go does time formatting based on pattern based layouts
	// we first have basic formatting based on the RFC3339 protocol
	// Go has a layout const for this
	t := time.Now()
	p(t.Format(time.RFC3339))

	// time.Parse using the same layout const
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// Format and Parse use example-based layouts. Usually you'll use a constant
	// from 'time' for these, but you can also supply custom layouts. Layouts must
	// must use Go reference time tokens 'Mon Jan 2 15:04:05 MST 2006' to show
	// Go the pattern with which to format/parse a given time/string. The example
	// time must be exactly as shown: the year 2006, 15 (or 3 12/hour time) for the
	// hour, Monday for the full weekday (or Mon), etc.
	// NOTE: the tokens fall in order from 0 - 7 with Mon being 0, Jan being 1,
	// 2nd being 2, 15(3) hour being 3, etc
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// For numeric representations you can also use standard string formatting with
	// the extracted components of the time value
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input explaining what the problem
	// might be
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err := time.Parse(ansic, "8:41PM")
	p(err) // remember that the parsing is based on pattern matching. So Go is using
	// the space delimiter as well to determine what you are trying to do. In this
	// example it sees "Mon" token and "8:41PM" to fill it, which it can't do.

}
