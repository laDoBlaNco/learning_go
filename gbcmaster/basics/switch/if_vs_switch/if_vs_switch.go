package main

import (
	"fmt"
	_ "log"
	"os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Gimme a month name.")
		return
	}
	//
	// m := os.Args[1]
	//
	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// fmt.Println("Fall")
	// } else {
	// fmt.Printf("%q is not a month\n", m)
	// }

	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Sprint")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month\n", m)
	}

}
