package main

import "fmt"

/*
	DATA SEMANTIC GUIDELINE FOR STRUCT TYPES

	As a guideline for us, if the data we ar working is a struct type, that is not an internal
	type as we've seen before, but a struct type, then we have to think about what the data
	represents to make a decision. A good general rule is to ask if the struct represents data
	or an API. IF THE STRUCT REPRESENTS AN API USE POINTER SEMANTICS.

	Below is a Time struct from the time package. If we consider Time to rep data, value semantics
	should be used for this struct.

	type Time struct{
		sec int64
		nsec int32
		loc *Location
	}

	For example in the Time package we can see the following factory function Now()

	func Now() Time{
		sec,nsec ;= now()
		return Time{sec + unixToInternal, nsec, Local}
	}
	Look at the return, its using value semantics for Time values which means every function
	gets its own copy of a Time value and fields in a struct should be  declared as values of
	type Time.

	func (t Time) Add(d Duration) Time{
		t.sec += int64(d/1e9)
		nsec := int32(t.nsec) + int32(d%1e9)
		if nsec >= 1e9{
			t.sec++
			nsec -= 1e9
		}else if nsec < 0{
			t.sec--
			nsec += 1e9
		}
		t.nsec = nsec
		return t // here's our copy of the value being returned.
	}
	Again here we see that Add gets its own copy of a Time value for mutation. It mutates its
	copy, then it returns a copy back to the caller. Once again this si the safest way to perform
	a mutation operation.


	THERE ARE SOME EXCEPTIONS THOUGH:

	func (t *Time) UnmarshalBinary(data []byte) error {}
	func (t *Time) GobDecode(data []byte) error {}
	func (t *Time) UnmarshalJSON(data []byte) error {}
	func (t *Time) UnmarshalText(data []byte) error {}

	so what's the difference? They are all implementing an interface where the method signature
	is locked in. Since the implementation requires a mutation, pointer semantics are the only
	choice.

	NOTE: HERE IS THE GUIDLINE:
	If value semantics are at play, we can switch to pointer semantics for some functions as long
	as we don't let the data in the remaining call chain switch back to value semantics. That means
	once we switch to pointer semantics, all future calls from that point need to stick to our
	pointer semantics, We can NEVER, EVER ,EVER GO FROM POINTER TO VALUE.
	ITS NEVER SAFE  TO MAKE A COPY OF A VALUE THA A POINTER POINTS TO.

*/

// A good way to identify what data semantic was chosen, when we are looking at an existing
// codebase, is to look for a factory function. The return type of a factory function should
// dictate the data semantic chosen by the dev.

// Here's some example code to show how to declare methods against named types (structs)

// duration is a named type that represents a duration of time in Nanosecond
type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

// setHours sets the specified number of hours
func (d *duration) setHours(h float64){
	// here we are using pointer semantics and not returning anything as this is a mutation
	*d = duration(h) * hour
}

// hours returns the duration as a floating point number of hours
func(d duration) hours()float64{
	hour:=d/hour
	nsec:=d%hour
	return float64(hour) + float64(nsec)*(1e-9/60/60) 
}


func main() {

	// declare a v ar of type duration set to its zero value
	var dur duration
	
	// change the value fo the dur to equal five hours
	dur.setHours(5) 
	
	// display the enw value of dur.
	fmt.Println("Hours:",dur.hours()) 

}
