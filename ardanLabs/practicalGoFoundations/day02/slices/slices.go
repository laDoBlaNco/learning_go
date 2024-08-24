package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var s []int                // s is a slice of int
	fmt.Println("len", len(s)) // len is 'nil safe'
	// you can't compare slices, I didn't know that. This was a design decision
	// for efficiency. You can only compare slices to nil
	if s == nil {
		fmt.Println("nil slice")
	}

	// You can also create slice literals. Remember the only difference form an arr
	// creation is that we don't put a number in the bracket [5]int vs []int
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 =%#v\n", s2)
	s3 := s2[1:4] // slicing operation, half-open range, meaning the upper-bound is
	// non inclusive
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100])  // this gives us a panic out of range

	s3 = append(s3, 100)
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2) // remember that slices reference an\
	// underlying array.

	// A slice is a struct with 3 fields. Anytime we use a slice we are copying this
	// struct which uses a pointer to memory. The underlying array with the pointer
	// pointing to its beginning. When we slice the same array, or slice a slice
	// we add another struct but with the pointer pointing to another start point.
	// with the len being the size of the array and the cap being from where that
	// pointer is set up until the end of the underlying array.

	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	// really all capacity does is tell append if it needs to get more memory

	// var s4 []int                 // default 0 values so it'll need reallocation
	s4 := make([]int, 0, 1_000)  // Single allocation O(1) because we know its big enough
	for i := 0; i < 1_000; i++ { // you can use _ in numbers in go for the normal commas to help visually
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))

	// CHALLENGE: Implement concat function
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) //[A B C D E]

	// median function using sort pack
	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 4, 3}
	fmt.Println(median(nil))
	fmt.Println(vs)

	fmt.Println(reflect.TypeOf(2))

	// although we are passing around a copy of a copy of a copy of the slices
	// they are point pointers to the same Array in memory, thus we are impacting
	// the original array. So what we can do is add some code to copy in order
} // before the sorting of the slice

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// Copy in order so as not to change underlying array
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	//if len(nums)%2 == 1 {
	if len(nums)&1 == 1 { // here we use &1 (bitwise AND operator  to find odds)
		return nums[i], nil
	}

	const n = 2
	v := (nums[i-1] + nums[i]) / n

	return v, nil
}

// Something to note. Go's concept of constants with numbers. Go is very strict with
// mixing types. If we put n := 2 and use n above then we get a compile error.
// but if we use the digit 2 (still a whole number) it works and if we put a const
// on the n that works as well. Their type is defined when they are used. If we use
// a const n with floats or with ints it'll still work. The whole number itself
// is a constant and it'll work in both scenarios as well.

// Also if we trying to get the median of nil we get a panic. In go we don't like
// panics. So we change the func to return a 2nd value, an error. With is more
// idiomatic in Go. If there's a possibility of an error, address it by returning
// a zero default and the error

// My way is more robust than the instructor he simply went straight for creating a
// new slice while i'm leaving it open to the possibility that my first arg isn't
// yet full. Thus not increasing the space complexity when there's no need.
func concat(s1, s2 []string) []string {
	i := len(s1)
	capNeed := len(s1) + len(s2)
	if capNeed > cap(s1) {
		temp := make([]string, capNeed)
		copy(temp, s1)
		s1 = temp
	}
	copy(s1[i:], s2) // s1[i:] is a slice of a slice, so copying into it actually
	// changes the underlying slice (s1). So while I was looking for something like
	// copyFrom or startFrom in slices pack, this way was available all the time.
	// apparently not many folks know because I couldn't find it on other tutorial
	return s1 //  (Restrictions: No for loops)
}

func appendInt(s []int, v int) []int {
	i := len(s)

	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1] // since value is one at a time, its adding 1 slot until its at cap
	} else { // need to grow the underlying array and copy over.
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1) // create new longer slice to copy original into
		copy(s2, s)                   // copy will always work since it copies what it can. It
		// returns the amount of items it was able to copy. (It is a shadow copy most times)
		s = s2[:len(s)+1] // then we rename it back to original slice
	}
	s[i] = v
	return s
} // IMPORTANT: The specifics on the builtin append are undefined and the Go team
// can change it as they need. This means that it'll mostly start doubling the
// cap, but then it'll move to smaller amounts, maybe 20%. Its what I originally
// thought that there is an arbitrary formula used to figure it out. ITs not
// doubled everytime.

// our append func is roughly what the append algo is doing. it can be O(1) if its just
// placing something in the slice, but from time to time it needs to creat a new slice and
// copy the entire thing with makes it O(n) of space complexity (not time)

// If you know in advance the memory you will need, use the 3rd arg in make for slices:
// make([]int,0,1_000) as it'll be a single allocation as can be seen in the example.
