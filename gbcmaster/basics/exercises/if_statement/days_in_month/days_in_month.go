// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month")
		return // again if you don't put this 'return' here it'll panic
	}
	mo := os.Args[1]
	leap := false
	y := time.Now().Year()
	// y = 2020 // this way rather than changing computer clock I just change my
	// y var to test.

	if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
		leap = true
	}

	if strings.ToLower(mo) == "january" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "february" && leap {
		fmt.Printf("%q has 29 days.\n", mo)
	} else if strings.ToLower(mo) == "february" {
		fmt.Printf("%q has 28 days.\n", mo)
	} else if strings.ToLower(mo) == "march" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "april" {
		fmt.Printf("%q has 30 days.\n", mo)
	} else if strings.ToLower(mo) == "may" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "june" {
		fmt.Printf("%q has 30 days.\n", mo)
	} else if strings.ToLower(mo) == "july" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "august" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "september" {
		fmt.Printf("%q has 30 days.\n", mo)
	} else if strings.ToLower(mo) == "october" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else if strings.ToLower(mo) == "november" {
		fmt.Printf("%q has 30 days.\n", mo)
	} else if strings.ToLower(mo) == "december" {
		fmt.Printf("%q has 31 days.\n", mo)
	} else {
		fmt.Printf("%q is not a month.\n", mo)
	}

}
