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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me a year number")
		return // without the return it continues to run the other if statements
		// in main.
	}

	y, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a valid year", args[0])
		return
	}

	if y%100 == 0 {
		if y%400 == 0 {
			fmt.Printf("%d is a leap year", y)
		} else {
			fmt.Printf("%d is not a leap year", y)
		}
	} else if y%4 == 0 {
		fmt.Printf("%d is a leap year", y)
	} else {
		fmt.Printf("%d is not a leap year", y)
	}
}

// main difference from Inanc's solution is that he used less print statements
// by using a bool result with an if to check and see if the year was leap
// or not first. Then "if leap {fmt.Printf.....}"
