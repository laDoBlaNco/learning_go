// https://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go
package main

import (
	"encoding/json"
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"reflect"
)

// ... for quick debugging
var p = fmt.Println

// INTRODUCTION TO INTERFACES:
// We need to know the Go  type system to really explain how to use Go interfaces
// effectively.

// An interface is 3 things:
// 1. it is a set of methods
// 2. but its also a set of types
// 3. and its also it itself a type

// Typically we start to learn interfaces with some contrived example so let's
// do that cuz its realistic?

// A core concept in Go type system is that instead of designing our abstractions
// in terms of the kind of data they can hold, we design them in terms of what
// actions they can execute.

// So we have an animal being any type that can speak. It takes no methods and
// returns a string. So any type that defines a method with this signature is said
// to satisfy the animal interface.
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "???????"
}

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!!!"
}

// Those are the basics now let's going into what the  budding gopher might not
// quite understand

// THE INTERFACE{} TYPE
// interface{} is the empty interface and a source of much confusion. The empty
// interface has no methods list, so therefore all types satisfy this interface
// becasue all types have 0 or more methods. So what new gophers think though is
// that if I pass any type to a func asking for a interface{} type than my arg
// is of 'any' type, but actually its of interface{} type. ALL VALUES HAVE  EXACTLY
// ONE TYPE AT RUNTIME SO OUR STATIC TYPE WOULD BE INTERFACE{}. GO WILL PERFORM
// TYPE CONVERSION AT RUNTIME (IF NECESSARY) AND CONVERT THE VALUE TO AN
// INTERFACE{} VALUE.

// So what is that? An interface value is similar to what I learned about how a
// slice relates to an Array. An interface value is constructed of two words of
// data:
// 1. a pointer to a method table for the value's  underlying type
// 2. a pointer to the actual data being held by that value.
// What we need to understand is that an interface value is two words wide and it
// contains pointers to the underlying data, that's typically enough to avoid many
// common pitfalls around using interfaces. More can be learned from Russ Cox:
// https://href.li/?http://research.swtch.com/interfaces

// so with our example note that we used a slice of Animal (our interface) types. We
// didn't have to say Animal(Dog{}) to turn it into an animal. Every type I put into
// the slice was converted automatically to an Animal type. but at the same time
// each of the items were a different underlying type. So why do I care?

// Well understanding this makes it very easy to understand certain things about
// Go that at times are confusing. Such as why if you have the following sig:
// func PrintAll(vals []interface{}){...} you can't just give it a []string and
// think it'll work. In the sig you are asking for a slice of interface{} types.
// So your slice needs to be field with interface{} values already. As in:
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		p(val)
	}
}

// POINTERS AND INTERFACES:
// Another issue that comes up at times is that when you are using an interface
// you don't really know if its based on a pointer or value receiver. there's no
// guarentee that the underlying type is or isn't a pointer. To see this, let's
// change the Cat type above to have a pointer recvr for its method Speak()
// Note how it breaks our program. It now says:
// "can't use Cat literal (type Cat) as type Animal in array element..."
// This is confusing but in actuality what's happening is that Go is telling us
// our *Cat{} satisfies our interface, but not Cat{}.We put a Cat literal
// in our slice of interface{}. This is because for an interface *Cat{} and Cat{}
// are two different things. When we do somethin similar with Dog, but rather than
// changing the receiver we simpily put a pointer to Dog{} directly in the slice
// note that it works, we didn't have to do anything to the receiver at all. This is
// because a pointer type can access the  methods of its
// associated value type, but not vice versa. O sea *Dog can utilize the Speak
// method but as seen above Cat can't access the  Speak method defined on *Cat
// When you think about it it makes perfect sense because everything in Go is
// passed by value. Every time you call a func, the data you're passing into it
// is copied. In this case of a method with a value receiver, the value is copied
// when calling the method. This makes more sense when you learn to think of
// everything has having a type. Meaning that this method sig:
// func (t T)MyMethod(s string){...} is a function of type func(T,string) :O
// (just as if we had created a associated func to work with the struct rather
// than a method). So method receivers are just passed into the function by value
// just like any other parameter. The receiver is just the 'self' or 'this' which
// in other langs is just a normal function with the first arg being 'self', etc.
// So if our receiver is a value, it passed into our func as a value (copy). Any
// changes to it inside the method won't be seen by the caller because the caller
// is using a completely different Cat value. So again it has no access to *Cat
// original methods. Conversely if we have a method on the Dog type, and we have a
// *Dog pointer, we know exactly which Dog value to use when calling this method,
// because *Dog pointer points to exactly one Dog.

// So after all that the basics are that given an interface that uses a method with
// a pointer receiver, I need to give it a pointer in order to work with it. But
// on the other hand if given an interface that uses a method with a value recv'r
// I can use either a pointer or a value and still have access to its methods.

// THE REAL WORLD: GETTING A PROPER TIMESTAMP OUT OF THE  TWITTER API
// Let's look at somethin real now and not contrived. Timestamps can be rep'd
// in any number of ways in a json doc, since they aren't part of the json spec.
// for example:

// let's start with a string rep of our JSON data
var input = `
{
	"created_at": "Thu May 31 00:00:01 +000 2022"
}
`

func main() {

	// now in our main func we can create a slice of our animals and make them
	// speak as they all satisfy the Animal interface
	// We fixed Cat below using new(), remember that new  allocates memory so in
	// turn what it does is return a pointer to the value's address. We also could've
	// done &Cat{}
	animals := []Animal{&Dog{}, new(Cat), Llama{}, JavaProgrammer{}} // note how we use
	// our interface as the TYPE for our slice
	for _, animal := range animals {
		p(animal.Speak())
	}
	p()

	// So for our interface{} example we note we can't just give it any slice
	names := []string{"stanley", "david", "oscar"}
	// PrintAll(names) // this gives us an error "cannot use names (variable of type
	// of type []string as []interface{} value in argument to Printall)"

	// we need to make a slice of interface{} and push each name into it, thus
	// converting each one individually
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
	p()

	//THE REAL WORLD:
	// our target will be of type map[string]interface{}, which is a pretty
	// generic type that will give us a hashtable whose keys are strings, and
	// whose value can be anything really (o sea of type interface{})
	var val map[string]interface{}

	// idiomatic way of checking for errs on things that don't return vals
	// but give side effects.
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}
	p(val)
	for k, v := range val {
		p(k, reflect.TypeOf(v))
	}

}
