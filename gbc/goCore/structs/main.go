package main

import (
	"fmt"
)

// ...for quick debugging
var pl = fmt.Println

// Structs allow you to store values with many different data types and also
// associated functions (not methods yet) to work with them

type customer struct { // its a custom data type
	name    string
	address string
	bal     float64
}

// functions (not yet methods) that are associated with these custom types
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}
func newCustAdd(c *customer, address string) { // note the use *customer since we want
	// to change the actual customer passed in and not the copy
	c.address = address
}

// Now let's create a shape
type rectangle struct{
	length,height float64
}

func (r rectangle) Area() float64{ // note now we are using what are called methods
// in Go. Its connected to the type not just receiving it as an arg
	return r.length*r.height
}

// Some people like to compare Go structs to objects in other languages. In other
// languags with OOP you have 'inheriting' while in Go there is no inheritance. There
// is embedding or composition
type contact struct{
	fName string
	lName string
	phone string
}

type business struct{
	name string
	address string
	contact
}

func (b business)info(){
	fmt.Printf("Contact at %s is %s - %s\n",
		b.name,
		b.contact.fName,
		b.phone) // note it wasn't necessary to do b.contact.phone
}

func main() {

	var ts customer
	ts.name = "Tom Smith"
	ts.address = "5 Main St"
	ts.bal = 234.56

	getCustInfo(ts)
	newCustAdd(&ts, "123 South St") // note that we had to pass a ref since we are
	// using a func that we created to take a pointer
	pl("Address:", ts.address)

	pl() 
	// we can also create struct literals
	ss := customer{"Sally Smith", "123 Main St", 0.0}
	pl("Name:", ss.name)
	
	pl()
	rect1:=rectangle{10.0,15.0}
	pl("Rect Area:",rect1.Area()) 
	
	pl()
	con1:=contact{
		"James","Wang","555-1212",
	}
	bus1:=business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info() 
	
	
}
