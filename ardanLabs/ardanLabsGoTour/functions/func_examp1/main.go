package main

/*
	Functions are at the core of the Go language and they provide a mechanism
	to group and organize our code to separate and distinct pieces of functionality.
	They can be used to be provide an API to the packages we write and are a core
	component to concurrency

	Let's first look at how functions return multiple values
*/

import (
	"encoding/json"
	"fmt"
)

// user is a struct type that declares user information
type user struct {
	ID   int
	Name string
}

func main() {

	// Retrieve the user profile.
	u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// display user profile
	fmt.Printf("%+v\n", *u)
}

// retrieveUser retrieves the user document for the specified user and returns
// a pointer to a user type value. snif snif returning a pointer, smells like something
func retrieveUser(name string) (*user, error) {

	// make a call to get the user in a json resposne
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of the user struct type
	var u user
	err = json.Unmarshal([]byte(r), &u)
	return &u, err
}

// Getuser simulates a web call that returns a json document for t he specified user.
func getUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}

// The point of this exaample is that our Go functions have the ability to return multiple 
// values, including errors as values.
