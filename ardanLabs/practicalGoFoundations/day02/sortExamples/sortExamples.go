package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implemnts sort.Interface for []Person based on the Age field
type ByAge []Person

// then we create the method set to be a sorter
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }    // key to order
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } //key to which field

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)

	// There are 2 ways to sort a slice. First, we can define the set of methods
	// as I did above for the slice type and call sort.Sort. In the example above
	// we use that technique and can do the following:
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less function, which can
	// be provided as a closure/callback. With this no methods are needed and if
	// they are implemented, they are just ignored. Here we re-sort  but in reverse
	// order
	fmt.Println()
	fmt.Println("Let's use a closure arg and sort in reverse:")
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
		// here we reversed expression i > j. We could have also done j < i
	})
	fmt.Println(people)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Sorting a struct type using programmable sort criteria:")
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return distance(p2, p1)
	}

	// Sort the planets by the various criteria
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)
}

// Now let's demonstrate a technique of sorting a struct type using programmable
// sort criteria

// First a couple of type defs to make the units clear
type earthMass float64
type au float64

// A planet defines the properties of a solar system object
type Planet struct {
	name     string
	mass     earthMass // the types I created
	distance au        // the types I created. They are floats but make it clearer
}

// By is the type of a "less" function that defines the ordering of its planet args
type By func(p1, p2 *Planet) bool // first time I create a func type
// I'm thinking this is how we create funcs like sort.Slice with func args ???

func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // the sort method's receiver is the func closure that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of planets to be sorted
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Closure used in the less method
}

// Len is part of the sort interface
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of the sort interface
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of the sort interface. It is implemented by calling the by closure.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}
