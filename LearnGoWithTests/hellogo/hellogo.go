package main

import (
	"fmt"
)

// ...for quick debugging
var pl = fmt.Println

func Hello() string {
	return "Hello, Gophers!"
}

func main() {

	fmt.Println(Hello())

}
