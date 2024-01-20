package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	// first create the Mutex in our account. Note this is in the same sync pkg
	Mutex *sync.Mutex // we put this here cuz we are adding the security to our overall
	// account and it comes with its on methods. The mutex needs to be part of struct because
	// the goroutines are going to be coming to this struct (account). Outside of this, each
	// of the thousands or millions of routines aren't going to know about the mutex. they need
	// to find the security in one common place. So they will come here to the account they need
	// to work on, and if it includes security, they will use it. if there's no security, they
	// still do their work, but there will be race conditions
}

// to guard against the race condition we first need to identify the critical pieces of our
// code that need protect. We do this by wrapping those lines of code in a Mutex

// NOTE: we also have to be careful of mutex Deadlock, similar to normal go routin deadlock where
// we are blocking and not releasing anywhere, which is if we lock our code but don't unlock it.
// we can see this happen if we comment out our a.Mutex.Unlock() below
// fatal error: all goroutines are asleep - deadlock!

// ANOTHER NOTE: Putting in mutexes does impact performance cuz you are making your async app
// into a sync app because funcs can only be run one after the other. But its safety vs performance.

func (a *Account) Withdraw(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()         // its kinda like putting a padlock on the account
	a.Balance -= value / Needs protection
	a.Mutex.Unlock()     // and taking it off.
	wg.Done()
}

func (a *Account) Deposit(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += value // Needs protection
	a.Mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Mutexes in Go!")
	// Mutex stands for Mutual Exclusion. Its used to guard the critical sections of our code
	// ensuring race conditions don't impact out code. In other words we aren't updating our data
	// in the wrong order or at the same time.

	// Note that mutexes are about protecting data and WaitGroups are about ensuring all the work
	// is done. One doesn't replace the other.

	// Let's look at this using a bank account

	var m sync.Mutex // create are actual mutex

	account := Account{
		Balance: 1000,
		Mutex:     &m, // add it to the accoun.
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go account.Withdraw(700, &wg)
	go account.Deposit(500,  &wg)
	wg.Wait()

	fmt.Println("Account Balances Updated")
	fmt.Println(account.Balance)

}
