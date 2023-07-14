package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
Needles and pins
Needles and pins
Sew me a sail
To catch me the wind
`

	// fmt.Println(text)
	wc := make(map[string]int)

	for _, w := range strings.Fields(text) {
		wc[strings.ToLower(w)]++
	}
	fmt.Println(wc)
}
