package main

// A sample program to show the syntax and mechanics of type switches (which have the same
// relation as regular switches to 'if statements', these are type switches on if statements
// that use type assertion) and empty interfaces

import "fmt"

func main(){
	
	// fmt.Println can be calld with values of any type. Why???
	fmt.Println("Hello world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true) 
	
	// Let's do the same using typeswitches and empty interfaces
	myPrintln("Hello world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true) 

}

/*
	- An interface is satisfied by any piece of data when thte data exhibits the full method
	  set of behavior defined by the interace.
	- The EMPTY INTERFACE defines no method set of behavior and therefore requires no 
	  specific method by the data being stored. O sea, all methods fall into its interface
	- The empty interface says nothing about the data stored inside the interface
	- Checks would need to be performed at runtime to know anything about it, since it aceepts
	  everything
	- We need to decouple around well defined behavior and only use the empty interface as an
	  exception when its is reasonable and practical to do so. In face in most cases I think
	  generics has replaced the need for this as generics is basically just based on the 
	  empty interface in its roots. 
*/

func myPrintln(a interface{}){ // with interface{}, this func can take any type at all
	switch v := a.(type){ // therefore we need to use type switch (type assertion) to determine
	// what we are actually dealing with 
	case string:fmt.Printf("Is string : type(%T) : value(%s)\n",v,v)
	case int:fmt.Printf("Is int : type(%T) : value(%d)\n",v,v)
	case float64:fmt.Printf("Is float64 : type(%T) : value(%f)\n",v,v)
	default:fmt.Printf("Is unknown : type(%T) : value(%s)\n",v,v)
	}
}
