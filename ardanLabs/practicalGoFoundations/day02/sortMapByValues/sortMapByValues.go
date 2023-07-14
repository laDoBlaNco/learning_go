package main

import (
	"fmt"
	"sort"
)

// In this tut we are going to learn how to sort a map by key or values

func main() {

	// first we create our map[string]int (fruit->count)
	basket := map[string]int{
		"orange":     5,
		"apple":      7,
		"mango":      3,
		"strawberry": 9,
	}

	// get a slice of keys from our map
	keys := make([]string, 0, len(basket))
	for k := range basket { // remember with maps it k,v not i,v
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// after sorting the slice of keys we just list the keys in order with their vals
	fmt.Println("Sorting by Keys ascending:")
	for _,k:=range keys{// this one is i,k since its a slice and not a map
		fmt.Println(k,basket[k]) 
	}
	
	
	sort.Sort(sort.Reverse(sort.StringSlice(keys))) 
	fmt.Println()
	fmt.Println("Sorting by Keys descending:") 
	for _,k:=range keys{// this one is i,k since its a slice and not a map
		fmt.Println(k,basket[k]) 
	}
	
	// So we still need a slice of the keys, but we can then use sort.SliceStable with
	// a Less func anony to sort those keys by their values. 
	
	sort.SliceStable(keys,func(i,j int)bool{
		return basket[keys[i]]<basket[keys[j]] 
	})
	
	fmt.Println()
	fmt.Println("Sorting by Values ascending:")
	for _,k:=range keys{
		fmt.Println(k,basket[k]) 
	}

	// reverse the sort just by changing the '<' to '>'
	sort.SliceStable(keys,func(i,j int)bool{
		return basket[keys[i]]>basket[keys[j]] 
	})

	fmt.Println()
	fmt.Println("Sorting by Values descending:")
	for _,k:=range keys{
		fmt.Println(k,basket[k]) 
	}


	
}
