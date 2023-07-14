package main

import(
	"fmt"
)

var pl = fmt.Println

// Go isn't limiting, it just wants your permission before it does a lot of things
// such as conversion. type(value) such as type conversion

// You can convert almost any value to another type as long as its convertible to
// that type.

func main(){
	speed:=100 // type is int
	force:=2.5 // type is float64
	// speed=speed*force  // this gives me an error. You can't use diff types together
	// we need to convert one of these vars
	speed = speed * int(force) 
	pl(speed) // but this result isn't correct. So we need to ensure we using the
	// correct var  to convert. int(force) truncated and we lost the .5 which takes
	// away accuracy (precision) from our number.
	pl(force,int(force)) // the conversion creates a new value so its doesn't actually
	// change the original value.
	
	// we also can't just convert speed, we need to  convert both speeds
	speed2:=100
	force2:=2.5
	// speed2=float64(speed2)*force // can't just convert the one speed as the assignment is still int
	speed2 = int(float64(speed2)*force2)  
	pl(speed2) 
}
