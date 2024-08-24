package main

import (
	"errors"
	"fmt"
	"net/http"
)

// in the  last video we talked about defer and how that helps us with panics and recover, but we
// can also use them to ensure connections are closed after we are done with them.

// Here we are going to understand what they really are and how they work. Defers are like helpful
// todos to help us clean up our code in the end.

func ReturnValues() (x int) {
	defer func() {
		x = 10 // this is actually what'll be returned as its set at the end of the func run
	}()
	x = 5
	return
}

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

func doStuff(e *Engineer, url string) {

	// Even though Go doesn't call the deferred function until the end, it does evaluate any
	// vars or args immediately. That means that our original name will be placed in our method
	// even though we updated it prior to our deferred method actually being run. This is interesting
	// stuff. ;)
	name := "Kelen D Whiteside"

	// we'll see this pattern of deferring named methods a lot with things such as
	// http requests, so lets add one in this same function
	defer e.UpdateName(name)
	fmt.Println("doing other exciting stuff!")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// we can then process the rest of the body and not worry about having to close this
	// at the end of the function.
	fmt.Println(resp.Status)

	// Now there is a gotcha with this which we'll see here as well.
	name = "Kevin Anthony Whiteside"
}

func MyAwesomeFunction() (err error) {
	defer func() {
		err = errors.New("I am an error")
	}()

	return nil
}

func main() {

	// defer functions are always executed at the end of the program, no matter when they are
	// run. Meaning all of main must complete and then the deferred stack will start.

	defer func() {
		fmt.Println("I will be executed last.")
	}()

	// its a stack that will run in LIFO order when there are multiple defer funcs
	defer func() {
		fmt.Println("Last in - but first out")
	}()

	fmt.Println("Defer in Go!")
	println()
	err := MyAwesomeFunction()
	if err != nil {
		fmt.Println(err.Error())
	}
	println()
	fmt.Println(ReturnValues())
	println()

	lado := &Engineer{
		Name: "Kelen",
	}

	fmt.Printf("%+v\n", lado)
	doStuff(lado, "https://google.com")
	fmt.Printf("%+v\n", lado)

}
