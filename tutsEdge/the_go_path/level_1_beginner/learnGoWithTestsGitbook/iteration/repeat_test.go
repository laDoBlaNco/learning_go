package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10) // note how we are using different vars for
	// in our tests for 'got' and 'want' and they still work.
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// writing benchmarks in Go is another first-class feature and its very
// similar to writing tests, but we use *testing.B
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

// testing.B gives us access to b.N. b.N runs and times how long it takes.
// the framework will determine the number of N times it runs .
// Then we run 'go test -bench=.

func ExampleRepeat() {
	repeat := Repeat("b", 15)
	fmt.Println(repeat)
	// Output: bbbbbbbbbbbbbbb
}
