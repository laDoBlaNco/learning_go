package main

import (
	"fmt"
)

// ...for quick debugging
var pl = fmt.Println

// specifying data types to be used at later times. Mainly used to create functions
// that can take multiple data types

type MyConstraint interface {
	int | float64 // we could also have use 'any' here but that's not smart
	// another general is 'comparable' which is any type that can be used with ==
}

func getSumGen[T MyConstraint](x, y T) T {
	return x + y
}

func main() {
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
	// pl("5.6 + 4.7 =", getSumGen("5.6", "4.7")) // receives error: string doesn't
	// implement MyConstraint interface (string missing in int | float64 )

}
