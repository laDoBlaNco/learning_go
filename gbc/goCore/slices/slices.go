package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// slices are like arrays but they are more manageable. They are really
	// pointers or refs to underlying arrays.

	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice size:", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, v := range sl1 {
		pl(v)
	}

	sl2 := []int{12, 25, 1976}
	pl(sl2)

	arr1 := [5]int{1, 2, 3, 4, 5}
	sl3 := arr1[:2]
	sl4 := arr1[:3]
	sl5 := arr1[2:]

	pl("1st 2:", sl3)
	pl("1st 3:", sl4)
	pl("last 3:", sl5)

	// Remember that since its a pointer, if you change the slice, the underlying
	// array will be changed as well and vice versa
	arr1[0] = 10
	pl("sl3:", sl3)

	sl3 = append(sl3, 12)
	pl("sl3:", sl3)
	pl("arr1:", arr1)

	sl6 := make([]string, 6)
	pl("sl6:", sl6)
}
