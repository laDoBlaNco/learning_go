// Use the template and follow the directions. You will be writing a web handler that performs
// a mock database call but will timeout based on a context if the call takes too long. You
// will also save state into the context

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// add imports

// Declare a new type named 'key' that is based on an int
type key int

// Declare a constant named 'userIPKey' of type 'key' set to the value of 0
const userIPKey key = 0

// Declare a struct type named 'User' with two 'string' based fields named 'Name' and 'Email'
type User struct {
	Name  string
	Email string
}

func main() {

	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)

}

// routes sets the routes for the web service
func routes() {
	http.HandleFunc("/user", findUser)
}

// Implement the findUser function to leverage the context for both timeouts and state
func findUser(rw http.ResponseWriter, r *http.Request) {

	// create a context that timesout in fifty milliseconds
	duration := 50 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// defer the cancel
	defer cancel()

	// Save the 'r.RemoteAddr' value in the context using the 'userIPKey' as the key, this
	// call should return a new context to replace the current  'ctx' value. The original ctx
	// is the parent context for this new child context.
	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	// create a channel with a buffer size of 1 that works with pointers of type  'User'
	ch := make(chan *User, 1)

	// Now use a goroutine to make the database call. Use the channel to get the user back
	go func() {

		// get the r.RemoteAddr value from the context and log the value you get back
		if ip, ok := ctx.Value(userIPKey).(string); ok {
			log.Println("Start DB for IP", ip)
		}

		// make the database call and return the value back on the channel
		ch <- readDatabase()
		log.Println("DB goroutine terminated")
	}()

	// Wait for the database call to finish or the timeout
	select {
	case u := <-ch:
		// respond with the user if that's what we get back
		sendResponse(rw, u, http.StatusOK)
		log.Println("Sent StatusOK")
		return

	case <-ctx.Done():
		// if we have the ability to cancel the database operation the gr is performing, we
		// should do that now. In this example we can't, so we'll respond with an error
		e := struct{ Error string }{ctx.Err().Error()}
		sendResponse(rw, e, http.StatusRequestTimeout)
		log.Println("Sent StatusRequest Timeout")
		return
	}

}

// let' finish up with our readDatabase function that performs a pretend database call with a
// second of latency
func readDatabase() *User {
	u := User{
		Name:  "Kevin",
		Email: "whitesidekevin@gmail.com",
	}

	// create 100 milliseconds of latency
	time.Sleep(100 * time.Millisecond)

	return &u
}

// sendResponse marshals the provided value into json and return that back to the caller.
func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v) 
}
