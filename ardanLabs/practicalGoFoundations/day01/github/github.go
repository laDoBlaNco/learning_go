package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"context"
	"time"  
)

func main() {
	resp, err := http.Get("https://api.github.com/users/ladoblanco")
	// this is a good example of error handling, returning errors from funcs instead
	// of using exceptions like the rest.
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.Printf("error: %s",err) //others like to do it this way, but Fatalf is easier
			os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// note that some use a map to get these values, but we don't do that in Go for http
	// first cuz http is case-insensitive and it can have more than one header, maps can't
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil { //using _ for unused var. and it can't be read from
	// 	log.Fatalf("error: can't copy - %s", err)
	// }
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	// fmt.Println(r)
	fmt.Printf("%#v\n", r) // again doing this with %#v gives you more info to look at, overall
	// type struct and its values - main.Reply{Name:"l@D0Bl@Nc0", Public_Repos:1}

	// Testing with my own function from above - and now adding in context
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	fmt.Println(gethubInfo(ctx,"ladoblano")) // missing 'c' so we can see the error
	name, repos, _ := gethubInfo(ctx,"ladoblanco") 
	fmt.Printf("The gethub info for me is - Name: %s and Number of Repos: %d\n", name, repos) 

} 

// githubInfo returns name and number of public repos for login
func gethubInfo(ctx context.Context,login string) (string, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login)) //pathescape ensures that if our login has spaces or anything that won't be read correctly, it still works
	// resp, err := http.Get(url)
	req,err :=http.NewRequestWithContext(ctx,http.MethodGet,url,nil) 
	if err != nil {
		return "", 0, err // instead of just panicing and print idiomatic go is to return 3
	} //...still. Normally those are the default empty vals and the error to handle
	resp,err := http.DefaultClient.Do(req)
	if err!=nil{
		return "",0,err
	}
	
	if resp.StatusCode != http.StatusOK {
		return "", -1, fmt.Errorf("%#v - %s", url, resp.Status) // another example is returning
	} // ...a -1 for the int or building the error message with fmt.Errorf()

	// Also going to use an anonymous strut instead of the one below since its a small prog
	var r struct {
		Name     string
		NumRepos int `json:"Public_Repos"`
	} // Notice we aren't obligated to add the second {} with the values since we are using it in the function.
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	} // One last thing, Json doesn't care if our struct has more or less fields and neither does Go

	return r.Name, r.NumRepos, nil
}

// Pointers are only used for referencing in Go. You can use a packaged named rightly "unsafe"
// in order to do the same things you can do in c/c++, but for normal Go its just for ref
// you want to use pointers for example to 'r' above, because if you send the actual struct
// will simply make a copy and you want be doing anything with the actual struct type.
// Therefore we use pointers to r (&r memory address of r)

// to help Json with hints we are creating our own type
type Reply struct {
	Name string
	// Public_Repos int
	NumRepos int `json:"Public_Repos"` // the json will put the info under our value
	// uninitialized values in go have a default zero value. If you see that
} // it meansit was just not found. But with a field tag we can name the val something and tag it something else.

//serialization is the act of taking a data struture  in our programming language and converting it to
// a seq of bytes. Its also call marshaling, deserialization or unmarshaling on the other side.
// I've heard all of these terms. We do this becasue you can only move data as a series of
// bytes. So there needs to be a translation. if you aren't transferring data to another
// machine, then you shouldn't really be marshaling.

//JSON, XML, CSV, etc all are serializtion strutures or formats.
// So how do we go about converting JSON to Go
/*
true/false <-> true/false
string <-> string - JSON also has utf-8 as does Go
null <-> nil
number <-> float64,float32,int8 int16,int32,int64,int,unit8, ... so we need to give hints here
array <-> []T - arrays in json or js are dynamic, but in Go you have to use
			    either []any or  before generics []interface{} (any is 1.18+)
object <-> map[string]any or a structconst

1. When working with marshaling you always have to know formats on both sides of your marshal

API for JSON encoding/decoding in Go
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
G0 -> []byte -> JSON: json.Marshal
*/
