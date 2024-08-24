package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// We aren't done yet. In the last example we created polymorphic functions by creating
// interfaces for Puller and Storer. We can also change a couple of things and create
// a PullStorer which will make our Copy func polymorphic. Then it just work with THE System,
// but with ANY system that knows how to Pull AND Store.

// NOTE: I'm not going to comment the stuff I commented in the first 2 examples

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

// Now we can simply embed our interfaces just as we can embed our concrete types and
// combine promoted behavior lists. The PullStorer interface is declared through the use of
// composition. Its composed of the Puller & Storer interfaces. We work towards composing larger
// interfaces from smaller ones. 
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

type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

type System struct{
	Xenia
	Pillar 
}

func Pull(p Puller,data []Data)(int,error){
	for i:=range data{
		if err:=p.Pull(&data[i]);err!=nil{
			return i,err
		}
	}
	return len(data),nil 
}

func Store(s Storer,data []Data)(int,error){
	for i:=range data{
		if err:=s.Store(&data[i]);err!=nil{
			return i,err 	
		}
	}
	return len(data),nil 
}

// And here we can now turn our Copy function into a polymorphic function by letting it
// accept PullStorers rather than just THE system
// Also NOTE how the PullStorer var 'ps' is being passed into the Pull and Store functions. How
// can this be??? We always need to remember that we are never passing an interface value around
// the program sinc they don't exist and are valueless, we can only pass concrete data. So the
// concrete data stored inside of the interface ps variable is what's being passed to Pull and
// Store. Isn't it true, the concrete value stored inside of ps must know to Pull and Store?
// Since a System is composed from Xenia and Pillar, System implements the PullStorer interface
func Copy(ps PullStorer,batch int)error{
	data := make([]Data,batch) 
	
	for{
		i,err:=Pull(ps,data)
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
		Xenia:Xenia{
			Host:"localhost:8000",
			Timeout:time.Second,
		},
		Pillar:Pillar{
			Host:"localhost:9000",
			Timeout:time.Second,
		},
	}
	if err:=Copy(&sys,3);err!=io.EOF{
		fmt.Println(err) 
	}
}

// With these changes, we can now create new concrete types that implement the PullStorer interface.



