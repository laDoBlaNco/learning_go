package main

import(
	"fmt"
	_"log"
	_"os"
	_"os/exec"
)

// ... for quick debugging
var p = fmt.Println


// 'range' iterates over elements in a variety of data structures. let's see how
// to use range with some of the data structures we've already learned.

func main(){
	
	// First here we use range to sum the numbers in a slice. Arrays work the
	// same as well.
	nums := []int{2,3,4} 
	sum := 0
	for _,num := range nums {
		sum+=num
	}
	p("Sum:",sum) 
	p() 
	
	// range on arrays and slices provides both the index and value of each
	// entry. Above we didn't need the index, so we ignored it with the blank
	// identifier _. Sometimes we actually want the indexes as well or instead
	for i,num:=range nums{
		if num == 3{
			p("index for 3:",i) 
		}
	}	
	p()
	
	// we can range on a map which will iterate over keys/values instead of index/
	// value.
	kvs := map[string]string{"a":"apple","b":"banana","c":"carrot"} 
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n",k,v) 
	}
	p()
	
	// range can also iterate over just the keys of a map (its the same with any
	// data collection. 2 items are returned (i,v) or (k,v) and if you want just
	// first then you can use only that. But if you want just the second you must
	// use _,v) 
	for k := range kvs{
		p("Keys:",k) 
	}
	p() 
	
	// range on strings iterates over Unicode code points, not just the bytes. The
	// first value is the starting byte index of the rune and the second is the
	// rune it self. 
	for i,c := range "G😀"{
		p(i,c) // Note this will print the bytes o sea the numbers, not the actual
		// rune
		p(i,string(c)) // for that we need to cast it back to a string.
	}
	p() 
	
}
