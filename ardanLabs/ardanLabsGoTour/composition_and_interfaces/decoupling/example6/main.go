package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// There is another change we can make to Copy below...

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

// type System struct{
// 	Puller
// 	Storer
// }

 

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

// This change makes the function even more precise and polymorphic again. Now the function is
// asking for exactly what it needs based on what the concrete data can do. Which this change
// we can now remove the System interface as the function is asking for the puller storer
// directly.
func Copy(p Puller,s Storer,batch int)error{
	data:=make([]Data,batch) 
	for{
		i,err := Pull(p,data)
		if i>0{
			if _,err:=Store(s,data[:i]);err!=nil{
				return err
			}
		}
		
		if err!=nil{
			return err
		}
	}
}



func main(){
	
	x := Xenia{
		Host:"localhost:8000",
		Timeout:time.Second,
	}
	p:=Pillar{
		Host:"localhost:9000",
		Timeout:time.Second,
	}
	
	if err:=Copy(&x,&p,3);err!=io.EOF{
		fmt.Println(err) 
	}
}

/*
	By removing the PullStorer and System types, the program simplifies. The main function
	can focus on constructing the concrete Puller and Storer values necessary for that moving
	data. The type system and APIs are  more precise. This idea of precision comes from Edsger
	W. Dijkstra.
	
	"The purpose of abstraction is not to be vague, but to create a new semantic level in which
	one can be absolutely precise" -- Edsger W. Dijkstra
	
	NOTE:	
	- This is much more than the mechanics of type embedding
	- Declare types and implement workflows with composition in minde
	- Understand the problem we are trying to solve first. This means understanding the data
	- The goal is to reduce and minimize cascading changes across your software
	- Interfaces provide the highest form of composition since we can combine virtually anything
	  as long as we focus on behavior and not state (or what we are) 
	- Don't group types by a common DNA but by a common behavior
	- Everyone can work together when we focus on what we do and not what we are.


*/


