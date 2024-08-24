package main

import(
	"fmt"
	"time"
)

func main(){

	fmt.Println("Runs on save...") 
	fmt.Println(time.Now().UTC().Second()) 
}


