package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// start by creating your handlers - functions with 2 values.
// the response you Write to and the Request that is the incoming  request

// healthHandler returns a server health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "OK")
}

// Here we are creating both sides to our mathHandler example. The request and the resp
// MathRequest - request of math operation
type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// MathResponse - response to MathRequest
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

// When working with handlers we mainly follow 3 steps:

// mathHandler returns result of calculation
func mathHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Decode & Validate
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		log.Printf("error: bad JSON: %s", err)
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("error: bad operator: %q", req.Op)
		http.Error(w, "uknown operator", http.StatusBadRequest)
		return
	}

	// Step 2: Do the work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unknown operation: %s", req.Op)
	}

	// Step 3: Encode my result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		//Can't return error to client here so we just log it
		log.Printf("can't encode %v - %s", resp, err)
	}

}

func main() {
	// start by routing the endpoints
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	} // This is a plan http server. We can also use  http.ListenAndServeTLS to serve
	// over TLS or https
}
