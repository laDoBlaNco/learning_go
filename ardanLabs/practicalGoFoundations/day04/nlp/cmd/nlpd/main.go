package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "expvar" 

	"github.com/gorilla/mux"

	"github.com/ladoblanco/nlp"
	"github.com/ladoblanco/nlp/stemmer"
)

func main() {
	// create server (dependency injection)
	logger := log.New(log.Writer(), "nlp ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger, // dependency injection
	}
	// routing
	// /health is an exact match
	// /health/ is aa prefix match
	// http.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/tokenize", tokenizeHandler)

	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// run server
	addr := ":8080"
	s.logger.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* before gorilla
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	*/
	// Usually handlers have 3 steps
	// 1. Get, convert and validate the data
	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1_000_000)
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return // NOTE: very important you must return. http.Error isn't an exception and won't
		// exit from the handler by itself
	}

	if len(data) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
		return
	}

	text := string(data)

	// 2. Work
	tokens := nlp.Tokenize(text)

	// 3. Emit output
	resp := map[string]any{
		"tokens": tokens,
	}
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't Encode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Exercise: Write a tokenizeHandler that will read the text from the r.Body
// and return JSON in the format "{"tokens": ["who","on","first"]}"

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// handlers get two things. What you are writing into (the response) and the request where
	// query, etc
	// TODO: Run a health check
	fmt.Fprintln(w, "OK gopher y now?!") // w uses the io.Writer interface so everything that uses that can
	// use our w
}

// We can test the basic routing and receiving on the commandline but if we want to do anything
// fancy we can use something like 'postman' or 'insomnia' to test posts, etc.
