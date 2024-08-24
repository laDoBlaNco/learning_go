package main

/*
After arrays and slices comes the 3rd data structure in Go. Maps.
*/

import (
	"fmt"
	"sort"
)

// sample program to show how to initialize a map, write to it, then read and delete from it

// user represents someone using the program
type user struct {
	name    string
	surname string
}

type player struct {
	name  string
	score int
}

// here users defines a slice of users
// type users2 []user

func main() {

	// Declare and 'make' a map that stores values of type user with a key of type string
	// NOTE the type signature [string]user
	// Also maps use value semantics. So its story its own copy of the data that it receives so its not a pointer
	// to some other data somewhere WHEN STORING IT. And it returns a copy. Copies in and copies out as we are
	// just moving data around.

	users := make(map[string]user)

	// Add key/value pairs to the map
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}
	fmt.Println(users)
	fmt.Println(len(users))

	// Read value at a specific key.
	mouse := users["Mouse"]
	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key
	users["Mouse"] = user{"Jerry", "Mouse"}
	mouse = users["Mouse"]
	fmt.Printf("%+v\n", mouse)

	// Delete the value at a specific key
	delete(users, "Roy")
	fmt.Println(users)
	fmt.Println(len(users))

	// it is safet to delete keys that are not actually there or already deleted
	delete(users, "Roy")
	fmt.Println("Goodbye")
	fmt.Println()

	// But what if the key doesn't exist?
	// Create map to track scores for players in a game
	scores := make(map[string]int)

	// Read the element at key "anna", it is absent so we get the zero default for this map's
	// VALUE TYPE. When we try to read off something that doesn't exists, it won't error out as we will
	// get a value. Remember that every var is started with a value.
	score := scores["anna"]
	fmt.Println("Score:", score)

	// If we need to check for the presence of a key first, then we can use
	// a 2 variable assignment, taking advantage of multiple assignment. The 2nd var is
	// a bool. NOTE: this is a very common Go idiom supported in various other situations
	score, ok := scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)

	// with this mechanism we can now leverage our zero default  values in other ways, such as simply
	// setting it to 1 with ++ which is a legal expression. so no need to first add Anna and then set it
	scores["anna"]++

	// Without this behavior we would have to code in a defensive way like this:
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"] // NOTE the use of = instead of := for this one
	fmt.Println("Score:", score, "Present:", ok)
	// this is a good example of how Go makes the zero value useful which is what most Go devs strive for
	// in there code.

	// Another NOTE is that the only value that can be used as a key is something that can be used as a
	// conditional operator.

	// Declare an make a map that uses a slice as the key
	// u := make(map[users2]int)
	// Results in: ./main.go:94:16: invalid map key type users2
	// Iterate over the map.
	// for key,value:=range u{
	// 	fmt.Println(key,value)
	// }

	fmt.Println()
	// we can also use map literal construction without the use of make
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},    // also note that we don't have to do user{"rob","roy"}
		"Ford":    {"Henry", "Ford"}, // since we know its users in this map, no need to be redundant
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"}, // note the trailing comma as I've gotten used to
	}

	// Iterate over the map printing each key and value with the value semantic form of for/range
	for key, value := range users2 {
		fmt.Println(key, value)
	}
	fmt.Println()
	// Also note that maps are random. Each time we run the code, it prints out in a random order. This is
	// by design so that we don't depend on Maps to hold a certain order in our algorithms. This protects us
	// from any changes in the future around maps not breaking our code if we aren't depending on a certain
	// ordering behavior.
	for key := range users2 { // pointer semantic version
		fmt.Println(key)
	}

	// Now let's try sorting with Maps. first we pull keys from the map and put into a new slice
	var keys []string // again, zero default so we use 'var'
	for key := range users2 {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically
	sort.Strings(keys)
	fmt.Println()
	// walk through the keys and pull each value from the map. so even though the map isn't in any order
	// we can pull the keys into a slice, sort the slice and then use that to pull out the map items
	// in a certain order.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
	fmt.Println()

	// Declare a map with initial values using a map literal
	players := map[string]player{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 21},
	}
	fmt.Println(players)

	// Trying to take the address of a map element will fail
	// anna:=&players["anna"] // This really is illegal cuz Go can't take the address of something that
	// really isn't stored anywhere yet. Because its not going to be placed in local memory directly here
	// there's no variable associated with it. 
	// NOTE: Remember that when working with maps everytime we store the map its going to be a copy
	// everytime we read the map its going to be a copy, so we can't just ask for the address of a value
	// that is internal to the  map as that would be a violation of integrity and that's why we can't 
	// use our pointer semantics here. So as seen below Go will let us get our own copy out of the map
	// work with it and then store it back into the map.
	// anna.score++
	// ./main.go:154:9: invalid operation: cannot take address of players["anna"] (map index expression of type player)

	// Remembering that maps are reference types, we get this error because get an address of this value until
	// we save it to a local variable first. Once we do that, then we can do what we wanted to, without using
	// the pointer semantics
	// Instead we take the element, modify it, and put it back
	player := players["anna"]
	player.score++
	players["anna"] = player
	fmt.Println(players)
	fmt.Println() 
	
	// REFERENCE TYPES
	// Remember with ref types we have the duality of moving it around with value semantics but reading
	// and writing with pointer semantics.
	scores = map[string]int{
		"anna":21,
		"jacob":12,
	}
	
	// Pass the map to a function to perform some mutation
	double(scores,"anna") 
	
	// See the change is visible in our original map
	fmt.Println("Score:",scores["anna"]) 
	
	

}

// double finds the score for a specific player and multiplies it by 2
func double(scores map[string]int,player string){
	scores[player]*= 2 
}
// Remember that we are working with Reference types so in reality when we get that COPY of our
// map with our value semantics, its really just a copy of a pointer. So we aren't copying the 
// full map here. That's another reason why we don't really need to ever do the double indirection
// of creating a pointer to a map since the map is already just a pointer.
