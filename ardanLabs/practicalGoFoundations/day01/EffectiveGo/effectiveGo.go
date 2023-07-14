package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	fmt.Println("-----------DEFER------------")
	b()
	/*
		for i := 0; i < 5; i++ {
			defer fmt.Printf("%d ", i) // print off  the deferred stack FILO, notice that
			// this is getting printed after everything else in my main func hahahaha
		}
		commenting this out as its defere to the end of main and confuses me with the
		output
	*/

	fmt.Println("\n", "-----------DATA------------")
	fmt.Println("-----------Arrays - Summing an array------------")
	fmt.Println(Sum(&[3]float64{7.0, 8.5, 9.1})) // I like the use of naked return on
	// the sum func, but not sure why we are using a pointer

	// Some slice examples - slices of slices
	fmt.Println("\n", "---------Some slice examples - slices of slice---------")
	text := LinesOfText{ // so LinesOfText is a slice filled with slices
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Println(text)

	// Appending slices - cap and len are still legal on the nil slice
	fmt.Println("\n", "--------Some slice examples - Nil slices and myAppend---------")
	mySlice := new([]int)
	fmt.Printf("Both len() and cap() are legal with nil slices:%#v - %#v\n", len(*mySlice), cap(*mySlice))

	anotherSlice := []int{1, 2, 3}
	anotherSlice = AppendInt(anotherSlice, []int{4, 5, 6})
	fmt.Printf("%#v\n", anotherSlice)
	// remember you can't get a nil slice from make, but with new you get back a pointer. with
	//dereferencing, len and cap work as they should.

	// MAPS - cap and len are still legal on the nil slice
	// Fetching map values
	fmt.Println("\n", "--------Some Map Examples---------")
	fmt.Println(timeZone["EST"])
	fmt.Println(timeZone["DLS"])

	// Using comma ok idiom to determine if something is actually in the map or set to 0
	seconds, ok = timeZone["DLS"]
	fmt.Println(seconds, ok)      // 0,false (false meaning its not in the map)
	_, present := timeZone["DLS"] // using comma ok just to see if something is present
	fmt.Println(present)
	// same as above but with my offset function from below
	fmt.Println(offset("DLS"))
	// deleting an entry
	fmt.Println(timeZone)
	delete(timeZone, "PST")
	fmt.Println(timeZone)
	timeZone["PST"] = -8 * 60 * 60 // this works because even though the entry doesn't exist
	// I'm working with an actual map. If I was working with a nil map then I'd get
	// an error until I initialize (make) and actual map
	fmt.Println(timeZone)
	// working with a nil map example
	fmt.Println(timeZone2)        // uninitialized map
	fmt.Println(timeZone2["PST"]) // works to show me zero
	// timeZone2["PST"]=-8*60*60 // this gives a panic, assignment nil map

	// Some Printing examples. Using Printf, Fprintf and Sprintf. All of these also have
	// their respective non format string versions (Println and Print).
	// Fprint family needs as its first arg any object that implements io.Writer interface
	// I can think of its as F(ormat)print, S(ring)print, P(rint)rint (o sea normal print)
	// F - send formated string to io.Writer, S - return a formated to string (not printed), P - print to stdout normal.
	fmt.Println("\n", "--------------Printing Examples - with Printf, Fprint, Println, and Sprint-------------")
	fmt.Printf("Hello %d\n", 23)              // note the need for the '\n'
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n") // here we see the addition of the io.Writer (os.Stdout)
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23)) // here we return the string from Sprint to Println which takes care of the \n

	// Something I found interesting already is that %d doesn't take flags like %f does. Go determines based on the type
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	// %v is for the default value and can print any value, even slices, arrays, structs, and maps. %v = value o sea Go's
	// value
	fmt.Printf("%v and %#v\n", timeZone, timeZone) // from the Timezone example, print maps

	// print stuct examples:
	fmt.Println("\n", "--------------Printing Examples - structs-------------")
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v - Value normal\n", t) // this is affected by the custome String() method below
	fmt.Printf("%+v - Value con los fields agregado\n", t)
	fmt.Printf("%#v - Alternative value con los fields agregados\n", t)
	fmt.Printf("%#v - same as above\n", timeZone)

	//%q will put quotes, %x will convert to hex print, and %T will print the type
	fmt.Printf("%T\n", timeZone)

	// As I learned with interfaces, if we want to control the default printing of customer types we add a String() method
	// as I did with T below
	fmt.Printf("%v\n", t)

	fmt.Println("\n", "--------------Using my Min func to demo ...args--------------")
	fmt.Println(Min(69, 24, 32, 6, 169))
	fmt.Println(Min([]int{69, 24, 32, 6, 169}...)) //this won't compile using []int unless its
	// used with ... on the end
	fmt.Println("\n", "--------------Using the built-in append---------------")
	aSlice := []int{1, 2, 3}
	aSlice = append(aSlice, 4, 5, 6)
	fmt.Println(aSlice)
	// now with a slice arg
	aSlice2 := []int{7, 8, 9}
	aSlice = append(aSlice, aSlice2...) // only works with anotherSlice... aSlice2 is not
	// of type int which is what our func sig is asking for.
	fmt.Println(aSlice)

	fmt.Println("\n", "-------------------Using our new ByteSlice method---------------")
	//The first returns and the other are inline as we used pointer receivers
	var myByteSlice ByteSlice
	fmt.Println(myByteSlice.Append([]byte("kevin")))
	myByteSlice.Append2([]byte("ladoblanco")) // this onei s changed inline
	fmt.Println(myByteSlice)
	fmt.Fprintf(&myByteSlice, " This is a test %d methods\n", 69)
	fmt.Println(string(myByteSlice)) // here we used our new Writer method with Fprintf
	// THE RULE: With pointer vs values for receivers, value methods can be invoked
	// on pointers and values, but pointer methods can only be invoked on pointers.
 
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

// Array example
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

/*
Our custom Append function, so we understand what the builtin function does and
how it works with slices.

The length of a slice may be changed as long as it still fits within the limits of
the underlying array; just assign it to a slice of itself. The capacity of a slice,
accessible by the built-in function cap, reports the maximum length the slice may
assume. Here is a function to append data to a slice. If the data exceeds the
capacity, the slice is reallocated. The resulting slice is returned. The function
uses the fact that len and cap are legal when applied to the nil slice, and
return 0.
*/
func AppendInt(slice, data []int) []int {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth
		newSlice := make([]int, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)] // here we are at the point where we have the space needed, whether it be
	// that it was already there or that we reallocated. We then then reslice slice to
	// take advantage of the new space and below copy the data onto the end of the new slice starting
	// right after the original data '[l:]'

	// We must return the slice because even though we've talked about slices being pointers
	// the slice itself (holding the pointer,length,capacity) is passed by value.
	copy(slice[l:], data)
	return slice
}

// Slices - multi-dimensional (array of array or slices of slice)
type Transform [3][3]float64 //array of array every set of brackets is a dimension
type LinesOfText [][]byte    //slice of slice same as above

// Note: slices can't be used as map keys because they don't have equality defined
// in them. Hold refs to an underlying structure

// Maps - also hold ref to an underlying structure, so like slices, if you pass to
// a function that changes contents, the changes will be visible to the caller

// composite literal map[string]int or timeZone->timeAdjustment
var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}
var seconds int
var ok bool

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

var timeZone2 map[string]int

// When printing structs we can use the '+' modifier which will annotate the fields, and any value we an use
// '#' for the full Go syntax
type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

// when creating String methods be careful not to get stuck in recursion by
// using %s to convert the string. The first will recur, the second and third
// will not
type MyString string

// func (m MyString) String() string {
// return fmt.Sprintf("MyString=%s", m) // This will compile with a warning that
// // 'format %s with arg m causes recursive...'
// }

func (m MyString) String2() string {
	return fmt.Sprintf("MyString=%s", string(m)) // here by using string() to convert
	// before the sub, we avoid conversion recursion because string() doesn't use
	// the string interface
}

func (m MyString) String3() string {
	return fmt.Sprintf("MyString=%v", m) // here we use the natural value with %v and
	// just let Sprintf return the entire string.
}

// all of the above compile, which is why you can easily fall into the trap of
// conversion recursion

// To finalize printing we look at ... on println. when I see ... before a type
// in a func sig it means that it'll take however many args and treat as a slice
// so ...T would be as many args as needed and would be converted to []T in the func
// If we are then using a slice as an arg we would say arg... to tell Go to use that
// arg as a list of names and not the entire slice together. For example

/*
func Println(v ...interface{}){
	std.Output(2,fmt.Sprintln(v...)) // Output takes parameters (int,string)
}
Note that we use ...interface{} as the func sig type and then in the func 'v' is
converted to a string so we need to use (v...) to then pass that to Sprintln so it
passes it as a list of args again rather than the slice it was converted into
*/

// Another example of using ... in the min function below. We can either use any
// amount of args, or pass it a slice...
func Min(a ...int) int {
	min := int(^uint(0) >> 1) //largest int, I'll need to figure out this bit stuff later
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// So knowing about ... we can see the actual append implementation sig of
// func append(slice []T,elements ...T)[]T
// builtin append takes your elements and appends them to the end of your slice
// and it doesn't care how many you put as long as its either comma separated args
// or []T... marked to be used as is an not converted to a []T on the inside.

// Initialization is more powerful in Go than in C/C++. Complex structures can be
// built during initialization.

// Constants are created at compile time no matter where they are defined. They can
// only be numbers, runes, strings or booleans. Exprssions that define them must be
// 'constant expressions' o sea evaluatable by the compiler like 1<<3. Something like
// math.Sin(math.Pi/4) can only be evaluated at run-time.

// Enumerated constants are created with 'iota' which is the Go enumerator. It can
// be implicitally repeated since its part of the expression
type ByteSize float64

const (
	_           = iota             // ignore the first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // so here the type and enumerator need not be set
	MB                             // explicitly in this expression
	GB
	TB
	PB
	EB
	ZB
	YB
)

// then we can add the String() method and create the formatting to be printed
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// Also note that we are good with the recursion because we are using %f and not %s

// METHODS - Pointers vs Values
// A method can be defined on any named type (except a poitner or an interface).
// The receiver doesn't have to a be struct. Below let's add our Append function
// to a []byte as a method and make it better with each go.
type ByteSlice []byte // first we created a named  type based on the builtin

func (slice ByteSlice) Append(data []byte) []byte { // same body as before
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

// but we can remove the return by using a pointer receiver so it works inline
func (p *ByteSlice) Append2(data []byte) {
	slice := *p //here we assign dereferenced p to slice so as not change the code
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice // and from slice back to *p
}

// Here we can turn it into a writer by fulfilling the interface
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(data), nil
}

/*
	IMPORTANT NOTE THAT WAS VERY CLOUDY IN THE BEGINNING:
	This rule arises because pointer methods can modify the receiver; invoking
	them on a value would cause the method to receive a copy of the value, so any
	modifications would be discarded. The language therefore disallows this mistake.
	There is a handy exception, though. When the value is addressable, the language
	takes care of the common case of invoking a pointer method on a value by
	inserting the address operator automatically. In our example, the variable b is
	addressable, so we can call its Write method with just b.Write. The compiler
	will rewrite that to (&b).Write for us.

*/

// Interfaces provide a way in Go to specify behavior. As I learned before like a
// job description. If something can do 'this' it can be used 'here'. We learned about
// this above with 'String() and Write()' methods. Another good example is the sort
// interface which has a list of len(), less(i,j int)bool, and swap(i,j int) methods
// along with a customer formatter with String(). For example:
type Sequence []int // we have a type Sequence which is a []int

// Methods req'd by sort.Interface are:
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// So with the 3 methods implemented above we now fulfill the reqs for a sorter
// o mejor dicho using sort.Interface
// Adding Copy to return a copy of the Sequence
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Now we do some custom formattering by creating a String() method to work with
// of our our print routines.
func (s Sequence) String() string {
	s = s.Copy() // make a copy so as not to overwrite the arg
	sort.Sort(s) // sort that copy
	str := "["
	for i, elem := range s { // Loop is O(N2); we'll fix that in next example
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// Interface conversions and type assertions
// A type switch is considered a conversion operation becasue it is in a way 
// converting an interface to the type of each case in the switch:
// type switch uses switch var := value.(type){} -- the value.(type) only works
// inside of a type switch and with an interface.

// if we only have 1 we can simply do a type assertion 'str := value.(string)' 
// when then use this with comma string to ensure that str is aa string.
