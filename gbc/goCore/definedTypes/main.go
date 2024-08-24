package main

import (
	"fmt"
)

// ...for quick debugging
var pl = fmt.Println

// we use defined types to enhance the quality of other data types. So its somewhat
// like overloading in other langs. We take a builtin datatype (or I guess even
// a custom one now that I think about it) and we base our custom type on top of it
// So you are basically adding your custom methods, fields, etc on top of what
// comes in the original type.

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}
func tBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {

	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f ML\n", ml2)

	pl("2 tsp + 4 tsp =", Tsp(2)+Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	pl()
	fmt.Printf("3 tsp = %.2f ml\n", tspToML(3))
	fmt.Printf("3 tbs = %.2f ml\n", tBToML(3))

	pl()
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
}
