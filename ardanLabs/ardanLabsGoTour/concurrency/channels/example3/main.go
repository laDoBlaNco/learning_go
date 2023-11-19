package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

// waitForResult: In this pattern, the parent goroutine waits for the child gr
// to finish some work to signal the result.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d:=<-ch
	fmt.Println("parent : recv'd signal:",d) 
}
