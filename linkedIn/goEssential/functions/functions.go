package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {

	val := add(1, 2)
	fmt.Println(val)

	fmt.Println("___________________________________")
	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)

	fmt.Println("___________________________________")
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	fmt.Println("___________________________________")
	val = 10
	double(val)
	fmt.Println(val)

	fmt.Println("___________________________________")
	doublePtr(&val)
	fmt.Println(val)

	fmt.Println("___________________________________")
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s2)
	}

	fmt.Println("___________________________________")
	worker()

	fmt.Println("___________________________________")
	//CHALLENGE http contentType
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

}

// simple func
func add(a int, b int) int {
	return a + b
}

// return more than one value
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

// doubles the val at a certain location in a slice, Note its not returning anything
func doubleAt(values []int, i int) {
	values[i] *= 2
}

// parameters pass args as values (copy). slices and maps work as pointers or refs
func double(n int) {
	n *= 2
}

// So let's use pointers
func doublePtr(n *int) {
	*n *= 2
}

// go functions returning more than one result is the basis of error handling in go
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	} // customary to return return the 0.0 value when there is an error
	return math.Sqrt(n), nil
}

// WE DON'T USE PANICS IN GO

// Working with resources in functions. The idiom of aquiring a resource, checking
// for an error and then deferring the release of the resource is very common in Go
func acquire(name string) (string, error) {
	return name, nil
}
func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r1)

	r2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r2)

	fmt.Println("worker") // this prints first then we see the defer call. The defers
	// are called in reversed order because they go onto a deferred stack (FILO)

}

// CHALLENGE:
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ERROR:%s", err)
	}
	defer resp.Body.Close() // make sure you release the resource

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // Return errir if Content-type header not found
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil
}
