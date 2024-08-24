package main

import (
	"bufio"
	"fmt"
	_"log"
	"os"
	"time"
)

func main() {

	scn := bufio.NewScanner(os.Stdin)

	now := time.Now()

	for {
		fmt.Printf("> ")
		if !scn.Scan() {
			break
		}
		fmt.Println(now.Format(scn.Text()))
	}
}
