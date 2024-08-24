package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// maps are collections with k,v pairs. The keys can be different types
	// than the values and needs to be a type that's comparable
	// var myMap map[keyType]valueType
	var heroes map[string]string     // this is a NIL map until its made
	heroes = make(map[string]string) // this completes creation of new empty map

	// we can also do it in one step:
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"
	pl()
	pl(heroes)

	// we can also do map literals
	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	pl()
	pl(superPets)
	pl()

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	fmt.Println("Calling something that doesn't exist - Chip:", superPets[3]) // this prints "" since  it doesn't exits
	// but we can check for that as well with comma ok
	_, ok := superPets[3]
	pl("Using 'comma ok' - Is there a third pet:", ok)

	pl()
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	delete(heroes, "The Flash")
	pl("After using 'delete' on Flash:", heroes)

}
