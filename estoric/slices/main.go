package main

import (
	"fmt"
)

func main() {
	// Create a slice using the make function
	// mySlice := make(<type>,<len>,<cap>) cap is optional
	// mySlice := make([]string, 3, 8) // type of a slice []<type>

	// Another way to create a slice is with a slice literal
	// mySlice := []string{"apple", "orange", "pear"}
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	mySlice = append(mySlice, 8) // note what this does to capacity

	mySlice[2] = 32 // remember that this alters the underlying array
	fmt.Printf("My length is: %d\nMy Capacity is: %d\n",
		len(mySlice), cap(mySlice))
	fmt.Printf("My values are: %v\n", mySlice)

	newSlice := mySlice[:4] // here we use  the [:] to create a new slice of slice
	fmt.Println(newSlice)

	// for INDEX,VALUE in range
	for _, i := range mySlice {
		fmt.Println(i)
	}

	// Appending one slice to another:
	mySlice = append(mySlice, newSlice...) // notice the need for ... elipses
	fmt.Println(mySlice)

}

/*
Go looks at arrays differently then other languges. But an array and slice
are lists of items that are all of the same type and each  itemis indexed (0 based)

in Go you are going to work with slices more than arrays but they are binded
together so we need to understand what Arrays are in Go

Array - fixed length. It can't be changed. To add to it you have to create a new
array and copy the items to it.

Slice - flexible length. You can change the length by either appending to it
or getting a slice of a slice.

A slice is built on top of an array. Thinking again of ThePrimeagen's explanation
of JS arrays vs actual Arrays. That's the same concept of slices vs arrays.

When you create a slice you are actually creating a 'backing array' and the slice
is reference that array. Its a representation of the array that stored in mem.
Its a Reference Type. Along with that it comes with props (starting point, Length,
and Capacity). We can't have a slice that exceeds the capacity. Remember we are
working on top of an underlying backing array. So the length is the size of the
slice and the capacity is the size of the underlying array (from the starting
index of the slice to the end of the array)

Using make to create a slice will not give it values, so it'll be created with
the default values based on its type.

using the slice literal it doesn't seem to allow  us to adjust capacity as we
are putting in a certain number of values from the start.

Few things to note at this point:
1. We are working with the underlying array so any changes are happening
to the array and will show in all slices built off of that array that include
the same indices
2. When we append to a slice we are placing more values on the end and if
needed Go will adjust the capacity to fit the new slice. This adjustment isn't
1:1 dependent on what you adding. i.e., adding 2 values doesn't mean Go is
going to simply add 2 more spots...when we use append and there isn't enough
space then it'll get rid of the original slice and create a new one with double
the capacity. That's why when we add 1 item to a 7 length slice it goes to 14
capacity. You need to keep this in mind  because it will impact memory if you
aren't aware. Every append doubles the capacity. But if we use the make function
then we set our on capacity.

*/
