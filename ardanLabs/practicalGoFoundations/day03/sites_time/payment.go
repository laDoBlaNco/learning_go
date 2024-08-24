package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := Payment{
		From:   "Wile E. Coyote",
		To:     "ACME",
		Amount: 123.34,
	}
	p.Process()
	p.Process()
}

func (p *Payment) Process() {
	t := time.Now()                    // below we also see a use of anony funcs 
	p.once.Do(func() { p.process(t) }) // we pass just the function (not the function call())
	// with this we don't have to worry about bool flags and mutexs, we use idempotence to
	// just do it once and no matter how many times its called after that, it won't do anything.
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s", ts, p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once // its lowercase because its an implementation detail and not exported.
}
