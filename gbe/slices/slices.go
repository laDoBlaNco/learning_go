package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Slices are an important data type in Go, giving a more versatile and powerful
// interface to sequences than arrays.
func main() {
	// Unlike arrays, slices are  typed only by the elements they contain (not
	// by the NUMBER of elements). To create an empty slice with non-zero lengths
	// you can use the built-in 'make'. Here we make a slice of strings of length
	// 3 (initially zero values)
	s := make([]string, 3)
	p("empty slice of 3 strs:", s)
	fmt.Printf("Same with quotes: %q\n", s)
	p()

	// we can set and get just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	p("set:", s)
	p("get:", s[2])
	p()

	// len also returns the length of a slice just like an array
	p("length:", len(s))
	p()

	// In addition to these basic operations, slices support several more than
	// make them richer than arrays. One is the builtin 'append', which returns
	// a slice containing one or more new values. Note that we need to accept
	// a return value from append as we may get a new slice value.
	s = append(s, "d")
	p("appended:", s)
	p("length and cap:", len(s), cap(s))
	p()

	// so if s is now len 4 and cap 6, I can technically append 2 more without
	// getting a new slice returned.-- ok so after the test, I still need to use
	// the return value to get the adjusted slice.
	s = append(s, "e", "f")
	p("appended:", s)
	p("length and cap:", len(s), cap(s))
	p()

	// Slices can also be copy'd. here we create an empty slice c of the same
	// len as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	p("copy'd:", c)
	p()

	// Slices support a "slice" operator with the syntax 'slice[low:high]'. For
	// example, this gets a slice of the elements from s[2],s[3], and s[4].
	l := s[2:5]
	p("slice1:", l)
	p()

	// this slices up to (but excluding) s[5]
	l = s[:5]
	p("slice2:", l)
	p()

	// and this slices up from (and including) s[2]
	l = s[2:]
	p("slice3:", l)
	p()

	// we can declare and initialize a var for slice in a single line as we do
	// with arrays (slice literals)
	t := []string{"g", "h", "i"}
	p("slice literal:", t)
	p()

	// and slices can also be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional
	// arrays and this is because the length of an array is part of its type.
	// This means that an array of type [5]string, must all be of len 5.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1 // so it'll grow on each iteration
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	p("2D:", twoD)

}

// I received a very interesting panic on the last example because I forgot to
// include line '90'. This reminds me that even though I can refer to it in the
// program and it'll compile, the second I try to assign to a slice that isn't
// made, go will panic.

// Also note that while slices are different types than arrays, they are
// rendered teh same by fmt.Println

