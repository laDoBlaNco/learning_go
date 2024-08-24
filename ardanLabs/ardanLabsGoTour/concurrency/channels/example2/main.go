package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this sample demonstrates the fan out channel pattern

func main() {
	fanOut()
}

// fanOut: In this pattern, the parent goroutine creates 2000 child goroutines
// and waits for them to signal their results. The work is "fanned out"
// and then collected.
func fanOut() {
	children := 2000
	ch := make(chan string, children) // buffer is same size as number of grs

	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------")
}

// NOTE a couple key items.
// 1. we created a channel buffered to the same size as the known number of grs
// We used a for/loop to ensure we collected all of the grs results before
// allowing the program to close on us. 
