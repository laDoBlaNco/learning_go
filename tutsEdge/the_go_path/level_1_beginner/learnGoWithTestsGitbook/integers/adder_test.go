package integers

import (
	"fmt"
	"testing"
)

// using package integers instead of 'main'

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
		// here we are using '%d' instead of %q so we can get an integer
		// instead of a string.
	}
}

// we can also put Example functions in our test files and they will compile
// as part of the package's test suite
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
	// NOTE: with the comment 'Output: 6' the Example func won't be run
}
