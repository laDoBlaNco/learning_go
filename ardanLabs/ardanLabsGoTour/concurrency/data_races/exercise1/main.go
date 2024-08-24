// fix the race condition in this program

package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time" 
)

// numbers maintains a set of random numbers
var numbers []int

// let's create a mutex
var mux sync.Mutex

// init is called prior to main
func init(){
	rand.Seed(time.Now().UnixNano()) 
}

func main(){
	
	// number of grs to use
	const grs = 3
	
 
	
	// wg is used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(grs) 
	
	// create three grs to generate random numbers
	for i:=0;i<grs;i++{
		go func(){
			random(10)
			wg.Done() 
		}()
	}
	
	// Wait for all the goroutines to finish
	wg.Wait() 
	
	// Display the set of random numbers
	for i,number:=range numbers{
		fmt.Println(i,number) 
	}
}

// random generates random numbers and stores them into a slice
func random(amount int){
	
	// generate as many random numbers as specified
	for i:=0;i<amount;i++{
		n:=rand.Intn(100)
		mux.Lock()
		numbers = append(numbers,n)
		mux.Unlock()  
	}
}
