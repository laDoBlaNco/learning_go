package main

import "fmt"

/*
	Interfaces give programs structure and encourage design by composition. They enable and
	enforce clean divisions between components. The standardization of interfaces can set
	clear and consistent expectations. Decoupling means reducing the dependencies betweeen
	components and the types they use. This leads to correctness, quality, and maintainability

	INTERFACE SEMANTICS - Interfaces allow us to group CONCRETE data together by WHAT THE DATA
	CAN DO. Its about focusing on what data can do and not what the data is. I remember one
	resource (Estroic Tech) described it as a list of job requirements or skills. You can have more
	than what's on the list, but you can't have less. And its focused again on what you can do and
	not what we are or what our title is. Interfaces help us decouple our code from change by
	asking for concrete data based what it can do. It's not limited to one type of data

	We must do our best to understand WHAT data changes are coming and use interfaces to decouple
	our program from those changes. Interfaces should describe BEHAVIOR and not state. They should
	be VERBS and not nouns.

	Generalized interfaces that focus on behavior are best. Interfaces with more than one method
	have more than one reason to change. Interfaces that are based on nouns, tend to be less
	resuable, are more susceptible to change, and defeat the purpose of the interface in the first
	place. Uncertaintly about change is not a license to guess but a directive to STOP AND LEARN
	MORE about our problem. I must distinguish between code that defends against fraud vs code
	that protects against accidents.

	When do we want to use an interface? When:
		- Users of an API need to provide an implementation detail on their own
		- API's have multiple implementations they need to maintain internally
		- Parts of the API that can change have been identified and require decoupling

	When do we not want to use an interface? When:
		- its just for the sake of using interfaces
		- its to generalize an algorithm
			- This makes me think about a problem that needs generics and we try to use interfaces
			  with it. Though I would argue that generics are just interfaces anyways
		- when users can declare their own interfaces
		- when its not clear how the interface makes the code better.


	INTERFACES ARE VALUELESS:
	The first important thing to understand about interfaces is that an interface type, though it
	might like it, doesn't have value. Is a VALUELESS type.

		type reader interface{
			read(b []byte])(int,error) // this is the method  sig / requirement or list
		}

	This type reader is NOT A STRUCT type, but an interface type. Its declaration is not based
	on state, but behavior. That's why we see a method signature as it declares a method set of
	behavior that concrete data must exhibit in order to satisfy/implement the interface. There
	is nothing concrete about interface types, therefore they are VALUELESS

	But then what does this mean:

		var r reader

	Since they are valueless, the construction of a variable  like this one 'r' seems odd, right?
	In our programming model, r doesn't really exist, it's valueless. There is nothing about r
	itself that I can manipulate or transform. This is a critical concept to understand. I'm
	never working with actual 'interface values', only concrete values. An interface does have a
	compiler representation (internal type) which is why we can do the 'construction' above,
	but from our programming model viewpoint, interfaces are VALUELESS.

	IMPLEMENTING INTERFACES:
	So how do we implement an interface?

*/

// Go is a language that is about convention over configuration. When it comes to a concrete
// type implementing an interface, there is no exception. Go devs don't configure and write code
// just to tell the compiler to put on that interface, etc

// In Go, all we have to do is declare a full method-set of behavior defined by an interface (reader)
// to implement that interface. In this case, that is what we've done since the reader interface
// only requires a single act of behavior named 'read'.
type reader interface {
	read(b []byte) (int, error) // here is our behavior / method list
}

// Here we declare a type named 'file' and then a method to that type named 'read'. So we an say
// that "The concrete type (meaning it has a value) now implements the reader interface using
// value semantics"
type file struct {
	name string
}

func (file) read(b []byte) (int, error) { // here's our sig, which means we implement the above interface
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// Let's create another concrete type 'pipe'.
type pipe struct {
	name string
}

// Then we declare a method named read. Because of this we can now say, "The concrete type pipe
// now implements the reader interface using value semantics". Here we now have 2 concrete types
// implementing our reader interface. Two concrete types with their unique implementations. One
// type is reading file systems and other networks. This means its POLYMORPHIC
func (pipe) read(b []byte) (int, error) {
	s := `{name: "Ladoblanco", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// Let's do something with these two  before we get into POLYMORPHISM
func main() {

	// create two values one of type file and one of type pipe
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call each retrieve function for each concrete type
	retrieveFile(f)
	retrievePipe(p) 
}

// Now what our retrieveFile function does is read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data[:len]))
	return nil
}

// and a separate retrievePipe to read from pipes and process the data
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data[:len]))
	return nil
}

// before moving on, the key idea here is that we've created an interface but what can we do with
// them? Why is it focused on behaviors and not state? Well look here above we have two different
// concrete types that 'implement' 'reader' since they have a read method. But we had to create
// two separate functions of retrieve. With interfaces we look at ways that we can DECOUPLE this
// change in what we are retrieving from, out of our code. So we can use one 'retrieve' for 
// both pipes and files. 
// Enter POLYMORPHISM
