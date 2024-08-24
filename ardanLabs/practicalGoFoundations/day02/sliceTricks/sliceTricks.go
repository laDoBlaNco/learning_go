package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("---------------Using Append and Copy for diff functionalities on slices ---")
	// so apparently there was a container/vector package that was removed  that had
	// a lot of trics that could be done on slices. Using just builtin append and copy
	// you can actually replicate those in a very idiomatic Go fashion.

	fmt.Println("Appending another slice:")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{11, 12, 13, 14, 15}
	fmt.Println(a, b)
	fmt.Println(append(a, b...))
	fmt.Println()

	//the thing with copy is remembering that the left  most arg is the dest.
	fmt.Println("Using Copy:")
	c := make([]int, len(a))
	copy(c, a)
	fmt.Println("a:", a, "\nc:", c)
	fmt.Println()
	// we can also append 1x1 to copy to do it in one line, but its a little slower
	// and not as effiencient, unless you are also appending other items after copy
	c = append([]int(nil), a...) // I've never used []int(nil), basically conversion nil
	// to an int slice
	fmt.Println("a:", a, "\nc:", c)
	fmt.Println("We created a nil slice with conversion:", []int(nil))

	// Never seen this before; a[:0:0]
	fmt.Println("Using [:0:0] slicing:", append(c[:0:0], b...))
	fmt.Println(c) // based on the experiments below using the [::] puts a limit on
	// the cap of the new slice so I don't overwrite the original underlying array
	// so a new underlying array is created, thus simulating copying an array. So
	// basically here we sliced off a NEW array from the original and filled it with
	// the needed elements.
	/*
		Explanation of the above by B Kennedy in Ardan labs:
		With the release of Go 1.2, slices gained the ability to specify the capacity
		when performing a slicing operation. This doesn’t mean we can use this index
		to extend the capacity of the underlying array. It means we can create a new
		slice whose capacity is restricted. Restricting the capacity provides a level
		of protection to the underlying array and gives us more control over append
		operations.

		Taking a slight detour into the Ardan Labs blog with to understand the 3
	*/
	fmt.Println()
	fmt.Println("Playing a bit with InspectSlice to understand [:0:0]:")
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	fmt.Println("Inspecting our source slice we created:")
	InspectSlice(source)
	// startign with a slice of len and cap 5. meaning we have access to the entire
	// underlying array
	fmt.Println("Slice of slice source:")
	takeOne := source[2:3]
	InspectSlice(takeOne)
	// here we only took the 3rd element, in face we can see that the address
	// of the third element is the first in takeOne. so working with same underlying
	// array. It has a len 1 and cap 3 cuz there 3 elements left in the underlying
	// array that are available to use.

	// What if we didn't want the new slice to have access to the remaining cap?
	// This is where [::] comes into play.
	fmt.Println("Slice of slice source with third index (Capacity):")
	takeOneCapOne := source[2:3:3] // use the 3rd index to set cap of new slice
	InspectSlice(takeOneCapOne)    // so we could append, etc without overwriting
	// the underlying array elements
	// if we try to make the new cap great than underlying array Go panics
	// takeOneCapFour:=source[2:3:6] // (6-2) attempts to set cap to 4 which is > avail
	// InspectSlice(takeOneCapFour) // slice bounds error

	// Now getting back to the slice tricks, if we append something to the new
	// slice that doesn't have access to underlying capacity?
	fmt.Println("Appending to a slice with new cap limits, creates a new underlying array:",
		"This protects your elements in original array from being overwritten")

	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	InspectSlice(takeOneCapOne) // it works and under the hood as we can see from
	// the mem addresses, When we append an element to takeOneCapOne, a NEW underlying
	// array is created for the slice (space complexity). This new array contains
	// a copy fo the elements being ref'd from the source and then it is extended as
	// normal.

	// How is this different from not using capacity limit?:
	fmt.Println("When we don't use the slice cap limit, we overwrite our underlying",
		"array:")
	InspectSlice(takeOne)
	takeOne = append(takeOne, "Kiwi")
	InspectSlice(takeOne)
	InspectSlice(source)

	// So this time it uses the existing cap it has access to and overwrote 'banana'
	// So this feature as of Go 1.2 protects us and our data from unwanted overwrites
	// The more we leverage these builtin functions and runtime to handle these
	// types of ops the better.

	// -------------Back to Slice Trix---------------------------
	fmt.Println()
	fmt.Println("Returning to Go Slice Trix:")
	fmt.Println("Still working on -- Copy (third way):")
	// This one-liner is equal to the other make+copy but its slower as of 1.16
	// but again we are making an empty slice and appending to it the elements of a
	fmt.Println(append(make([]int, 0, len(a)), a...))

	// What about cutting from a slice
	fmt.Println()
	fmt.Println("Cutting from a slice:")
	//Here we append using slices of slices to basically append to a slice starting
	// on top of the elements we want to cut out. Cutting out the '3'
	fmt.Println(append(a[:2], a[3:]...))
	fmt.Println()
	// Now we can also delete, which I think would be the same as above in essence
	fmt.Println("Deleting an element:")
	//
	fmt.Println(append(a[:3], a[3+1:]...))
	//or
	fmt.Println(a[:3+copy(a[3:], a[3+1:])])
	fmt.Println(a) // these are deleting, but they are leaving extra elements on the
	// end. also using Copy is changing the underlying array

	// From here I've decided to just touch on those trix that I see most practical.
	// I can always come back and check for some other more obscure ones later.

	// Resetting my slices after getting changed with the cutting and deleting.
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b = []int{11, 12, 13, 14, 15}

	fmt.Println()
	fmt.Println("Inserting in the middle of a slice:")
	fmt.Println(append(a[:4], append(make([]int, 5), a[4:]...)...))

	fmt.Println()
	fmt.Println("Here we extend our slice with empty slots:")
	fmt.Println(append(a, make([]int, 50)...))

	fmt.Println()
	fmt.Println("Here we have an in-place filter:")
	n := 0
	for _, x := range a {
		if x%2 != 0 {
			a[n] = x
			n++
		}
	}
	a = a[:n]
	fmt.Println(a)
	fmt.Println()

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b = []int{11, 12, 13, 14, 15}

	fmt.Println("Above we inserted empty slots in a slice, This time a slice in a slice:")
	fmt.Println(append(a[:5], append(b, a[5:]...)...))

	fmt.Println()
	fmt.Println("Here we are going to implement Push, which is normal append:")
	fmt.Println(append(a, b...))

	fmt.Println()
	fmt.Println("Here we are implementing Pop:(note we're assigning a var as pop",
		"normally returns the popped element)")
	x := a[len(a)-1]
	fmt.Println(x, a[:len(a)-1])

	fmt.Println()
	fmt.Println("Now lost try from the Front with Unshift and shift:")
	fmt.Println(append([]int{69}, a...))
	x = a[0]
	fmt.Println(x, a[1:])

	fmt.Println()
	fmt.Println("We'll finish of two ways of reversing our slice:")
	fmt.Println("First way:")
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)
	fmt.Println()
	fmt.Println("Second way: (back to normal)")
	for l,r:=0,len(a)-1;l<r;l,r=l+1,r-1{
		a[l],a[r]=a[r],a[l] 
	}
	fmt.Println(a) 
	
	fmt.Println("Aqui termino por ahora.") 


}

// This is the 'InspectSlice' function from Ardan Labs to help us see what's
// happening.
func InspectSlice(slice []string) {
	// Capture the address to the slice structure
	address := unsafe.Pointer(&slice)

	// Capture the address where thel ength and cap size is stored
	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	// Create pointers to the length and cap size
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr [0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for i := 0; i < *lenPtr; i++ {
		fmt.Printf("[%d] %p %s\n",
			i,
			&(*addPtr)[i],
			(*addPtr)[i])
	}
	fmt.Printf("\n\n")

}
