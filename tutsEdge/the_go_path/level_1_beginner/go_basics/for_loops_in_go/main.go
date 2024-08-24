package main

import "fmt"

func main() {
	println()
	println("For Loops in Go!")
	ages := []int{42, 28, 30, 27, 18}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}
	println()
	for i := 0; i < len(ages); { // basically a while loop with the intialization of i after for
		// rather than putting it outside the loop like we are below. same thing in two ways.
		fmt.Println(ages[i])
		i++
	}
	println()
	// since Go only has the 'for' loop we use it for everything. Below we have an infinite loop
	// which would need  condition to break. A while loop would be the same but with the condition
	// included after for. As you see we are just moving what's here i:=0;i<len(ages);i++ into the
	// body of the loop with the intialization of the var outside as with any while loop
	i := 0
	for i < len(ages) {
		fmt.Println(ages[i])
		i++
	}
	println()
	// Here's the infinite example:
	i = 0
	for {
		fmt.Println(ages[i])
		i++
		if i >= len(ages) {
			break
		}
	}
	println()

	// here we are going to use 'continue' instead of break so that we can skip certain elements
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// finally we have the for..range which will allow us to work with the index,value
	for index, value := range ages {
		fmt.Print(index, "...")
		fmt.Println(value)
	}

}
