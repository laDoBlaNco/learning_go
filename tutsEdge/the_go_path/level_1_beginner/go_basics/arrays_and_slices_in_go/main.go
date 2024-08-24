package main

import "fmt"

func main() {
	println()
	println("Arrays and Slices in Go")
	// arrays is a container for element types that we want to store in Go
	// [1,2,3,4]
	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planets)

	// above was the shorthand, but the more verbose approach is here:
	var planetsArray [8]string
	planetsArray[0] = "mercury"
	fmt.Println(planetsArray)
	// There is an even more verbose approach with is using the 'make' func. But we'll get into
	// that later I guess.

	// Slices:
	// main difference is that we don't specify the size of the slice because they are mutable
	// Every slice has an underlying array and based on the len and cap of that underlying array
	// the slice can be adjusted under the hood to fit whatever we give to it.
	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	// append func handles the resizing of the underlying array for us when needed.
	planetSliceVerbose = append(planetSliceVerbose, "mercury")
	fmt.Println(planetSliceVerbose)
}
