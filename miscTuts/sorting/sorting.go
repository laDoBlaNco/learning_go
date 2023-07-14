package main

import (
	"fmt"
	"sort"
)

// basic sorting of strings integers and floats in Go

func main() {

	// sorting a slice of strings
	strs := []string{"quick", "brown", "fox", "jumps"}
	sort.Strings(strs)
	fmt.Println("Sorted strings: ", strs)

	// sorting a slice of ints
	ints := []int{56, 19, 78, 67, 14, 25}
	sort.Ints(ints)
	fmt.Println("Sorted integers: ", ints)

	// Sorting a slice of Floats
	floats := []float64{176.8, 19.5, 20.8, 57.4}
	sort.Float64s(floats)
	fmt.Println("Sorted floats: ", floats)

	fmt.Println()
	fmt.Println("Using Reverse & ...Slice interfaces and funcs:")

	// Now let's do reverse order which is just using some func composition
	// so while sort.Strings, .Float64s, and .Ints are all sorting functions
	// Reverse and StringSlice return interfaces that are used by other sorting funcs
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("Sorted strings in reverse order: ", strs)

	// reversing our sorted Integers
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Sorted integers in reverse order: ", ints)

	// Sorting a slice of floats
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Println("Sorted floats in reverse order: ", floats)

	fmt.Println()
	fmt.Println("Now we use a comparator function, which is like a JS callback")

	// When we want to sort something other than in its natural order we can tell
	// Go how to do it on the fly with a comparator function. We can first use
	// sort.Slice() and SliceStable()  which accept a 'less' func as an arg.
	// The difference between Slice and SliceStable is that equal elements may
	// switch other. For a stable sort use SliceStable
	strs = []string{"United States", "India", "France", "United Kingdom", "Spain"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("Sorted strings by length: ", strs)

	// Stable sort
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("[Stable] Sorted strings by length: ", strs)

	fmt.Println()
	fmt.Println("Sorting a slice of strings in the reverse order of length")

	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[j]) < len(strs[i])
	})
	fmt.Println("[Stable] Sorted strings by reverse order of length: ", strs)

	// Sorting a slice of structs by a field
	fmt.Println()
	fmt.Println("Sorting a slice of structs using a comparator function")
	users := []User{
		{
			Name: "Rajeev",
			Age:  28,
		}, {
			Name: "Monica",
			Age:  31,
		}, {
			Name: "John",
			Age:  56,
		}, {
			Name: "Amanda",
			Age:  16,
		}, {
			Name: "Steven",
			Age:  28,
		},
	}

	fmt.Println("First by age:")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("Sorted users by age: ", users)

	// Stable sort
	sort.SliceStable(users, func(i, j int) bool {
		return users[j].Age < users[i].Age
	})
	fmt.Println("[Stable] Sorted users by age: ", users)

	// Finally let's look at the most generic way of sorting a collection of
	// primiitives or user-defined structs in Go. To enable custom sorting of a
	// collection of any type, you need to define a corresponding type that
	// implements the generic Interface provide by the sort package. Once I
	// implement this interface I can sort any collection that implements the
	// sort.Interface interface
	sort.Sort(UsersByAge(users))
	fmt.Println("Sorted users by age: ", users)
	sort.Stable(UsersByAge(users))
	fmt.Println("[Stable] Sorted users by age: ", users)

	fmt.Println()
	fmt.Println("Now let's try by name:")
	sort.Sort(UsersByCulo(users)) 
	fmt.Println("Sorted users by name: ", users)

}

type User struct {
	Name string
	Age  int
}

// define a collection type that implements sort.Interface
type UsersByAge []User

func (u UsersByAge) Len() int {
	return len(u)
}
func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

// A test to see if I can re-implement to order by name
type UsersByCulo []User

func (u UsersByCulo) Len() int {
	return len(u)
}
func (u UsersByCulo) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByCulo) Less(i, j int) bool {
	return u[i].Name < u[j].Name
}
// By changing the name I proved to my self that the name has nothing to do with it
// Its the Less function that identifies which field we are going to be working with.
// And the swap function as to how we will sort (normal or reverse), etc.
// Also in the end what we are doing is giving sort.Sort a sort.Interface value to
// work with. 
