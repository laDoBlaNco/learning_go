package main

import "fmt"

// Slices are an incredibly important data structure in Go. They formt he basis for how we manage
// and manipulate data in a flexible, peformant, and dynamic way. Its very beneficial for all Go
// devs to learn how slices work and how to use them.
// Slices are a 3-word structure which holds the same as an array plus the capacity
// Array = [*pointer][length]    Slice = [*pointer][length][capacity]
// The length of the slice represents the number of elements THAT CAN BE ACCESSED on the backing array
// The capacity of the slice represents the TOTAL number of elements that EXIST on the backing array

// Slice construction

func main(){
	
	// we can construct a slice in several different ways.
	// Slice of string set to its zero default state - [*nil][""]
	var slice []string // note that there's nothing in [] this means its a slice
	fmt.Printf("%q\n",slice)
	
	// Slice of string set to its empty state - [*{}][""] ?
	slice2 := []string{} // 
	fmt.Printf("%q\n",slice2)
	
	// Slice of string set with a length and capacity of 5. Note the use of 'make'
	slice3 := make([]string,5) // if the cap isn't specified, its the same as the length
	fmt.Printf("%q\n",slice3)
	
	// Slice of string set with a length of 5 and capacity of 8.
	slice4 := make([]string,5,8)
	fmt.Printf("%q\n",slice4)
	
	// Slice of string set with values with a length and capacity of 5
	slice5 := []string{"a","b","c","d","e"}
	fmt.Printf("%q\n",slice5)
	
	// As we see here the built-in function 'make' allows us to pre-allocate both length and capacity for the 
	// backing array. if the compiler knows the size at compile time, the backing array can be constructed
	// on the stack. So by preallocating we could positively impact performance if needed.
	
	
	// Slice Length vs Capacity
	
	// The length of a slice represents the number of elements that can be read and written to. The capacity
	// represents the total number of elements that exist in the backing array FROM THAT POINTER POSITION
	// With syntactic sugar, slices look and feel just like arrays, but don't get them confused.
	slice6 := make([]string,5)
	slice6[0] = "Apple"
	slice6[1] = "Orange"
	slice6[2] = "Banana"
	slice6[3] = "Grape"
	slice6[4] = "Plum"
	
	fmt.Println(slice6)
	
	// We can tell the difference between a slice and an array construction, since an array has a known size
	// at compile time and slices necessarily don't [5]string vs []string
	// If we try to access an element beyond the length, even if its in the capacity, we'll get a runtim 
	// error or PANIC
	fmt.Println(slice6[5]) 
	
	
	
}
