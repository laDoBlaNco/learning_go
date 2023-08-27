// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

/*
// A Duration represents the elapsed time between two instants as
// an int64 nanosecond count. The representation limits the largest
// representable duration to approximately 290 years.

type Duration int64

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.

const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)

// Add returns the time t+d.
func (t Time) Add(d Duration) Time
*/

// Sample program to show how literal, constant and variables work
// within the scope of implicit conversion.
package main

import (
	"fmt"
	"time"
)

func main() {
	
	// Here we see various examples of using constants in real-word.
	// Use the time package to get the current date/time.
	now := time.Now()

	// Subtract 5 nanoseconds from now using a literal constant. Remember that a literal
	// constant is simply an untyped number in this case. a KIND of numeric value. In Go when
	// we talk about literal constants we are talking about any literal data that we create
	// without the use of a TYPE constructor or New function, etc. One small correction or 
	// better yet clarification. Because we are using these literal constants for MATHEMATICAL
	// operations (adding, subtracting, multiplying, etc) then some implicit conversion must be
	// done to ensure that they untyped literal matches what we are combining it to.
	literal := now.Add(-5)

	// Subtract 5 seconds from now using a declared constant. Meaning that we are going to 
	// declare the constant and give it a name. Its still an untyped constant (KIND) but now
	// with a name 'timeout' and of course being implicitly converted to a time.Duration
	// KIND
	const timeout = 5 * time.Second // time.Duration(5) * time.Duration(1000000000)
	constant := now.Add(-timeout)

	// Subtract 5 nanoseconds from now using a variable of type int64. So here we are actually
	// putting our untyped literal KIND into a Typed 'box', which means it'll have the same
	// footprint or limitations of the type we are using 'int64'.
	// minusFive := int64(-5)
	// variable := now.Add(minusFive) 

	// example4.go:50: cannot use minusFive (type int64) as type time.Duration in argument to now.Add

	// Display the values.
	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Literal : %v\n", literal)
	fmt.Printf("Constant: %v\n", constant)
	// fmt.Printf("Variable: %v\n", variable)
}
