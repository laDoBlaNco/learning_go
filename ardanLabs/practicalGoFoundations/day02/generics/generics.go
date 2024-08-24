package main

import(
	"fmt"
)

func main(){
	// Initialize a map for the integer values
	ints:=map[string]int64{
		"first":34,
		"second":12,
	}
	
	// Initialize a map for the float values
	floats:=map[string]float64{
		"first":35.98,
		"second":26.99,
	}
	
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats)) 
		
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string,int64](ints), 
		SumIntsOrFloats[string,float64](floats))
		// first time I see this used, but both work and gopls is giving me a 
		// warning that I don't need the type args when calling. Go infers them
		// at compile time based on your actual args.
		
	// here without the type args as they are inferred as mentioned above:
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints), 
		SumIntsOrFloats(floats))
		
	// here we use our Number constraint interface
	fmt.Printf("Generic Sums with Constraint Interface: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats)) 
	
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64)int64{
	var s int64 // using that zero default to start
	for _,v:=range m{
		s+=v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64)float64{
	var s float64 // using that zero default...and lots of repeated code. 
	for _,v:=range m{
		s+=v
	}
	return s
}
// Note how the only difference in these functions is the name and the type used...
// Prime candidate for using generics.

// So let's add a generic function that can take either maps as args. So basically
// replacing both of our current functions with a single func. 

// This is done using Type parameters in addition to the normal func params. This is
// what makes the function generic, enabling it to work with different types. Each
// type param will have a CONSTRAINT taht acts as a kind of meta-type for the 
// param. It will specify the  permissable type args that can be used.

// At compile time the type parameter stands for a single type though the type
// parameter's constraint reps a set of types. So if the type used isn't allowed
// by the constraint, then it won't compile.

// here we use a constraint that only allows Ints and Floats

// SumIntsOrFloats - sums the values of map m. It supports both int64 & float64
func SumIntsOrFloats[K comparable,V int64|float64](m map[K]V)V{
	var s V // I've typically only seen generics with T but this is map K,V
	for _,v:=range m{
		s+=v
	}
	return s
}

// So here we see the type constraint in []s before the normal params. We use K & V
// as our generic args and the func param uses those to identify the types instead
// of map[string]int64/Float64.
// In that constraint specify that K must be comparable and V can be int64 or float64
// The comparable constraint is pre-declared by Go and allows any type that can be
// used with == and !=. so map keys must be comparable and ensures that we are using
// a legal type for the keys of our map. '|' specifies a union of two types o sea
// we can use int64 or float64 for V. 
// Finally we specify the "normal" func param using our generic identifiers:
// map[K]V (NOTE: this is a valid map type cuz K is comparable. If we hadn't declared
// K comparable, the compiler would reject the ref to map[K]V)

// We an also put our contraints into their own interfaces which allows us to reuse
// them in other places. this is not something I knew about but definitely falls in
// line with the magic of Go. 

// We declare a type constraint as an interface. It will allow any type implementing
// the interface. So we can declare a type constraint with three methods, then use it
// with a type parameter in a generic function, type args used to call the function
// must have all 3 of these methods for it to be of that interface.

type Number interface{
	int64|float64 // here we are saying the interface can be either of these types
}

// Now let's do our func again but with our new type (interface)
// SumNumbers sums the values of map m. It supports both ints and floats as values
func SumNumbers[K comparable,V Number](m map[K]V)V{
	var s V
	for _,v:=range m{
		s+=v
	}
	return s
}
