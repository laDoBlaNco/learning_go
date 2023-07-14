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
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

// DON'T DO IT THIS WAY IN PRODUCTION:
// ITS HARD TO READ
// BUT ITS JUST AN EXERCISE

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Requires age")
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf("Wrong age: %q\n", os.Args[1])
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}

/* This was my solution. Same result, but Inanc solution combined a branch
so there's 1 less branch, which is better.

func main() {

	if age := os.Args[1:]; len(age) != 1 {
		fmt.Println("Requires age")
	} else if n, err := strconv.Atoi(age[0]); err != nil {
		fmt.Printf("%s can't be converted to a number", age[0])
	} else if n < 0 {
		fmt.Printf("Wrong age: %q", age[0])
	} else if n < 13 {
		fmt.Println("PG-Rated")
	} else if n <= 17 {
		fmt.Println("PG-13")
	} else if n > 17 {
		fmt.Println("R-Rated")
	}
}

*/
