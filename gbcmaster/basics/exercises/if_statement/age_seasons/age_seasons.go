// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"
var p = fmt.Println

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func main() {
	// Change this accordingly to produce the expected outputs
	age := 61

	// Type your if statement here.
	if age>60{
		p("Getting older")
	}else if age>30{
		p("Getting wiser")
	}else if age>20{
		p("Adulthood")
	}else if age>10{
		p("Young blood")
	}else{
		p("Booting up") 
	}
}
