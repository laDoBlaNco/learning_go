package main

import "fmt"

// Typically if you are working with Go, you want to keep you 'main' func as small as possible.
// its just the entrypoint, so it shouldn't be doing a whole lot. We get around this using funcs

func HelloWorld(name string, age, height int) {
	fmt.Println("Hello:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
}

// Go funcs can also return values, or multiple values
func AddTotal(a, b int) (int, int) { // note that you must use ()s to return multiple vals
	return a + b, a - b
}
func main() {
	println()
	fmt.Println("Functions in Go!")
	HelloWorld("ladoblanco", 46, 200)
	total, negativeTotal := AddTotal(2, 3)
	fmt.Println("Total:", total)
	fmt.Println("Negative Total:", negativeTotal)

}
