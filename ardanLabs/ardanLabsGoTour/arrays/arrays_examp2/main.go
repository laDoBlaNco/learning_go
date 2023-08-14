package main

// Its interesting to see what the compiler provides as an error when assigning arrays of the same types
// that are of different lengths

func main(){
	
	var five [5]int // a zero defaulted array of ints of length 5
	four := [4]int{10,20,30,40}
	
	five = four 
	// ./prog.go:11:9: cannot use four (variable of type [4]int) as [5]int value in assignment
	// O sea, [4]int and [5]int are two different types
	
	// Here  I declare an array of 4 and 5 integers initialized to its zero value state. Then try to 
	// assign them to each other and the compiler screams, "You can't use that type[4]int as a type [5]int
	// value
	// Its important to be clear about what the compiler is saying here. Its saying that an array of 4 ints
	// and an array of 5 ints represent data of different types. The size of the array is part of its type 
	// information. In Go, the size of an array has to be known at compile time.
	
}
