package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// if you initialize an array or slice with no values, you get the zero defaults
	// for the type of array you are working with.

	// Array size can't be changed. The size is part of the arrays type
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5} // composite literal
	pl("Index 0:", arr2[0])
	pl("Array Length:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	str1 := "abcde"
	rSlice := []rune(str1)
	for _, v := range rSlice {
		fmt.Printf("Rune Array: %d\n", v)
	}
	bSlice := []byte{'a', 'b', 'c'}
	bStr := string(bSlice)
	pl("I'm a string:", bStr)

}
