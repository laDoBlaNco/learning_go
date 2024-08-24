package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"log" 
)

var pl = fmt.Println

// Types and conversions

func main() {
	var vName string = "Ladoblanco"
	var v1, v2 = 1.2, 3.4
	var v3 = "Hello"
	v1 = 4.5
	pl(vName, v1, v2, v3)
	pl("Working on different types:")

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf('🦍'))

	// to cast
	cv1 := 1.5
	cv2 := int(cv1)
	pl(cv2)

	cv3 := "5000000"
	cv4, err := strconv.Atoi(cv3)
	pl(cv4, err, reflect.TypeOf(cv4))

	// integer to a string
	cv5 := 5000000
	cv6 := strconv.Itoa(cv5)
	pl(cv6)

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
		pl(cv8)
	}

	cv9 := fmt.Sprintf("%f\n", 3.14)
	pl(cv9)

	// If conditional
	// conditional ops: > < >= <= == 1=
	// Logical Ops: && || !

	age := 8
	if age >= 1 && age <= 18 {
		pl("Important Birthday")
	} else if age == 21 || age == 50 {
		pl("Important Birthdays")
	} else if age >= 65 {
		pl(" All of these are important")
	} else {
		pl("Not an important birthday")
	}

	pl("!true =", !true)

	// We've also got formatted output using these tokens:
	// Go's version of C's printf
	// %d: integer
	// %c: character
	// %f: float
	// %t: boolean
	// %s: string
	// %o: base 8
	// %x: base 16
	// %v: Guesses based on data type and uses Go's normal format for said type
	// %T: type of supplied value.

	pl()
	fmt.Printf("%s %d %c %f %t %o %x\n", "stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)

	// Now let's use this to solve a problem.
	// What is your age:
	// Receive customer data (Their age)
	// Google how to trim whitespace from input
	// Age < 5 Too you for school
	// Age == 5 Go to Kindergarten
	// Age > 5 or <= 17 Go to grade GRADE
	// Default Go to college.
	pl("Exercise 1 - Get user input:")
	pl()
	fmt.Print("What is your age: ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) 
	}
	ageData, err := strconv.Atoi(strings.TrimSpace(data))
	if err != nil {
		log.Fatal(err) 
	}
	// Now let's use the data
	if ageData < 5 {
		pl("Too young for school")
	} else if ageData == 5 {
		pl("Go to  Kindergarten")
	} else if ageData > 5 && ageData <= 17 {
		fmt.Printf("Go to grade %d", ageData-5)
	} else {
		pl("Go to college")
	}
}
