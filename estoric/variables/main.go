package main

import "fmt"

// Use two forward slashes like this to enter comments in code

/*
This is a multi-line comment
*/

/*
VARIABLES AND ALL THE BASIC PRIMITIVE TYPES IN GO

Variables are used to store values in memory
Go is a Statically Type language. The data type can't be changed after declaring it, though the value can.

A byte ------> 8 bits
A bit is the smallest unit of data store in a computer

Integers can be signed or unsigned

Unsigned means its a postive number
uint8		unsigned 8-bit integers (0 to 255)
uint16		unsigned 16-bit integers (0 to 65535)
uint32		unsigned 32-bit integers (0 to 4294967295)
uint64		unsigned 64-bit integers (0 to A REALLY BIG NUMBER)

Signed means its a negative or positive number
int8		signed 8-bit integers (-128 to 127)
int16		signed 16-bit integers (-32768 to 32757)
int32		signed 32-bit integers (-2147483648 to 2147483647)
int64		signed 64-bit integers (A REALLY BIG NEGATIVE NUMBER TO A REALLY BIG POSITIVE NUMBER)

uint		unsigned, either 32 or 64 bit value
int		signed, either 32 or 64 bits

float32		32-bit floating-point numbers (1.23)
float64		64-bit floating-point numbers

byte		alias for uint8 (also used for chars)
rune		alias for int32 (also used for unicode and larger byte chars)

bool		true or false
strings		"Strings look like this and are surrounded in double quotes in Go."

*/

// SO HOW DO YOU CREATE A VARIABLE?

// Explicit Declaration
// We declar with - var | varableName | type
// Creates a location in memory that holds that type

// Location in memory that holds a string
var greeting string // In memory but without a value

func main() {
	// Store a value in our variable aka Initialization
	greeting = "What it is"
	fmt.Println(greeting)

	// Implicit Declaration
	// Declare and initilize a var using :=
	a := 30
	fmt.Println(a)
	year := "2022"
	myString := "We are almost ending "
	myBool := 5 > 8

	// We can also use the var with implicit dec as well
	var myFloat float32 = 5.8 // equal sign without :

	fmt.Println(myString + year)
	fmt.Println(myBool)
	fmt.Println(myFloat)
}

// Zero Default Value Type - Every type has a default value type:
// int = 0
// bool = false
// string = ""
// float = 0
// reference = nil
