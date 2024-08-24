package main

import (
	"fmt"
	"time"
)

// CONTROL STATEMENTS
// IF ELSE, ELSE IF, FOR, SWITCH

func main() {
	x := 5
	if x < 10 {
		fmt.Println("x is less than 10") // simple example
	}
	fmt.Println("This will run no matter what")
	fmt.Println("====================================================")

	if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is not equal to 10") // with an else statement
	}
	fmt.Println("This will run no matter what")
	fmt.Println("====================================================")

	// Go doesn't put ()s around if or for statements.
	// Go has block scoping and we see this with if statements. Below a is in the if statement block.
	// so it can be accessed in that same block. But not in the Else block
	if x > 3 {
		a := x
		fmt.Println("x and a are larger than 3", a, x)
	} else {
		fmt.Println("a is not larger than 3")
		//fmt.Println(a) // this gives us an "undefined" error
	}
	fmt.Println("====================================================")

	if a := x - 3; a > 3 { // here we declare a as part of the if conditional and now its available in all branches
		fmt.Println("x and a are larger than 3", a, x)
	} else {
		fmt.Println(a)
	} // but its still not accessible outside of the if/else statement

	fmt.Println("====================================================")
	// we also have the else if statement
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("====================================================")

	// We can also talk about the for statements which are very similar to other languages, but with no ()s
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("====================================================")
	// Here's another example of a for loop which is more practical using a go slice
	names := []string{"Sam", "Tom", "Joe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("====================================================")
	// Another thing with for loops is the use of the break statement, if you need to leave the loop early
	// for example after you find what you are looking for.
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
		if i == 1 {
			break
		}
	}

	fmt.Println("====================================================")
	// There are also times when don't want to break the loop but simply skip over an iteration. For that
	// we use 'continue'
	for i := 0; i < len(names); i++ {
		if names[i] == "Sam" {
			continue
		}
		fmt.Println(names[i])
	}
	fmt.Println("====================================================")
	// Go decided not to use a while statement and just use the for for that as well
	i := 0
	for i < 5 {
		fmt.Println(i)
		i += 1 // just like a while statement, without this altering the counter, we'll run forever
	}

	fmt.Println("====================================================")
	// Go also uses 'for' as the forever loop. A simple empty for statment (for{}) will run forever
	i = 0
	for {
		fmt.Println(i)
		i += 1
		if i > 5 { // in a forever loop we need to break at some point
			break
		}
	}

	fmt.Println("====================================================")
	// Finally we have the 'for range' loop which gives us two values to work with in the loop as the range
	// returns either the index & value or the property & value depending if you are looping through a slice,
	// a map, a struct, etc
	for k, v := range names { // here we use k,v (key/value) where key is the index
		fmt.Println(k, v)
	}
	fmt.Println("====================================================")
	// Now let's talk about the switch statement. If you find yourselve with too many branches of if else
	// then a switch is more practical. Its works from top to bottom and stops once it finds a match.
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is a weekday.")
	}

	fmt.Println("====================================================")
	switch hour := time.Now().Hour(); { //note here we are just checking each case against 'true'.
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
	// Note that there is no switch expression only an assignment. If this is the case we must end the assignment
	// as we do in the if or for statements with an ;, and then Go sees there is no expression and will then knows
	// that we want it to compare the case against 'true'
	fmt.Println("====================================================")
	// Go again only runs the case that matches. This way no 'break' is needed. But if you need it to fall through
	// multiple cases we just use the 'fallthrough' keyword.
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

}
