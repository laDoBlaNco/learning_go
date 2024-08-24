package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1) // printing empty uninitialized instance X:0 y:0

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2) // print initialize based on order of fields X:1 Y:2

	i3 := Item{
		Y: 10,
		// X:20,
	}
	fmt.Printf("i3: %#v\n", i3) // print initialized based on named fields, if left out
	// then it will be given the zero default value.
	fmt.Println(NewItem(10, 20))  // this one will give us our Item
	fmt.Println(NewItem(10, -20)) // this will give us our error

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)
	// also note that we used a pointer receiver, but  in the function we didn't
	// give it a pointer, we gave it a value. Go compiler is smart enough to know
	// your intention

	p1 := Player{
		Name: "Ladoblanco",
		Item: Item{500, 300}, // we can embed other structs and types in a struct
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X) // we then have access to all original fields
	// of the embeded type
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X) // You can also do this, but its not
	// idiomatic Go
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1) // methods also continue to work with embedded
	// types igual
	//THIS IS NOT INHERITANCE, ITS EMBEDDING:
	//Everytime you call Move, what it will get is an Item. In OOP it'll get the
	// 'this' pointer or 'self' pointer and the type will change depending on the
	// context that its being called in. This is not the case with Go. When you call
	// Move it will only have access to an Item, that's it. So you can't log that
	// you are moving a player with a specific name because it doesn't know anything
	// about that. There is a lot of debate in OOP regrading extending and embedding
	// and reasons why Go has decided to go with embedding and not extending.

	ms := []mover{
		&i1, // This shows us that when we use *pointer receivers we an use them
		&p1, // normally with out worry as we used p1.Move above and it wasn't as
		&i2, // a &pointer. But when we convert it to an interface we need to be more
	} // exact. If you have a value T and you have a value receiver, then its ok to
	// pass the T. But if its a pointer receiver you can't do that. This is why
	// the above code explodes if we try to use i1,p1,i2 values and not &i1,&p1, etc
	// If you have a pointer type *T, you can pass it to value receiver or pointer
	// receiver and both are happy, depending on what you are trying to get with it.
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	// using Go's enum (iota)
	k := Jade
	k1 := Copper
	k2 := Crystal
	k3 := Key(69)
	fmt.Println("k:", k)         // the result of starting with iota+1
	fmt.Println("key:", Key(17)) //
	fmt.Println("k1:", k1)

	// How does Go no how to write time in json?
	// time.Time import json.Marshaler interface
	json.NewEncoder(os.Stdout).Encode(time.Now())

	p69 := Player{
		Name: "ladoblanco",
		Item: Item{500, 500},
		Keys: []Key{k, k1},
	}
	fmt.Println(p69.FoundKey(k2))
	fmt.Println(p69.FoundKey(k))
	fmt.Println(p69.FoundKey(k3))
	fmt.Printf("%#v\n", p69)

	fmt.Println()
	fmt.Println()
	fmt.Println("Sort by distance implementation:")
	lado := NewPlayer("ladoblanco", 69, 69)
	odalis := NewPlayer("odalis", 66, 66)
	kelen := NewPlayer("kelen", 69, 66)

	players := []Player{
		lado,
		odalis,
		kelen,
	}
	fmt.Println(players)
	fmt.Println()
	fmt.Println("Now let's sort em: Who's on top?")
	sort.Sort(ByDistance(players))
	fmt.Println(players)

	fmt.Println()
	fmt.Println("Now from the back?!?!:") 
	sort.Sort(ByDistanceRev(players)) 
	fmt.Println(players) 
}

/*
	Exercise:
	1. Add a 'Keys' field to Player which is a slice of key
	2. Add a 'FoundKey(k Key) error' method to player which will add k to keys if it
	   it's not there
	   --Err if k is not one of the known keys

*/

// In Go we like to see a string rep of everything. So we can create a method
// String(). Due to the interface println will now use this String method and the
// above which was print out 'k: 1' is now print 'k: jade'
// This works because we implement fmt.Stringer interface.
func (k Key) String() string { // method connected to our key type below.
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k) // if we change this %d to a %s or %v we get an
	// error that %s or %v is causing recursion. This is because everytime we want
	// to use a string it'll check our type for  String method and go to our switch
	// and then fall back into the same line asking for %s which will trigger another
	// search for the String method, etc.
}

// Two things to note here: We have const with type. Also iota inside a const group
// is always incremented by 1. Iota in Go is used to represent constant increasing
// sequences. When repeated in a constant, its value gets incremented after each
// specification.
// This is Go's version of enum
const (
	Jade Key = iota + 1 // this starts iota on 1 instead of 0 (1,2,3) vs (0,1,2)
	Copper
	Crystal
	// Jade Key = iota // Same as above. But in const group we can omit the
	// Copper = iota   // repitition by just putting it on the first specification
	// Crystal = iota
	invalidKey // internal and not exported - take advantage of iota (enum) so now
	// we can compare to its iota always knowing its the last in the list.
)

type Key byte

// Rule of thumb: ACCEPT INTERFACES, RETURN TYPES

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Someone asked if they could create a type that is a merge of two different types
// or can handle different types. Normally this can't be done with since 1.18 we
// have generics which makes it possible with constraints. Here we accept an arg
// and return 1 of two types depending on that arg received.

func NewUnion[T int | float64](kind string) T {
	if kind == "int" {
		return 0
	}
	return 0.0
}

// We define interfaces at the point of consumption. In Go interfaces says what we
// NEED not what we provide. AND WE KEEP THEM SMALL: The bigger the interface
// the weaker the abstraction. The more complex you make it, the less flexible it
// becomes. Interface 1 method can except more types, 2 or more methods means less
// types are going to be able to use it.
// We need to remember 2 uses of interfaces:
//  1. To be flexible about what you accept. You decouple interface from the concrete
//     type of thinking. "what we need vs what we provide"
//  2. By implementing a specific interface, it changes the way Go is looking at our
//     type. There are various interfaces in the GoSTL that  we can use and once out
//     type implements them, Go will look differently at them.
//
// Any and all types can implement the empty interface (interface{}) which is the
// smallest interface you can have in Go
//
// The json marshalers/unmarshalers and fmt stringers are the ones we will probably
// use them most in practice, but there are several others and it takes time to learn
// them. Best to learn about them when you need them.
type mover interface {
	Move(x, y int)
	//Move(int,int)  // you can also create the sig like this without names, only types
}

type Player struct {
	Name string
	// X    int // here we have a conflict with X in Item. it still compiles but...
	//now the result of p1.X is 0 and p1.Item.X is still correct. So we have a case
	// of 'shadowing' just like other variables.
	Item // Embeds Item down to any arbitrary level (fields and methods)
	// T // This creates an error an ambiguous selector, if you have two embeded
	// types with the same fields and the parent type doesn't have that field.
	// o sea, if there  was no player.X then when you ask for p1.X, Go doesn't know
	// if you want Item.X or T.X. If you have a Player.X then it'll just use that.
	// Again you get around all the shadowing using the specific p1.T.X or
	// p1.Item.X if needed. As you can see the others continue to work fine even
	// with the T conflict.
	Keys []Key
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey { // the comparison using the firt and last iota
		return fmt.Errorf("invalid key: %#v\n", k)
	} // this is more efficient than my version as I am O(n) this is O(1)
	// Im noting that in Go its idomatic to always take are of the errors first!!
	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil

	// So even with the additional help func (containsKey) his solution was much
	// shorter than mine. Efficiency about the same as he also used a for range loop
	// but in the help func rather than at the top of the function. So really it is
	// more efficient than mine because he checks for errors first, Breaking out first
	// due to an error will always be more efficient than checking for errors last.

	// The above is Go code. Its simple and clear and will do what you need it to do.
	// Yes we'll right more lines of code than in a Python, etc, but its easier to
	// write and maintain and the best of both worlds (low-level systems & dynamic-feel)

	/*
	   MY VERSION

	   var included bool

	   	for _, v := range p.Keys {
	   		if v == k {
	   			included = true
	   		}
	   	}

	   switch {
	   case k == Jade:

	   	switch {
	   	case !included:
	   		p.Keys = append(p.Keys, k)
	   	case included:
	   		break
	   	}

	   case k == Copper:

	   	switch {
	   	case !included:
	   		p.Keys = append(p.Keys, k)
	   	case included:
	   		break
	   	}

	   case k == Crystal:

	   	switch {
	   	case !included:
	   		p.Keys = append(p.Keys, k)
	   	case included:
	   		break
	   	}

	   default:

	   		return fmt.Errorf("%v isn't a legal key", k)
	   	}

	   return nil
	*/
}

func containsKey(keys []Key, k Key) bool { // this is what I was going to do.
	// using a loop to check if the []Key contains my key
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

type T struct {
	X int
	Y int
}

// i is "the reciever" (like 'this' in JS or self in python )
func (i *Item) Move(x, y int) { // must use pointer because again everything is passed
	// by value. if you want ot mutate you have to use pointer receiver.
	i.X = x
	i.Y = y
}

type Item struct {
	X int // these  must be capital so they are exported and others can use outside
	Y int // of the package. If not they would receive a struct Item without visible
	// fields.
}

// We use 'new' functions to initialize our struct types most times and there are
// 4 types

// func NewItem(x,y int) Item {} // Return initialized type
// func NewItem(x,y int)*Item {} // Return a pointer to the initialized type
// func NewItem(x,y int) (Item,error) {} // Return initialized type and an error tag
func NewItem(x, y int) (*Item, error) { // Return pointer to type and error
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
		// note that 'nil' is the zero default for pointers as we are returning here
	}
	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

// 'New' functions in go replace all the constructors, initializers, etc. if its
// just a single type, then just name it New

// zero vs missing value is a delicate question in Go. Is it zero because it was SET
// or because it was the default value?  Normally it doesn't matter, but if you do
// want to know, there's no way to tell in Go.

const (
	maxX = 1000
	maxY = 600
)

// These new functions arent connected to our types in any way. They aren't methods.
// so other than the fact that they return our  new type/struct, they aren't locked
// in

// There are 2 places that a program can allocate memory. It can allocate it on the
// stack which is all the local vars for a function or it can allocate it on the
// heap meaning it is going to outlive the function. So the local vars etc to a
// func will use the stack, but this will be cleared out and able to use by another
// function once the current function ends. But when we return a &pointer we are
// referencing something outliving the function. In C or C++ the person receiving
// that pointer would get an invalid variable. In Go it isn't.

// The Go compiler does "escape analysis" and will allocate items to the heap.
// Go will see that we are returning a pointer to this memory and it will
// allocate memory on the heap to outlive the function.

// Stack: a block of memory allotted to each function. Memory is calculated at
// compile time.

// Heap: a block of memory alloted/demanded by your software at runtime.

// To see what is happening on the heap you can run 'go build -gcflags=-m' and this
// give you the details of the optimizations that the compiler is making. One of
// which here is throwing 'i' to the heap:

/*
$ go build -gcflags=-m
# github.com/ladoblanco/game
./game.go:7:13: inlining call to fmt.Println
./game.go:8:12: inlining call to fmt.Printf
./game.go:11:12: inlining call to fmt.Printf
./game.go:17:12: inlining call to fmt.Printf
./game.go:19:13: inlining call to fmt.Println
./game.go:20:13: inlining call to fmt.Println
./game.go:41:2: moved to heap: i     <-----------------------------------
./game.go:38:25: ... argument does not escape
./game.go:38:26: x escapes to heap
./game.go:38:26: y escapes to heap
./game.go:38:61: maxX escapes to heap
./game.go:38:67: maxY escapes to heap
./game.go:7:13: ... argument does not escape
./game.go:7:13: i1 escapes to heap
./game.go:8:12: ... argument does not escape
./game.go:8:13: i1 escapes to heap
./game.go:11:12: ... argument does not escape
./game.go:11:13: i2 escapes to heap
./game.go:17:12: ... argument does not escape
./game.go:17:13: i3 escapes to heap
./game.go:19:13: ... argument does not escape
./game.go:20:13: ... argument does not escape
*/

// Escape Analysis: Gc compiler does global escape analysis across function and
// package boundaries. It checks a memory that it really needs to be allocated
// to a heap or if it could be managed within the stack itself.

// Function Inlining: Only short and simple functions are inlined. The actual code
// replaces the function call reducing the amount of references and actions needed.

// Escapes to the Heap means that the variable needs to be shared across the function
// stack frames [between main() and Println() for example]

// On returning a pointer the variable is used in the main function will be accessed
// outside of the life of this function. Globally accessed variables must be MOVED to
// the head not just escaped. So this data is moved to the heap memory and now main
// function must access that variable from the heap.

// Heap allocation is usually slower and now the GC is involved, so performance-wise
// its preferible to allocate things to the stack, but at times you can't avoid it.

// In my implementation of the day02 exercise Miki isn't very specific. so I'm
// deducing that distance will be X * Y, I'll also need to create a slice of players
// along with the normal items to be a sort.Interface which will called
// I'm going to measure the distance from 0,0 so for example a player in 9,9 would
// be 9*9 - 0.

func NewPlayer(name string, posX, posY int) Player {
	return Player{
		Name: name,
		Item: Item{posX, posY},
	}
}

// SortByDistance:
// first let's start with my []Player
type ByDistance []Player

func (bd ByDistance) Len() int {
	return len(bd)
}
func (bd ByDistance) Swap(i, j int) {
	bd[i], bd[j] = bd[j], bd[i]
}
func (bd ByDistance) Less(i, j int) bool {
	return bd[i].Item.X*bd[i].Item.Y < bd[j].Item.X*bd[j].Item.Y
}

type ByDistanceRev []Player

func (bd ByDistanceRev) Len() int {
	return len(bd)
}
func (bd ByDistanceRev) Swap(i, j int) {
	bd[i], bd[j] = bd[j], bd[i]
}
func (bd ByDistanceRev) Less(i, j int) bool {
	return bd[i].Item.X*bd[i].Item.Y > bd[j].Item.X*bd[j].Item.Y
}
