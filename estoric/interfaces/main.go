package main

import (
	"fmt"
	"math"
	"reflect"
)

// So an interface is like a job description. The things the function or
// person has to do to satisfy the requirement of the job description or the
// interface.

// Two of the most used interfaces in Go are the io.Writer & io.Reader
// interfaces.

// The only thing that matters is that the requirements are filled. the "role"
// or type doesn't matter. It also doesn't matter if the type can do more than
// what the requirements are asking for.

// An interface allows us to speak in more general terms (Abstract type)

// If there were no interfaces everything would have to be created by type
// and only accepts  certain types and really locked down. Everything being
// from scratch. But interfaces opens the world of being able to plug different
// types into what we are building, without actually knowing what the type is
// at first.

// We are going to use 'shapes' to practice interfaces

type shape interface { // we create an interface with type and put inside
	Area() float64 // the requirements for the type we will use or "give the job"
	Perimeter() float64
}

type polygon interface {
	GetSides() int
}

var i interface{} // note we use var with an empty interface as we are just assigning??

type rectangle struct {
	width, height float64
	sides         int
}

type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) GetSides() int {
	return r.sides
}

func Calculate(s interface{}) {
	//fmt.Printf("My type is %T\n", s)// the type is there we just need to pull it out
	//fmt.Printf("My value is %v\n", s)
	if rect, ok := s.(circle); ok { // so type assert  returns 2 results like functons

		fmt.Printf("My type is %T\n", rect)
		fmt.Println(rect.Perimeter())
		//fmt.Println(s.Area())
		//fmt.Println(s.Perimeter())
	} // here we do type assertion and save underlying type to 'rect'

	//fmt.Println(rect.GetSides()) // now we can call rect's GetSides method
}

func calculate2(s interface{}) {
	// myType := reflect.TypeOf(s) // We return a type Type that gives us views
	// fmt.Println(myType.Name())  // into the underlying type
	// fmt.Println(myType.Kind())  // here we get underlying 'kind'
	// fieldInfo := myType.Field(0)
	// fmt.Println(fieldInfo.Name)
	myType := reflect.ValueOf(s)
	fmt.Println(myType)
	switch x := s.(type) { // note we can't use 'fallthrough' with type switch statements
	case circle:
		fmt.Printf("I am a circle and my value is %v\n", x)
	case rectangle:
		fmt.Printf("I am a rectangle and my value is %v\n", x)
	case int:
		fmt.Printf("I am an int and my value is %v\n", x)
	default:
		fmt.Println("None of the above types work for us")
	}
	fmt.Println("____________________________________________")
}

func main() {
	r := rectangle{width: 5, height: 8, sides: 4}
	c := circle{radius: 6}
	// n := 2

	calculate2(r)
	calculate2(c)
	// calculate2(&n)

	i = 20                                  // here we put an in into the empty interface. I was tring to
	fmt.Printf("The value of i is %v\n", i) // to do this earlier and didn't know who

	i = "some string"
	fmt.Printf("The value of i is %v\n", i)

	i = true
	fmt.Printf("The value of i is %v\n", 1)
}

// Note how our function that takes a parameter with interface shape which
// means we can use any type that fits this "job description" so both
// rectangle and circle work with the calculate function.

// But what about using a type that has more than just what's described in the
// interface? Type assertion allows us to get access to the "underlying type" of
// the instance of an interface. So we put rectangle into calculate as a "shape"
// but under that shape interface we have a concrete type with methods and values
// defined.

// So the type assertion is taking the interface param we want to assert a type
//  to by putting the type we are saying it is in ()s -- s.(rectangle)

// We can also assert that there is a second interface. So we are asserting that
// there's an interface (job description) already created that requires the other
// skill or method we are looking to use

// So am empty interface 'type T interface{}' is an empty job description. We
// accept anyone of any role with any or no methods. So its a way to accept
// anything with having compile issues for not matching the type and dealing
// with it in the actual funcion or program. Any value can be rep'd by an
// empty interface. Somewhere we will see this a lot is with maps. If we need
// to accept values of any type (for example when working iwth JSON) then we
// can use: var myMap map[string]interface{} -- A map with string props and
// empty interface values, or ANY value type we want.

// Type switch statements play the same role with intervaces as they do with
// if/else statements in general. When you need to assert about some types
// and then make a decision based on the results of that assert. Again you may
// need to do this when you don't know what types you are going to recieve but
// want to flexibility in your program.

// Reflection builds on that with another set of tools to find out info on the
// underlying type and value (Type, Kind, Value) by importing "reflect"

// So unlike other languages when you use reflect.TypeOf() you get back an
// Type, not just a string rep of what the type is. This is actual reflection

// Then using the reflect.TypeOf().Kind() gives us the underlying 'kind'. There
// is a differnce between the 'type' and the 'kind' as we see above. If we go
// back to the job description. The type would be the role ('Backend Developer
// or 'Unemployed' or 'Frontend Dev') and the kind would be what kind of type
// or role that is. So we have  types made of structs, maps, other types such
// as  ints, floats, etc. Note that we use Kind() off of the TypeOf returned
// value. Then we go off on a tangent with the many different things we can
// look at from that Type type.

// Lastly we have ValueOf() instead of TypeOf(). which will give us the actual
// values instead of the type. This is interesting that our result is the same
// value that we get when use fmt.Printf with %v. This is because in the 
// implementation of "fmt" its really just using the reflection functions to
// dig into and give the types, values, etc after peeling back the onion.

// Primary thing to remember is that when you pass in an interface (job description)
// there's always an underlying type, the guy we hire always has a 'role' that we
// are taking advantage of for our job, but may have more skills (methods) for us
// to take advantage of. The reflection package allows us to see those.
