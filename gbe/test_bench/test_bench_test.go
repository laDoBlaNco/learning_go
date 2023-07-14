// Go by Example: Testing and Benchmarking
// Unit testing is an important part of writing principled Go programs. The testing
// package provides the tools we need to write unit tests and the go test command
// runs tests.

// For the sake of demonstration, this code is in package main, but it could be any package
// Testing code typically lives in the same package as the code it tests

package main

import (
	"fmt"
	"testing"
)

// we'll be testing this simple implemenation of an integer minimum. Typically
// the code we're testing would be in a source file named 'test_bench.go' and the
// test file for it would be in test_bench_test.go (the '_test' needs to be inserted before
// the '.go' ext)
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name beginning with 'Test'
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue executing the test
		// t.Fatal* will report test failures and stop the test immediately
		t.Errorf("IntMin(2,-2) = %d; want -2", ans)
	}
}

// Writing test can be repetitive, so it's idiomatic to use a table-driven style, where test
// inputs and expected outputs are listed in a table and a single loop walks
// over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	// t.Run enables running subtests, one for each table entry. These are shown
	// separately when executing go test -v
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmarks are apparently just another type of test :-)
// They typically go in _test.go files as well and are named beginning with Benchmark.
// The testing runner executes each benchmark function several times, increasing b.N on each
// run until it collects a precise measurement.
func BenchmarkIntMin(b *testing.B) { // note the diff. '(b *testing.B)' vs '(t *testing.T)'
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

// Run all tests in the current project in verbos mode = 'go test -v'

// Run all benchmarks in the current project. - 'go test -bench=.c'

// All tests are run prior to benchmarks. The bench
// flag filters benchmark funcs with a regexp.
