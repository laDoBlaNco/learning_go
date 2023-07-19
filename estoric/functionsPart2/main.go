package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Another way is to create a type based on the func sig we want to use
type opFuncType func(int, int) int // now we can use this type in the map
// And it has the same effect as putting the func sig in the map

// we create a map which specifies the function signature of the values
// it will accept. If you try to use a function that has a different sig
// var opMap = map[string]func(int, int) int{ // putting the func sig directly in map
var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": multiply,
	"/": divide, // note the last one still needs a comma
	//"?": fmt.Println, // this gives me an error due to the Println sig
}

func main() {

	// here we use a slice of a slice of strings (basically multidimensional)
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expressions: ", expression)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op] // getting a value from map returns 2 values val/bool
		if !ok {
			fmt.Println("unsupported operation:", op)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	// Basic Anonymous func example.
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside the anonymous func")
		}(i) // similar to JS just adding the () after the def, also no name needed
	}

	// Another basic example of Anonymous funcs with the closures use case
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	// Passing a function in as a param
	people := []struct {
		Name string
		Age  int
	}{
		{"John", 35},
		{"Sam", 23},
		{"Alice", 53},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
	// works the same as JS sorts for array with callbacks
}

func add(i, j int) int      { return i + j }
func sub(i, j int) int      { return i - j }
func multiply(i, j int) int { return i * j }
func divide(i, j int) int   { return i / j }

func intSeq() func() int { // note intSeq returns a func that returns an int
	i := 0
	return func() int { // the func returned is bound to 'i' in this scope and thus
		i++ // will keep track of its state
		return i
	}
}

// Function signature is the parameters and return type. We need to keep this
// in mind to treat functions as first class values or citizens

// Note this time the go fmt let me keep the entire function on one line??? This
// is completely legal

// In the example aboe we see how we were able to store functions in a map
// as long as they func sig matched the map reqs. Then after implementing our
// logic we simply access the func that we need with the numbers provided.
// This is a great example of the modularity of first class functions in a lang.

// Another way to do the the mapping is creating our own type off of the func
// sig that we are requiring. Very interesting option. We can design types out
// of almost anything which is super powerfull. Its  best practice, if we are going
// to use something multiple times then we can assign it a name.

// In Go we can also anonymous funcs. A use cases for anonymous functions is closures
// This is mainly so we can reference vars bound to the function at a later time.

// The last  function case we are looking at is when you would pass in a function as
// a parameter. Why would you need this?
