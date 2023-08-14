package main

/*
ARRAYS - Arrays are a special data structure in Go that allow us to allocate contiguous blocks
of fixed memory. Arrays have some special features in Go related how they are declared and 
viewed as types.

DECLARING AND INITIALIZING VALUES
Declare an array of five strings initialized to its zero default with 
var strings [5]string - [5]string is the type name
*/

func main(){
	
	var strings [5]string
	// A string is an immutable, two word, data structure representing a pointer to a backing array of
	// bytes and the total number of bytes in the backing array. Since this array is set to its zero
 	// default value state, every element is also set to its zero value state. This means that each 
	// string has the first word set nil and the second word set to 0. [*nil][0]
	
	// STRING ASSIGNMENT - What happens when a string is assigned to another string? o sea a string literal
	// to a string [*nil][0]. Well when a string is assigned to another string, the two word value is copied
	// resulting in two different string values, both sharing the same backing array. This way the cost
	// of the copy is the same since its a pointer and a length connected to the same backing array.
	strings[0] = "Apple"
	
	// ITERATING OVER COLLECTIONS - Go provides two different semantics for iterating over collections as
	// I've already seen. We can do so using VALUE semantics or POINTER semantics. 
	
	// Value Semantics: meaning its iterating over its own copy
	for i,fruit:=range strings{
		println(i,fruit)
	}
	
	println()
	println("Strings:")
	// Pointer Semantics: meaning that even though its still passed by value, its iterating over the original 
	for i:=range strings{
		println(i,strings[i]) // this is why its pointer semantics. when we only use one arg its the index and
		// have to then use the index with the original collection from outside of the for/range
	}
	
	// When using value semantic iteration, two things happen. 
	// 		- First, the collection I'm iterating over is copied and we iterate over the copy.
	// 		  In the case of an array, the copy could be expensive since the entire array is copied.
	// 		  In the case of a slice, there is no real cost since only the internal slice value is copied
	// 		  and not the backing array. 
	// 		- Second I get a copy of each element being iterated on.
	
	// When using pointer semantic iteration, I iterate over the original collection and access each element
	// associated with the collection directly. Again why in this second vesion we use strings[i] and in the
	// first there's no need as everything we need is already copied into the program boundary/frame/stack
	
	// VALUE SEMANTIC ITERATION
	// Now looking at the following code
	println()
	println("Strings2:")
	strings2 := [5]string{"apple","orange","banana","grape","plum"}
	for i,fruit := range strings2{
		println(i,fruit)
	}
	
	// The strings2 variable is an array of 5 strings. The loop iterates over each string in the collection
	// and displays the index position 'i' and the string value 'fruit'. Since this is value semantic
	// iteration, the for range is iterating over its own shallow copy of the array and on each iteration
	// the fruit variable is a copy of each string (the two word data structure). So the copy is technically
	// still point to a backing array, but its still a copy of the 2 word data structure and not the original.
	
	// Notice how the fruit variable is passed to the print function (another program boundary) using value
	// semantics as well. The println function is getting its very own copy of the string value as well.
	// By the time the string is passed to the print function, there are 4 copies of the string value (array,
	// shallow copy, fruit variable and the print function's copy). All 4 copies are sharing the same backing
	// array of bytes. Making copies of the string value is important because it prevents the string value
	// from ever escaping to the heap. This eliminates non-productive allocation to the heap.
	
	// So what about POINTER SEMANTIC ITERATION?
	// Let's look at the following code:
	println()
	println("Strings3:")
	strings3 := [5]string{"apple","orange","banana","grape","plum"}
	for i := range strings3{
		println(i,strings3[i])
	}
	// Once again, the strings3 variable is a array of 5 strings. The loop iterates over each string in the
	// collection and displays the index position and the string value. Since this is pointer semantics though,
	// the for/range is iterating over the strings array directly and on each iteration, the string value for
	// each index position is accessed directly for the print call. 
	
}
