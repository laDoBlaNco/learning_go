package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for json encoding and decoding (also known as marshaling
// and unmarshaling), including to and from built-in and custom data types.

// We'll use these two structs to demonstrate encoding and decoding of custom
// types below.
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in json. Fields must start with
// capital letters to be exported
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// First we'll look at encoding basic data types to json strings. Here are
	// some examples for atomic values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// And here are some examples for slices and maps, which encode to json
	// arrays and objects as you'd expect
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The json pkg can automatically encode your custom data types. It will only
	// include exported fields in the encoded output and will by default use the
	// names as the json keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can also use tags on struct fields declarations to customize the encoded
	// json key names. check the definition of response2 above to see an example of
	// such tags
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//========================================================
	// Now let's look at decoding/unmarshaling json data into Go values. Here's
	// an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We first need to provide a variable where the json package can put the
	// decoded data. This map[string]interface{} will hold a map of strings to
	// any arbitrary data type which is basically what all json will be in the end
	var dat map[string]interface{}

	// first let's look at the actual decoding and checking for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Now that we have the data decoded, we need to be able to use the values
	// in the decoded map, so we'll need to convert them to their appropriate type.
	// For example here we convert the value in num to the expected float64 type
	// with a type assertion

	// Before we go into this example I need to remind myself what type assertion
	// actually returns. I though bool, but apparently its actually a conversion
	var x interface{} = "foo"
	var s string = x.(string)
	fmt.Println(s)
	// So a type assertion doesn't really convert an interface to another data type
	// but it provides access to an interface's concrete value, which is typically
	// what you want. The type assertion x.(T) asserts that the concrete value of x
	// is of type T and that x is not nil.
	// 		If T is not an interface, it asserts that the dynamic type of x is identical
	// 		to T.
	// 		If T is an interface, it asserts that the dynamic type of x implements T
	s, ok := x.(string) // we can also do comma ok here
	fmt.Println(s, ok)

	n, ok := x.(int)
	fmt.Println(n, ok) // note that if we use comma ok Go won't panic but it'll just
	// give us the zero default along with the bool. If we don't use comma ok and
	// assert incorrectly then Go will panic.
	// n = x.(int) // panic: interface conversion: interface {} is string, not int

	// Now back to our regularly scheduled examples:
	num := dat["num"].(float64) // asserting that the map value at key "num" is a float64
	fmt.Println(num)

	// Accessing nested data requires a series of conversions/assertions, because you
	// must convert them top to bottom in the nest, so to speak
	strs := dat["strs"].([]interface{}) // first we assert that the value at map key "strs"
	// is a slice of 'any' (interface{}).
	str1 := strs[0].(string) // then we assert that the first value of that slice is string
	fmt.Println(str1)

	// we can also decode/unmarshal into custom data types. This has the advantages
	// of adding additional type-safety to our programs and eliminating the need
	// for type assertions when accessing the decoded data. so technically its a
	// different and better way to do the above.
	// This is technically the same as the above. A json object with strings, nums,
	// and a slice
	str := `{"page":1,"fruits":["apple","peach","pear"]}`
	res := response2{}                // we instantiate a custom our custom struct
	json.Unmarshal([]byte(str), &res) // like when we Unmarshaled into our generic
	// map in the other example.
	fmt.Println(res)
	fmt.Printf("%+v\n", res) // good example of the dif between '#' and '+' verbs
	fmt.Printf("%#v\n", res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings as intermediates
	// between the data and json representation on standard out. We can also stream
	// json encodings directly to the os.Writers like os.Stdout or even http
	// response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}

// We've covered the basics of json in Go here, but we can later check out:
// https://go.dev/blog/json
// https://pkg.go.dev/encoding/json ... for more

// Now let's check out XML
