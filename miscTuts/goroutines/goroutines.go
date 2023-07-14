package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(c chan int, co chan int) {
	for {
		x := <-c
		y := <-c
		fmt.Println(x, "+", y, "=", x+y)
		co <- x + y
	}

}

func main() {
	c := make(chan int, 6)
	co := make(chan int, 3)
	go sum(c, co)
	c <- 10
	c <- 15
	c <- 99
	c <- 1
	c <- 23
	c <- 7

	r := <-co
	fmt.Println(r)
	r = <-co
	fmt.Println(r)
	r = <-co
	fmt.Println(r)

	fmt.Println("____________________________________________")

	msg := primemsg{42, false}

	// create channels
	cin := make(chan primemsg, TEST_LEN)
	cout := make(chan primemsg, TEST_LEN)

	// create workers
	for i := 0; i < NUM_WORKERS; i++ {
		go isPrime(cin, cout)
	}

	// Fill the input channel
	for i := 0; i < TEST_LEN; i++ {
		msg.num = rand.Intn(1000000) + 1000000
		cin <- msg
	}

	// Read the answers
	for i := 0; i < TEST_LEN; i++ {
		msg = <-cout
		fmt.Println(msg.num, msg.isPrime)
	}
}

// SquareRoot / Primes
type primemsg struct {
	num     int
	isPrime bool
}

const TEST_LEN = 100
const NUM_WORKERS = 3

func isPrime(cin chan primemsg, cout chan primemsg) {
	id := rand.Intn(1000000)
	i := 0
	for {
		msg := <-cin
		num := msg.num
		fmt.Println(id, "is testing", num)
		sq_root := int(math.Sqrt(float64(num)))
		for i = 2; i < sq_root; i++ {
			if num%i == 0 {
				msg.isPrime = false
				cout <- msg
				break
			}
		}
		if i > sq_root {
			msg.isPrime = true
			cout <- msg
		}
	}

}

// THIS DIDN'T WORK. I'LL COME BACK TO IT DOWN THE ROAD TO SEE IF i CAN FIGURE OUT
// ITS DEADLOCKING.
