package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArr(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64
	var numSize float64 = float64(len(nums))
	for _, v := range nums {
		sum += v
	}
	return sum / numSize
}

func main() {

	f4 := 10
	pl("f4:", f4)
	pl("f4 Address:", &f4)

	var f4Ptr *int = &f4
	pl("f4 Address:", f4Ptr)
	pl("f4 Value:", *f4Ptr)

	*f4Ptr = 11
	pl("f4 Value:", *f4Ptr)

	pl()
	pl("f4 before function:", f4)
	changeVal2(&f4)
	pl("f4 before function:", f4)

	pl()
	pArr := [4]int{1, 2, 3, 4}
	dblArr(&pArr)
	pl(pArr)
	
	pl()
	mySlice:=[]float64{11,13,17} 
	fmt.Printf("Average is %.3f\n",getAverage(mySlice...)) 
}

// The keys to working with pointers:
// 1. Knowing when you need pointers and when you don't. If you want to change
//    the original item directly, use pointers.
// 2. Keep track of the use of * and & and when they are needed. *var is getting
//    the value of a varPtr. &var is getting the address of your value var or in
//    words creating a pointer. *T when * is when a type you are saying that your
//    var is actually a pointer type.
//
//    Now when using functions that ask for pointers 'changeVal2(arr *int)' requiring
//    a pointer type arg, you must give it your arg with &, in otherwords give it
//    a pointer (address). And to reset the value of a pointer use *ptr to get to the
//    value at that address "dereferncing"

// 3. Note that when I work with Arrays above I need to use the pointer operators
//    creating the pointers, etc. But when working with slices I don't. This is 
//    because slices are already pointers or references to underlying arrays??? I 
//    assume this is the case, I'll investigate myself, but the example above has 
//    no reference to pointers. Its kind of a stupid example for this lesson.
