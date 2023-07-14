// Go offers built-in support for xml and xml-like formats with the encoding.xml
// package

package main

import (
	"encoding/xml"
	"fmt"
)

// 'Plant' will be mapped to xml .Similarly to the json examples, field tags
// contain directives for the encoder and decoder. Here we use some special features
// of the xml pkg: the 'XMLName' field name dictates the name of the xml element
// representing this struct: id,attr means that the 'id' field is an xml attribute
// rather than a nested element
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string { // making this a stringer
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin,
	)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil", "Dominican Republic"}

	// Emit xml representing our plant; using MarshalIndent to produce a
	// more human-readable output. (note: we are still using 'marshal/unmarshal')
	out, _ := xml.MarshalIndent(coffee, " ", "   ")
	fmt.Println(string(out))
	
	// To add a generic xml header to the output, append it explicitly
	fmt.Println(xml.Header + string(out)) 
	
	// Use Unmarshal to parse a stream of bytes with xml into a data structure. If 
	// the xml is malformed or cannot be mapped into Plant, a descriptive error will
	// be returned
	var p Plant
	if err := xml.Unmarshal(out,&p);err!=nil{
		panic(err) 
	}
	fmt.Printf("%T\n",p)  
	fmt.Println(p) 
	
	tomato := &Plant{Id: 81,Name:"Tomato"} 
	tomato.Origin = []string{"Mexico","California"} 
	
	// The parent>child>plant field tag tells the encoder to nest all plants under
	// <parent><child>...
	type Nesting struct{
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"` 
	}
	
	nesting := &Nesting{} 
	nesting.Plants = []*Plant{coffee,tomato} 
	
	out,_ = xml.MarshalIndent(nesting," ","    ")
	fmt.Println(string(out)) 
	
}

// Next let's take a look at Time
