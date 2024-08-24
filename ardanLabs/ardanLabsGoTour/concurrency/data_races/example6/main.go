package main

import (
	"fmt"
	"os"
	"sync"
)

// This program is to show a more complicated race condition using an interface value.
// This produces a read to an interface value after a partial write.

// Speaker looks for the speaking behavior
type Speaker interface {
	Speak() bool
}

// Ben is a person who can speak
type Ben struct {
	name string
}

// Lrt's create the method for Ben to say hello. It returns false if the method is
// called through the interface value after a partial write
func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
	}
	return false
}

type Jerry struct {
	name string
}

// Speak allows Jerry to say hello. It returns false if the method is called through
// the interface value after a partial write
func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
		return false
	}
	return true
}

func main() {

	// let's create values of type Ben and Jerry
	ben := Ben{"Ben"} 
	jerry := Jerry{"Jerry"} 

	// now lets assign the pointer to the Ben value to the interface value
	person := Speaker(&ben)

	// have a gr constantly assign pointer of the Ben value to the interface and then speak
	go func() {
		for {
			person = &ben
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// Let's have another gr constantly assigning the pointer of the Jerry value to the
	// interface and then Speak.
	go func() {
		for {
			person = &jerry
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// just hold main from returning. The data race will cause the program to exit
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
