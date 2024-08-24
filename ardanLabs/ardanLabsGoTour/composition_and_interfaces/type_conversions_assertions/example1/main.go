package main

import "fmt"

/*

	TYPE CONVERSIONS AND ASSERTIONS
	So a type conversion allows the data of one t ype to convert to another type. Go prefers
	conversion over casting. A type assertion allows us to ask the question if there is a value
	of the given type stored inside an interface.
	
	First let's look at implicit interface conversions. As we saw in a previous example
	An interface value of one type can be passed for a different interface type if the
	concrete value stored inside the interface implements both behaviors. This could
	be considered an implicit interface conversion, but its better to think about how
	concrete data is being moved through interfaces in a decoupled state.
	

*/

type Mover interface{
	Move()
}

type Locker interface{
	Lock()
	Unlock() 
}

type MoveLocker interface{
	Mover
	Locker
}

// our bike represents a concrete type for this example
type bike struct{} 


// then we implement the interfaces
func(bike)Move(){
	fmt.Println("Moving the bike") 
}

func (bike)Lock(){
	fmt.Println("Locking the bike") 
}

func(bike)Unlock(){
	fmt.Println("Unlocking the bike") 
}

func main(){
	// declare vars of teh MoveLocker and Mover interfaces set to their zero value
	var ml MoveLocker 
	var m Mover 
	
	// create a value of type bike and assign it to MoveLocker interface value
	ml = bike{} 
	
	// An interface value of type MoveLocker can be implicitly converted into a value of type
	// Mover. They both declare a method named Move
	m = ml 
	
	// but it won't work the other way as Move doesn't have Lock and Unlock
	 // ml = m 
	// example1/main.go:65:8: cannot use m (variable of type Mover) as MoveLocker value in assignment:
	// Mover does not implement MoveLocker (missing method Lock)

	// Interface type Mover does not declare methods named Lock or Unlock. Therefore the compiler
	// doesn't perform an implicit conversion to assign a value of interface type Mover to
	// an interface value of type MoveLocker. It is irrevelant that the concrete type value of 
	// type bike that is stored inside of the Mover interface value implements the MoveLocker
	// interface. 
	
	// We can perform type assertion at runtime to support the assignment. performing type assertion
	// against the Mover interface value to access a COPY of the concrete type value of the type
	// bike that was stored inside of it. Then assign the COPY of the concrete type to the MoveLocker
	// interface.
	b:=m.(bike) 
	ml = b 
	
	// its important to note that the tpe assertion syntax provides a way to state what type of the
	// value is stored inside the interface. This is more powerful from a language and readability
	// standpoint, than using casting syntax, like in other languages. 
	
	
}

// Type assertion mainly allows us at runtime to ask a question, is there a value of the given
// type stored inside this interface. We see the syntax m.(bike). In this case we are asking
// if there is a bike value stored inside of m at the m oment the code is executed. If there is
// then the variable b is given a copy of the bike value stored. Then the copy can be copied 
// inside of the ml interface varible instead of m. 

// If there isn't a bike value stored inside the interface value, then the program will panic. This
// is what we want so that there is absolutely no question if there is or is not a bike value stored.
// What if there is a chance there isn't and that is valid? Then we need the second form of the type 
// assertion with is the comma,ok idiom
// b,ok := m.(bike). If this form is ok = true, there is a bike value stored in the interface, and
// if ok is false, then there isn't, but the program doesn't panic. NOTE the var will still be a 
// type bike, but it'll be at its zero default value. 

