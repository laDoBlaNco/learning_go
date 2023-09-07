package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// The next question to ask is if our polymorphic functions are as precise as they otherwise
// could be?  This is a part of the engineering process that can't be skipped. The answer is
// NO, two changes can be made below. First with Copy...

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

// type PullStorer interface {
// 	Puller
// 	Storer
// }

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

type System struct{
	Puller
	Storer
}

 

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

// Copy no longer needs to be polymorphic since there is only one System type which covers
// all possible combinations of PullStorers. The PullStorer interface type can be removed 
// from the program altogether. NOTE: Remember that we moved the polymorphism INSIDE of the 
// concrete type when we composed the interface types. So we put it back to *System which
// again is now a polymorphic type
func Copy(sys *System,batch int)error{
	data:=make([]Data,batch) 
	for{
		i,err := Pull(sys,data)
		if i>0{
			if _,err:=Store(sys,data[:i]);err!=nil{
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


