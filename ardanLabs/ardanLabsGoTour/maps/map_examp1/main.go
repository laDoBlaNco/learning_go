package main

import "fmt"

/*
	A map is a data structure that provides support for storing and accessing data
	based on a key. It uses a hash map and bucket system that maintains a contiguous
	block of memory as well underneath. In fact it itself is also implemented with
	an array. Meaning we get the same mechanical sympathies that we do with slices.

	We can declare and construct maps in several ways:
*/

type user struct {
	name     string
	username string
}

// Construct a map set to its zero value, that can store user values based on a key of type string
// Trying to use this map will result in a runtime error (panic).
// var users map[string]user

func main() {
	// Construct a map initialized using make, that can store user values based on a key of type string
	users2 := make(map[string]user)
	fmt.Println(users2)

	// Construct a map initialized using empty literal construction, that can store user values based
	// on a key of type string
	users3 := map[string]user{} // we tend to not want to use empty literals for constructions
	fmt.Println(users3)

	//  a map set to its zero value is not usable and will result in our program panicking .
	// The use of the built-in function 'make' and literal construction constructs a map ready for use.
	users2["Roy"] = user{"Rob", "Roy"}
	users2["Ford"] = user{"Henry", "Ford"}
	users2["Mouse"] = user{"Mickey", "Mouse"}
	users2["Jackson"] = user{"Michael", "Jackson"}

	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// Once data is stored inside a map, to extract any data a key look is required.
	// what does that look like?
	user1, exists1 := users2["Bill"]
	user2, exists2 := users2["Ford"]
	fmt.Println("Bill:", exists1, user1)
	fmt.Println("Ford:", exists2, user2)

	/*
		To perform a key lookup, square brackets are used with the map variable. Two values can be returned
		from a map lookup, the value and a boolean that reps if the value was found or not. If you don't need
		to know this, you can leave the 'exists' var out.

		When a key is not found in a map, the operation returns a value of the map type set to its zero
		value state. You can see this with 'Bill' key lookup. We can't use the zero value to determine
		if a key exists in the map since the zero value might be a valid value that was actually stored
		for the key.
	*/

	// we delete using the built-in 'delete' func
	delete(users2, "Roy")
	fmt.Println(users2)
	fmt.Println()
}

/*
	If the built-in 'make' function is used to construct a map, then the assignment operator can be
	used to add and update values in the map. The order of how keys/values are returned when ranging
	over a map is undefined by the spec  and up to the compiler to implement. This ramdomness is designed
	to make sure that we don't rely on a specific order for our design.

	The current algo since 1.16 will return the results in a random order once the number of values
	reaches a certain limit. Once again, this is a compiler implementation that is allowed to change.
	We can't depend on it.

*/
