package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type PullStorer interface {
	Puller
	Storer
}

type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("error reading data from Xenia")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Pillar struct{
	Host string
	Timeout time.Duration
}

func(*Pillar)Store(d *Data)error{
	fmt.Println("Out:",d.Line)
	return nil
}
// Since a system is composed from a Xenia and Pillar, System implements the PullStorer interfacd
// With these changes, we can now create new concrete types that implement the PullStorer interface
// When we think about this more, declaring different System types for all possible combinations is
// not realistic. It will work, but the maintenance nightmare requires a better solution.

// We could decide to COMPOSE a concrete system type from to interfaces. This is an intersting 
// solution because it would allow us to inject the concrete Pillar and Storer into the
// system at application startup
type System struct{
	Puller
	Storer
}

// So the idea is creating an interface composed of interfaces we are creating a concrete type
// composed of interfaces. 

func Pull(p Puller,data []Data)(int,error){
	for i:=range data{
		if err := p.Pull(&data[i]);err!=nil{
			return i,err
		}
	}
	return len(data),nil 
}

func Store(s Storer,data []Data)(int,error){
	for i:=range data{
		if err := s.Store(&data[i]);err!=nil{
			return i,err
		}
	}
	return len(data),nil
}

func Copy(ps PullStorer,batch int)error{
	data:=make([]Data,batch) 
	for{
		i,err := Pull(ps,data)
		if i>0{
			if _,err:=Store(ps,data[:i]);err!=nil{
				return err
			}
		}
		
		if err!=nil{
			return err
		}
	}
}



func main(){
	
	sys := System{
		Puller:&Xenia{
			Host: "localhost:8000",
			Timeout:time.Second,
		},
		Storer:&Pillar{
			Host:"localhost:9000",
			Timeout:time.Second,
		},
	}
	
	if err:=Copy(&sys,3);err!=io.EOF{
		fmt.Println(err) 
	}
}

// and it all still works. AS we can see we are now able to inject our pointer to Xenia and 
// Pillar directly into our System of interfaces. This one system type implements the PullStorer
// interface for all possible combinations of concrete types. As long as they fulfill the 
// method requirements. 

