package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	fmt.Println("Hello", stuff.Name)
	intSlice := []int{2, 3, 5, 7, 11}
	strSlice := stuff.IntSliceToStrSlice(intSlice)
	fmt.Println(strSlice)
	pl(reflect.TypeOf(strSlice))
}
