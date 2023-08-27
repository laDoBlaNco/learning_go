package main

import (
	"fmt"

	"github.com/ladoblanco/ardanLabsGoTour/exporting/toy"
)

func main() {

	// Use the New function from the toy package to create a vlaue of type toy
	theToy := toy.New("fun_bean", 5)

	// Use the methods from the toy value to set some "initializing" values
	// I put "initializing" in "" because I didn't have it that way on the first try. I misread
	// the instructions and had args in the New func to initialize. After re-reeading and making
	// the corrections, the methods below are actually initializing the data.
	theToy.UpdateOnHand(6)
	theToy.UpdateSold(9)

	fmt.Printf("Name: %s\n", theToy.Name)
	fmt.Printf("Weight: %doz\n", theToy.Weight)
	fmt.Printf("Inventory: %d\n", theToy.OnHand())
	fmt.Printf("Total Sold: %d\n", theToy.Sold())

}
