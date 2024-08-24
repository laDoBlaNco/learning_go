package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	var intArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intArgs = append(intArgs, val)
	}
	max := 0
	for _, val := range intArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value:", max)
}
