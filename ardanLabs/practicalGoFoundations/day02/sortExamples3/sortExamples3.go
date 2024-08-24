package main

import (
	"fmt"
	"sort"
)

// This example is using a sort wrapper
type Grams int

func (g Grams) Sting() string {
	return fmt.Sprintf("%dg", int(g))
}

type Organ struct {
	Name   string
	Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int {
	return len(s)
}
func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ByName will finish implementing the sort.Interface using the Less method directly
// and they promoted Len and Swap methods from the embedded Organs value
type ByName struct {
	Organs
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

// Same for ByWeight...
type ByWeight struct {
	Organs
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	// here we are going to create the organ slice and with our wrappers which are
	// basically allowing us to generalize the Swap and Len methods as embedded and
	// wrap them in our custom BySomething type.
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s}) // why are we using {} instead of ()???
	fmt.Println("Organs by weight:")
	printOrgans(s)

	fmt.Println()

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)

}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}
