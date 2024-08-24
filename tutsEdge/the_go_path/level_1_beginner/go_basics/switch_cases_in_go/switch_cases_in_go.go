package main

import "runtime"

func main() {
	println("If statements in Go")
	var customerHeight int = 140
	customerAge := 18

	switch { // switch statements are preferred when we have more complex branches
	case customerHeight >= 150 || customerAge >= 18:
		println("can access ride")
	case customerHeight >= 120:
		println("can access children's rides")
	default:
		println("cannot access rides")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("os x")
	case "linux":
		println("linux machine")
	default:
		println("something else")
	}

	// if customerHeight >= 150 || customerAge >= 18 {
	// 	println("can access ride")
	// } else if customerHeight >= 120 {
	// 	println("customer can access children's rides")
	// } else {
	// 	println("customer is to small")
	// }
}
