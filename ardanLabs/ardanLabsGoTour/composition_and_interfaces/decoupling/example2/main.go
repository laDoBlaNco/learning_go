package main

import(
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time" 
)

/* 
	The next step is to understand what could change in the program. In this case, what can change
	is the systems themselves. Today its Xenia and Pillar, tomorrow it could be Alice and Bob.
	with this knowledge, we want to decouple the existing concrete solution from this future
	change. To do that, we want to change the concrete functions to be polymorphic functions. 
	NOTE the use of the contrasts here, concrete functions vs polymorphic functions. 
*/

func init(){
	rand.Seed(time.Now().UnixNano()) 
}

type Data struct{
	Line string
}

// here we create our interfaces to be able to create polymorphic functions. Remember that before
// the Pull function accepted a Xenia value and a Pillar value, but in the end, it wasn't Xenia
// or Pillar that were important, what's important is a concrete value THAT KNOWS HOW TO PULL AND
// STORE. So we change the concrete functions to polymorphic by asking for data based on
// WHAT IT CAN DO INSTEAD OF WHAT IT IS.
// Puller declares the behavior or method list for pulling data
type Puller interface{
	Pull(d *Data)error
}

// Storer declares behavior for storing data
type Storer interface{
	Store(d *Data)error 
}

type Xenia struct{
	Host string
	Timeout time.Duration
}

// now our Xenia type implements the Puller interface. (it is a puller)
func(*Xenia)Pull(d *Data)error{
	switch rand.Intn(10){
	case 1,9:return io.EOF
	case 5:return errors.New("error reading data from Xenia") 
	default:
		d.Line = "Data"
		fmt.Println("In:",d.Line)
		return nil 
	}
}

type Pillar struct{
	Host string
	Timeout time.Duration
}

// Store knows how to store, making it fill the contract and become a storer
func(*Pillar)Store(d *Data)error{
	fmt.Println("Out:",d.Line) 
	return nil 
}

// Then we have our system
type System struct{
	Xenia
	Pillar 
}

// Now in this example Pull is asking for a Puller, o sea any type that fills the puller 
// interface
func Pull(p Puller,data []Data)(int,error){
	for i:=range data{
		if err:=p.Pull(&data[i]);err!=nil{
			return i,err 
		}
	}
	return len(data),nil
}

// Same as above for Store
func Store(s Storer,data []Data)(int,error){
	for i:=range data{
		if err:=s.Store(&data[i]);err!=nil{
			return i,err 
		}
	}
	return len(data),nil 
}
// Now these functions are polymorphic. When Alice and Bob are declared and implemented as a 
// Puller and a Storer, they can also be passed to these new polymorphic functions. 

// Copy again gets promoted methods to pull and store data to THE system. 
func Copy(sys *System,batch int)error{
	data:=make([]Data,batch) 
	
	for{
		i,err:=Pull(&sys.Xenia,data)
		if i>0{
			if _,err:=Store(&sys.Pillar,data[:i]);err!=nil{
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
