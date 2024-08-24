package main

import (
	"fmt"
	"time"
)

// ...for quick debugging
var p = fmt.Println

func main() {

	// We start by getting the current time
	now := time.Now()
	p(now)

	// time is a struct. We can build one by providing year,month,day, etc. Times
	// are always associated with a location, i.e. 'time zone'
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	fmt.Printf("%#v\n", then)

	// we can extract each component of a go time
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// we also have the weekday string available
	p(then.Weekday())

	// We can start to compare two times with .Before, .After, and .Equal
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub method returns the difference (subtraction) or duration between
	// two times. (in 24h24m24.24s format)
	diff := now.Sub(then)
	p(diff)

	// We can then get the duration in different units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// with .Add we can advance a time by a given duration (or use - to move
	// backwards by a duration)
	p(then.Add(diff)) // this didn't come print as I expected cuz Go takes into
	// consideration your local time zone unless you tell it differently.
	p(then.Add(-diff))

}
