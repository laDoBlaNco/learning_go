package main

import "fmt"

/*
	In our continuation of talking types we move from builtin types (variables) to user defined struct types.
		- Remember that if you don't understand your data you don't understand the problem
		- Many times you need to construct or model a special type based on your understanding of the problem
		- "Implicit conversion fo types is the Halloween special of coding. Whoever thought of them deserves
		   to go to their own special hell."- Martin Thompson

*/

// Type provides 2 pieces of information to the compiler, the size and what it represents

// Example represents a type with different fields. Remember  that this is a composite type or based on
// other existing types (builtins)
type example struct {
	flag    bool
	counter int16
	pi      float32
}

// After working with the literal types and structs below, let's take those and create some more named
// types
type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type nancy struct {
	flag    bool
	counter int16
	pi      float32
}

// NOTE: If we were to add these up to see how much contiguous memory will be used we might say it 7 bytes
// That would technically be correct if we just summed up the bool + int16 + float32, But Go wants to be
// mechanically sympathetic so it uses alignments. Remember that we need to remember the machine and Go does
// whatever it can to make the machine's job easier. So aligning has to do with aligning the data to these
// memory word boundaries. So let's say we continue witih our odd continguous memory (instead of 7) let's
// say its 11. If we stayed at 11 bytes on a 64 bit machine, then the
// memory would have 2 8 byte word boundaries. But this would mean that part of the data from our struct would
// then be split overlapping both words. So Go uses alignment (padding) to get the 11 bytes up to some multiple
// of 8 (on a 64bit machine & 4 on a 32bit machine) so that we have evenly split our data into the 2 words.
// this helps the machine because if data in in one word, then it takes 1 operation to get to it. If its split
// across 2 words then we would need 1 op to get part of the data and a 2nd op to get the rest, typically
// with data that would evenly fit into just 1 operation. Its done by looking at the data and making sure it
// aligns with its bytes. So a 2 byte value needs to fall on a 2 byte alignment, it needs to start at 0, 2, 4
// etc. Or a 4 byte alignment needs to start at address 0, 4, 8, etc.So with these structs it needs to go
// through and align everything which causes some gaps that need to be padded.

// Knowing how this works helps us look at a struct and just know its an 8 byte alignment. Typically we won't
// worry about this because our correctness comes first. But if we for some reason need to worry at that level
// and need to do some type of micro optimization for space on the machine, then we could look at out we order
// our fields in the struct to ensure more efficients splits and less padding.

// all of this comes down to knowing the cost of a struct and how much memory will be allocated. Remembering
// that being an engineer means knowing the cost of your decisions.

func main() {

	// Declare a var of type 'example' set to it zero value (0 allacation) defaults
	fmt.Printf("\nStruct Types:\n")
	var e1 example // so this would be a 8 byte allocation

	// display the value
	fmt.Printf("%+v\n", e1)

	// but we could also create a literal and not set the struct to its zero default.
	// Declare a variable of type example and init using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// then display the values
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	// We can also have anony structs.
	// Values can be based on a named type or it can be on a 'literal' type
	// Basically the literal type is one that doesn't have a name (or need a name), unless we give it one later
	// Declare a var of an anony type set to its zero value
	// the difference here from above is that we aren't creating a named type and then a var off that type and then
	// adding values to it. Here we use var and create a literal type (no-named) right off the var name.
	fmt.Printf("\nAnony Structs:\n")
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value
	fmt.Printf("%+v\n", e3)

	// Here we can take it one step further and declare a variable of anony type and init using a
	// struct literal.
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Printf("%+v\n", e4)
	fmt.Println("Flag", e4.flag)
	fmt.Println("Counter", e4.counter)
	fmt.Println("Pi", e4.pi)

	fmt.Printf("\nExplicit vs Implicit Conversion of like Types:\n")
	var b bill
	var n nancy
	// b=n // this would give an error even though these two types are identical and they are compatible (meaning
	// that they have the same memory model). since Go is strongly typed, its not going to do any implicit
	// conversion from one to another, even though in other langs the compiler would say that these are close
	// enough so let's convert bill to nancy or nancy to bill and move on. With Go, they aren't the same type
	// and so they can't be assigned to each other. The only way we can make this happen is doing an EXPLICIT
	// conversion, telling Go that you want to convert one to the other to do the operation
	b = bill(n)

	// A realworld case where this saves the day is the 1000s or millions of bugs caused by perhaps trying
	// to assign an uint to an int. Here implicit conversion would happen in other langs but because this is
	// signed vs unsigned, that conversion causes loss of data and loss of accuracty and integrity. 

	fmt.Println(b, n)

	// One other interesting thing here is that this is only hoppening cuz we are working with 2 NAMED types. If 
	// we were to assigned a struct literal like e4 to our bill var, the compiler would allow it, cuz it can
	// see they are identical and compatible. They aren't two different types becasue one is named and the other
	// is literal (or unnamed). There is no conflict.

	// A real-world example with this situation is with funcs. Funcs in Go are first class values and get 
	// passed around all the time. A func is a literal variable name, a literally typed, or non-name typed
	// variable. This is where we see the flexibility of letting the compiler see if they are identical and
	// compatible and allowing for us to pass them around and assign them. If they were named types,
	// we would have issues. 

	// Recap:
	// 		- Struct types allow us to define our new types ourselves. We can do this with a named struct or
	// 		  literal struct.
	// 		- A type itself helps the compiler know the amount of memory and its representation
	// 		- We talked about padding and alignment and we only order fields based on that when its necesssary
	// 		  for optimization, but that's after we've designed for correctness and readability, and only
	// 		  if we have memory constraints and need to order those fields from largest to smallest
	// 		- Go is very strict type system, so values that are based on name types we need to have like types
	// 		  both sides, but if we are using literals we have a bit more flexibility in using our types.

	// Moving forward the majority of what we do in Go will be based on these user-defined struct types

}
