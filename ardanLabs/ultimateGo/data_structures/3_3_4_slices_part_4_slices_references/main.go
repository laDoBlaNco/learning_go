package main

import "fmt"

/*
	Let's start looking at side effects and mutation when we are working with slices. This way
	we'll be able to start to identify things when reading more and more Go code.
	
	For example losing data when we make copies and our slices grow. When pointers are updated we
	could lose data if we don't update our pointers. We could have data flowing to an old backing 
	array.
	
	We must remember that when we are working with our reference types we might be using our 
	value semantics to  move things around but we are using our pointer semantics to read 
	and write and we need to remember that. 
	
	Anytime we see 'append' outside of a func and it could replace a backing array, we need to
	stop our code review and double check it. That could be an issue or bugs waiting to happen.
	
	Value semantics are safe, but when we are using pointer semantics, stranges things could happen.

*/

type user struct {
	likes int
}

func main() {

	// declare a slice of 3 users
	users := make([]user, 3)

	// Share the user at index 1
	shareUser := &users[1]

	// Add a like for the user that was shared
	shareUser.likes++

	// Display the number of likes for all users
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}
	
	// add a new user
	users = append(users,user{})
	
	// add another like for the user that was shared
	shareUser.likes++
	
	// display the number of likes for all users
	fmt.Println("**********************************")
	for i:= range users{
		fmt.Printf("User: %d Likes: %d\n",i,users[i].likes) 
	}
}

func inspectSlice(slice []string) { // the arg its asks for a a slice VALUE not a pointer to a slice
	fmt.Printf("Length[%d]  Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { // and here we use the VALUE semantic version of the for/range
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
	// How do we know the form of the for/range to use:
	// What is this a collection of? Slice of strings.
	// What is a string? A built-in type
	// What do we use for built-in types? Value semantics
}
