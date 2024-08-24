// https://tutorialedge.net/golang/getting-starting-with-go-generics/
// Getting Started With Go Generics - Tutorial
package main

import (
	"fmt"
)

func oldNonGenericFunc(myAge int64) {
	fmt.Println(myAge)
}
func newGenericFunc[age int64 | float64](myAge age) { // note that 'age'
	// is used as the generic type var. It doesn't have to be T all the time
	// which is what was throwing me off previously.
	fmt.Println(myAge)
} // and it works as expected taking in both ints and floats. We are
// basically saying in the type param []s that 'age' represents both
// types int64 and float64, making 'age' a generic type.

// we could also do the same using 'any' as in type with no restrictions
func newGenericFunc2[age any](myAge age) {
	fmt.Println(myAge)
}

// Now let's try doing some casting and adding 1 to age
/*
func newGenericFunc3[age any](myAge age){
	val := int(myAge) + 1
	fmt.Println(val)
}// this doesn't work cuz using 'any' won't allow us to cast anything.
// we need to be more specific about the types allowed and then Go will
// say whether it can cast or not
*/

func newGenericFunc4[age int64 | float64](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

// Now in most cases Go will infer the type from the parameter for our
// generic functions. But in case it can't and we need to more deliberate
// can use our same type parameters in the function call, stating the
// type we are passing in brackets after the func name.
// newGenericFunc[int64](testAge)

// So what about if we want to create constraints outside of the type
// parameters. Constraints in generics are just interfaces, but now in
// addition to being able to create a method list, we can create a type
// list. Putting all of these conditions or constraints into one
// interface allowing us to simply use that in our generic function
type Age interface {
	int64 | int32 | float32 | float64
}

func newGenericFunc5[age Age](myAge age) { // note our use of 'Age' as the
	// constraint for the generic type age
	val := int(myAge) + 1
	fmt.Println(val)
}

// We fan also have more complex constraints. For example let's look
// at getSalary below where we take anything that satisfies a given
// type constraint. We define an interface and use it in our generic
// function
type Employee interface { // we have an Employee interface
	PrintSalary()
}

func getSalary[E Employee](e E) { // our type constraint is that the
	// type of our arg fits the Employee interface, and since Employee has
	// a method, the arg used must also have method PrintSalary()
	e.PrintSalary() // the reason it must have this method, same as any
	// other interface, but now with generic types.
}

// so now we can create other types and as long as they fit the description
// (interface/constraint) of being an Employee, they will fit our
// type constraint.
type Engineer struct { // so we create a struct type Engineer with the
	// right fields and methods and now he's an Employee
	Salary int32
}

func (e Engineer) PrintSalary() {
	fmt.Println(e.Salary)
}

type Manager struct { // also a Manager (Employee)
	Salary int64
}

func (m Manager) PrintSalary() {
	fmt.Println(m.Salary)
}

// we could've done the same just using interfaces as well because
// using interfaces is a type of "generic programming", but there are
// differences between the approaches and benefits of one over the other

// THE BENEFITS OF GENERICS
// So far we've covered  the basics with generics and what we'll encounter in the
// wild. Let's create an example of something can can be beneficial in our own Go
// apps.

// Generics with a Standard BubbleSort:
func BubbleSort(input []int) []int{ // taking an array of int and return the same
	n := len(input)
	swapped := true
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements i n our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i] > input[i+1] {
				// Log that we ar swapping values for posterity
				fmt.Println("Swapping")
				// swap values using Go's tuple assignment (idiomatic)
				input[i], input[i+1] = input[i+1], input[i]
				// set swapped to true - this is important if the
				// loop ends and swapped is still equal to false
				//, our 'rithm will assume that the list is fully
				// sorted.
				swapped = true
			}
		}
	}
	return input
} // so here we defined our rithm for int. So if we try to run it with in32? See
// below in main

// Now let's try the generic version
type Number interface{
	int16|int32|int64|float32|float64
}

func BubbleSortGeneric[N Number](input []N)[]N{
	n:=len(input)
	swapped:=true
	for swapped{
		swapped=false
		for i:=0;i<n-1;i++{
			if input[i]>input[i+1]{
				input[i],input[i+1]=input[i+1],input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5
	var testString string = "ladoblanco"

	newGenericFunc(testAge)
	newGenericFunc(testAge2)

	newGenericFunc2(testAge)    // works the same with any
	newGenericFunc2(testAge2)   // works the same with any
	newGenericFunc2(testString) // works the same with any

	// newGenericFunc3(testAge) // cannot convert myage (variable of type
	// newGenericFunc3(testAge2) // age constrained by 'any') to type int

	newGenericFunc4(testAge)
	newGenericFunc4(testAge2)

	newGenericFunc5(testAge)
	newGenericFunc5(testAge2)

	// Now for our complex Type Constraint
	fmt.Print("\n")
	fmt.Println("Complex Constraints:")
	engineer := Engineer{Salary: 10000}
	manager := Manager{Salary: 100000}

	getSalary(engineer)
	getSalary(manager) // and our generic function works with both because
	// both fit the Employee interface. THIS IS WHAT I ENCOUNTERED IN
	// BOOT.DEV

	fmt.Println()
	fmt.Println("BubbleSort Example:")

	list := []int32{4, 3, 1, 5, 69}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2, 6.9}

	// sorted := BubbleSort(list) // we get an error for tying to use int32 with int
	// fmt.Println(sorted)
	sorted:=BubbleSortGeneric(list)
	fmt.Println(sorted)
	
	sortedFloats:=BubbleSortGeneric(list2)
	fmt.Println(sortedFloats) 
	
	fmt.Println()
	fmt.Println("Are slices inline?")
	fmt.Println(list,list2) 
	fmt.Println("Yes...Yes they are. There was no need to build a new list")
	// So as I thought, no need to return the slice as they are changed inline. 
	// Remember that slices, maps, etc are passed by reference/pointer in Go.
	

}
