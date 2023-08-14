package main

import "fmt"

/*
	SLICING SLICES - Slices provide the ability to avoid extra copies and heap allocations
	on the backing array when needing to isolate certain elements of the backing array 
	for different operations.
	
	The slicing syntax represents the list notation [a:b] which means, include elements from
	index a through b, without including b. 
*/

func main(){
	
	// Create a slice with a length of 5 elements and a capacity of 8
	slice1 := make([]string,5,8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	
	inspectSlice(slice1)  
	fmt.Println() 
	
	// Take a slice of slice1. We want just indexes 2 and 3.
	// parameters are [starting_index : (starting_index + lenght)] 
	slice2 := slice1[2:4] 
	// The var slice2 is a new slice value that is now sharing the same backing array that 
	// slice 1 is using. We can see this in the results of this example. However slice2 only
	// allows you to access the elements at index 2 and 3 ("Banana" and " Grape") of the original
	// slice's backing array. The length of slice2 is 2 and not 5 like slice1 and the capacity
	// 6 since there there are only 6 elements left in that backing array from that pointer
	// position.
	// A better way to think about slicing is to focus on the length using this notation
	// [a:a+len] index a through a plus the length of your new slice. This will reduce 
	// errors in calculating new slices. 
	inspectSlice(slice2) 
	
	// as mention after running this example we NOTE how the slice2 is sharing the same backing
	// array. We can see this by comparing the addresses. The nice thing about this is there 
	// are no allocations. The compiler knows the size of the backing array for slice1 at 
	// compile time. Passing a copy of the slice value down into the inspectSlice function
	// keeps eveyrthing on the stack.
	
	// MUTATIONS IN THE BACKING ARRAY
	// When we use slice2 to change the value of teh string at its index 0, any slice value that
	// is sharing the same backing array (where the address for that index is part of that slice's
	// length) will see the change. 
	fmt.Println() 
	slice2[0] = "CHANGED"
	inspectSlice(slice1)
	inspectSlice(slice2) 
	// We need to always be aware when we are modifying a value at an index position if the 
	// backing array is being shared with another slice. 
	
	// But we could just use append right? Wrong
	fmt.Println() 
	slice2 = append(slice2,"CHANGED")
	inspectSlice(slice1)
	inspectSlice(slice2) 
	// All it does is move the problem, mainly because at no point are we creating a new backing
	// array and therefore any change, whether  on the slice length or appending will overwrite
	// the backing array if its falls on an element already in use. 
     
    // The way to get out of this issue is to use the 3rd arg for slicing which sets the cap 
    // as well as the length. If we tell Go that we want to create a slice of a slice, but with
    // the same length and Cap, then we append  to it, Go will determine that its the same len & cap
    // and will create a new backing array for your slice, thus removing the mutation from the
    // original backing array.
    fmt.Println() 
   	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	
	slice2 = slice1[2:4:4] 
	slice2 = append(slice2,"Changed") 
	inspectSlice(slice1)
	inspectSlice(slice2) 
	// and there we go, a new backing array away from danger. 
	
	// Here is another example of using the third index slice
	fmt.Println()
   	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	inspectSlice(slice1) 
	fmt.Println()
	
	// take a slice of a slice again, just index 2
	takeOne := slice1[2:3]
	inspectSlice(takeOne) 
	fmt.Println()
	
	// Take a slice of just index 2 with a length and capacity of 1
	takeOneCapOne := slice1[2:3:3] // using the third index position to set the cap to 1 o sea
	// keep it the same as the length, whatever that is
	inspectSlice(takeOneCapOne) 
	fmt.Println()
	
	// now we can append a new element without hurting the backing array
	takeOneCapOne = append(takeOneCapOne,"Kiwi") 
	inspectSlice(slice1) 
	inspectSlice(takeOneCapOne) 
	
                	
}

func inspectSlice(slice []string){
	fmt.Printf("Length[%d]  Capacity[%d]\n",len(slice),cap(slice))
	for i := range slice{ // this is the pointer semantic version
		fmt.Printf("[%d]  %p  %s\n",i, &slice[i],slice[i]) // note the 16 byte stride
	}
}
