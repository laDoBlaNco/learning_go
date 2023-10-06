// a quick program to show how maps ARE NOT THREAD SAFE
package main

import(
	"fmt"
	"sync" 
)

// scores will hold values incremented by multiple grs
var scores = make(map[string]int) 

func main(){
	
	var wg sync.WaitGroup
	wg.Add(2) 
	
	go func(){
		for i:=0;i<1000;i++{
			scores["A"]++
		}
		wg.Done()
	}()
	
	go func(){
		for i:=0;i<1000;i++{
			scores["B"]++
		}
		wg.Done() 
	}()
	
	wg.Wait() 
	fmt.Println("Final scores:",scores) 
}
