package main

import(
	"fmt"
	"time"
	"sync" 
)
// ...for quick debugging
var pl = fmt.Println

// Go allows us to have multiple blocks of code running at the same time
// which is parallelism or sharing the same runtime thread or routine which is 
// actually concurrency

func printTo15(){
	for i:=1;i<=15;i++{
		pl("Func 1:",i)
	}
}
func printTo10(){
	for i:=1;i<=10;i++{
		pl("Func2 1:",i)
	}
}
func nums1(channel chan int){
	channel <-1
	channel <-2
	channel <-3
} 
func nums2(channel chan int){
	channel <-4
	channel <-5
	channel <-6
}


type Account struct{
	balance int
	lock sync.Mutex // this is mutual exclusivity that basically allows us to lock and unlock
	// access to values.
}

func(a *Account) GetBalance()int{
	a.lock.Lock()
	defer a.lock.Unlock() // unlock after the transaction takes place
	return a.balance
}

func(a *Account)Withdraw(v int){
	a.lock.Lock()
	defer a.lock.Unlock() 
	
	if v > a.balance{
		pl("Not sufficient funds in Account")
		return
	}
	fmt.Printf("%d Withdrawn : Balance : %d\n",v,a.balance) 
	a.balance-=v
}

func main(){

	go printTo15()
	go printTo10() 
	
	// the main function is running on itself as well, so it will close once done
	// and not wait for the other routines. so we need to make it pause for a bit
	// to give time to the others. 
	
	// time.Sleep(2*time.Second)  // in order to not need to do this we can assign
	// the main goroutine (main) receive the values from our goroutines. this is done
	// with channels.
	pl()
	
	channel1:=make(chan int)
	channel2:=make(chan int)
	
	go nums1(channel1)
	go nums2(channel2) 
	
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)
	// The above as I saw in the bootdev  gives us  more control around receiving 
	// data one at a time only when the channel is called upon.
	
	// Let's simulate bank count transactions.
	pl() 
	var acct Account
	acct.balance = 100
	pl("Balance:",acct.GetBalance()) 
	for i:=0;i<12;i++{
		go acct.Withdraw(10) 
	}
	time.Sleep(2*time.Second) 


}
