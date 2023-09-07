package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

/*
	The best way to take advantage of embedding is through the compositional design pattern.
	Its the key for maintaining stability in your software by having the ability to adapt to
	the data and transformation changes that are coming

	Here we will create a sample program demonstrating struct composition

*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data is a struct of the data we are copying
type Data struct {
	Line string
}

// Xenia represents a system that you need to pull data from. The implemenation is not important
// for what we want to do. What is important is that the method Pull can succeed, fail, or not
// have any data to pull, o sea pull empty
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Pull knows how to pull data out of Xenia
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF // so this is the empty pull
	case 5:
		return errors.New("error reading data from Xenia") // this is the failed pull
	default:
		d.Line = "Data" // and this is the pull of actual data
		fmt.Println("In:", d.Line)
		return nil
	}
}

// The Pillar type represents a system that we need to store data into. What is important here
// is again that the method Store can succeed or fail
type Pillar struct{
	Host string
	Timeout time.Duration
}

func(*Pillar)Store(d *Data)error{
	fmt.Println("Out:",d.Line)
	return nil 
}

// System wraps Xenia and Pillar together into a single system. These two types represent a
// primitive layer of code that provides the base behavior required to solve the business
// problem of pulling data out of Xenia and storing that data into Pillar. 
type System struct{
	Xenia
	Pillar 
}

// The next layer of code is represented by two functions, Pull and Store. They build on the
// primitive layer of code by accepting a collection of data values to pull or store in the
// respective system. These functions focus on the concrete types of Xenia and Pillar since 
// those are the systems the program needs to work with at this time.
func Pull(x *Xenia,data []Data)(int,error){
	for i:=range data{
		if err:=x.Pull(&data[i]);err!=nil{ // we use &data cuz the method looks for a pointer to data
			return i,err 
		}
	}
	
	return len(data),nil 
}

func Store(p *Pillar,data []Data)(int,error){
	for i:=range data{
		if err := p.Store(&data[i]);err!=nil{
			return i,err 
		}
	}
	
	return len(data),nil 
}

// Now the Copy function builds on top of the Pull and Store functions to move all the data that
// is pending for each run. NOTE the first parameter to Copy, its type call System
func Copy(sys *System,batch int)error{
	data:=make([]Data,batch) 
	
	for{
		i,err:=Pull(&sys.Xenia,data) 
		if i>0{
			if _,err := Store(&sys.Pillar,data[:i]);err!=nil{
				return err
			}
		}
		if err!=nil{
			return err
		}
	}
}

// Now in the main function we can write to construct a Xenia and Pillar within the composition
// of a system. Then the system can be passed to the Copy function and data can begin to flow
// between the two systems

// This is our first draft of a concrete solution to our concrete problem

func main(){
	
	sys := System{
		Xenia:Xenia{
			Host:"localhost:8000",
			Timeout: time.Second,
		},
		Pillar:Pillar{
			Host:"localhost 9000",
			Timeout:time.Second,
		},
	}
	
	if err := Copy(&sys,3);err!=io.EOF{
		fmt.Println(err) 
	}
	
}
// In the next example we are going to decouple our concrete implementation from future 
// changes with interfaces.

