package main

import (
	"errors"
	"fmt"
)

/*
	Create two error variables, one called ErrInvalidValue and the other called
	ErrAmountTooLarge. Provide the static message for each variable. Then write
	a function called checkAmount that accepts a float64 type value and returns
	an error value. Check the value for zero and if it is, return the ErrInvalidValue.
	Check the value for greater than $1,000 and if it is, return ErrAmountTooLarge
	Write a main function to call the checkAmount function and check the return error
	value. Display a proper message to the screen.

*/

var (
	// 1. Declare an error variable name ErrInvalidValue using the New function
	// from the errors package
	ErrInvalidValue = errors.New("You have an invalid value")

	// 2. Declare an error variable named ErrAmountTooLarge using the New function
	// from the errors package
	ErrAmountTooLarge = errors.New("Your amount is too large")
)

//  3. Declare a function name checkAmount that accepts a value of type float64
//     and returns an error interface value
func checkAmount(amt float64) error {
	// 4. Is the parameter equal to zero. If so then return the error variable
	if amt == 0 {
		return ErrInvalidValue
	}
	// 5. Is the parameter great than 1000. If so then return the other error variable
	if amt > 1000 {
		return ErrAmountTooLarge
	}
	// 6. Return nil for the error value
	return nil
}

func main() {

	// 7. Call the checkAmount function and check the error, then use a switch/case to compare
	//    the error with each variable. Add a default case. Return if there is an error
	err := checkAmount(9999)
	if err != nil {
		switch err {
		case ErrInvalidValue:
			fmt.Println(err)
			return
		case ErrAmountTooLarge:
			fmt.Println(err)
			return
		default:
			fmt.Println(err)
			return
		}
	}

	// 8. Display everything is good
	fmt.Println("Everything is good")
}
