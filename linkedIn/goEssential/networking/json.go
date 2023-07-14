package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
	"net/http" 
	"bytes"
)

// Request expected are  bank transactions
type Request struct {
	Login  string  `json:"user"`   // since we know what the data fields in our json
	Type   string  `json:"type"`   // look like, we can add 'json tags' to identify
	Amount float64 `json:"amount"` // exactly how our structs maps to the data
}

var data = `
{
	"user":"Scrooge McDuck",
	"type":"deposit",
	"amount":123.4
}`

func main() {
	// create a reader to simulate a file/socket or whatever other reader
	// REMEMBER: READERS ARE THINGS THAT CAN BE READ, THEY DON'T  DO THE READING WE DO
	// WRITERS ARE THINGS THAT CAN BE WRITTEN TO, AGAIN THEY DO THE WRITING, WE DO
	rdr := strings.NewReader(data) // simulate a file/socket

	// Decode request
	dec := json.NewDecoder(rdr)

	var req Request                          // create a req var of our type above Request
	if err := dec.Decode(&req); err != nil { // note we are calling decode on a pointer
		log.Fatalf("error: can't decode - %s", err) // to our struct, not a copy
	}

	fmt.Printf("got: %+v\n", req)

	// We also work from the other end and need to marshal/serialize or encode our data
	// to  json
	prevBalance := 1_000_000.0      //Loaded from a database for example.
	resp := map[string]interface{}{ // string -> whatever
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Let's encode it
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s\n", err)
	}
	
	// APIs - json over http with net/http package
	//GET request
	resp2,err:=http.Get("https://httpbin.org/get")
	if err!=nil{
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp2.Body.Close() 
	
	io.Copy(os.Stdout,resp2.Body)
	
	// We can also use the http post method to make a Post request
	
	// Job is a job description
	type Job struct{
		User string `json:"user"`
		Action string `json:"action"`
		Count int `json:"count"`
	}
	
	// POST request
	job:=&Job{
		User:"Saitama",
		Action:"punch",
		Count:1,
	}
	
	var buf bytes.Buffer // is an in-memory reader/writer
	enc=json.NewEncoder(&buf)
	if err:=enc.Encode(job);err!=nil{
		log.Fatalf("error: can't encode job - %s\n",err)
	}
	
	resp2,err=http.Post("https://httpbin.org/post","application/json",&buf)
	if err!=nil{
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp2.Body.Close() 
	
	io.Copy(os.Stdout,resp2.Body)

}
