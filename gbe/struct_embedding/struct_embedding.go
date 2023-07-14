package main

import(
	"fmt"
	_"log"
	_"os"
	_"os/exec"
)

// ... for quick debugging
var p = fmt.Println

// Go supports embedding of structs and interfaces to express a more seamless
// composition type. This isn't to be confused with //go:embed which is a go
// directive introduced in Go 1.16+ to embed files and folders into the app
// binary..


type base struct{
	num int
}
func (b base)describe()string{
	return fmt.Sprintf("base with num=%v",b.num) 
}

// a container embeds a base. an embedding looks like a field without a name
type container struct{
	base
	str string
}

func main(){
	// When creating structs with literals, we have to initialize the embedding
	// explicitly; here the embedded type serves as the field name
	co:=container{
		base:base{
			num:1,
		},
		str:"some name",
	}
	// We can access the base's fields directly on co, e.g. co.num
	fmt.Printf("co={num:%v, str:%v}\n",co.num,co.str) 
	p()
	
	// Alternatively, we can spell out the full path using the embedded type
	// name
	p("also num:",co.base.num) 
	p() 
	
	// Since container embeds base, the methods of base also become methods of
	// a container. Here we invoke a method that was embedded from base directly
	// on co
	p("describe:",co.describe()) 
	p() 
	
	// I don't know that i've ever seen a type/interface created within a 
	// function before
	type describer interface{
		describe() string
	}
	
	// Embedding structs with methods may be used to bestow interface implementations
	// onto other structs. Here we see that a container now satisfies the 
	// describer interface because it embeds base.
	var d describer = co 
	p("another describer:",d.describe()) 
	
	
}
