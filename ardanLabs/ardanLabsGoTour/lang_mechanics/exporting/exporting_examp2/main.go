package main

import(
	"fmt"
	
		"github.com/ladoblanco/ardanLabsGoTour/exporting/examples"
)

// This example shows how we won't be able to compile if we try to access an unexported 
// type direclty
func main(){
	counter := examples.alertCounter(10) 
	
	fmt.Printf("Counter: %d\n",counter) 
	
	// I get the error: exporting_examp2/main.go:12:22: alertCounter not exported by package counters
}
