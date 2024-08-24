/*
Create a program that declares two anony funcs. One that counts down from
100 to  0 and one that counts up from 0 to 100. Display each number with an
unique identifier for each goroutine. Then create goroutines from these funcs
and don't let main return until the goroutines complete.

Run the program in parallel (if I have a max of 1 proc, I don't think I can technically
run it in parallel, so I'm going to change that below.)
*/
package main 

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// I'm allocating 2 logical processors for the scheduler to use
	runtime.GOMAXPROCS(2)
	// fmt.Println(runtime.GOMAXPROCS(0))
}

func main() {

	// 1. Declare a wait group and set the count to two
	var wg sync.WaitGroup
	wg.Add(2) 
	
	// 2. Declare an anony func and create a goroutine
	go func(){
		
		// 3. Declare a loop that counts down from 100 to 0 and displays
		//    each value
		for i:=100;i>=0;i--{
			fmt.Printf("[A:%d]\n",i)  
		}
		
		// 4. Tell main we are done
		wg.Done() 
	}()
	
	// 5. Declare another anony func and create a goroutine
	go func(){
		
		// 6. Declare a loop that counts up from 0 to 100 and displays
		//    each value
		for i:=0;i<=100;i++{
			fmt.Printf("[B:%d]\n",i)  
		}
		
		// 7. Tell main we are done
		wg.Done() 
	}()
	
	// 8. Wait for the goroutines to finish.
	wg.Wait() 
	
	// Display "Terminating Program"
	fmt.Println("Terminating Program") 
	
	
}
