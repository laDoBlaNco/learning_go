package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Q: what is the most common word (ignoring case) in sherlock.txt?
// Word frequency

func main() {
	file, err := os.Open("sherlock.txt")
	// fmt.Printf("%#v\n",file)
	// goland: freq/sherlock.txt
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	// fmt.Println(wordFrequency(file))

	w, err := mostCommon(file, 5)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	fmt.Println(w)

	/*
	   	path := "c:\to\new\report.csv" // note the special chars cause we are using \
	   	fmt.Println(path)              // Go see a tab, a newline, and a carriage return
	   	path = `c:\to\new\report.csv`
	   	fmt.Println(path) // now its prints correctly.

	   	// multi-line strings: (as normal the tabs are seen as well so needs to be off
	   	// the left margin to print correctly)
	   var request = `GET /ip HTTP/1.1
	   Host: httpbin.org
	   Connection: Close

	   `
	   	fmt.Println(request)
	*/

	// mapDemo()

}

//Need to change this to:
// func mostCommon(r io.Reader,n int)([]string,error)
// We are going to  need to sort by wc and then pull n from the top of the
// slice.

// Done deal. I ended up finding out how to sort a slice using sort.Slice and
// sort.SliceStable (difference being how it treats items with matching sort vals)
// This sort function takes an anony Less func and that Less func is based on
// the map that my slice of keys comes from. So in the end we sort the slice of
// keys and  take the first n and done. I also printed out the top words and their
// counts - DONE!

// First we create the type for the sort.Interface:
type ByCount map[string]int

func (bc ByCount) Len() int              { return len(bc) }
func (bc ByCount) Swap(i, j string)      { bc[i], bc[j] = bc[j], bc[i] }
func (bc ByCount) Less(i, j string) bool { return bc[i] > bc[j] } // reversed, Greatest to least

func mostCommon(r io.Reader, n int) ([]string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(freqs))
	for k := range freqs {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freqs[keys[i]] > freqs[keys[j]]
	})
	fmt.Printf("These are the type %v words:\n", n)
	keys = keys[:n]
	for _, w := range keys {
		fmt.Println(w, freqs[w])
	}
	return keys, nil
}

// Now we need to get the maximum - Again its just Go code. Nothing fancy and when
// you come back in two months you'll have no issues in knowing what's going on.
func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map") // note the return of "", zero default for a string
	}
	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, nil
}

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`) // always document your regexp
// an example of the input and what the result should be

// Note: When we declare a package level var, we can use a function call which
// means that it is running before main. Init is another  function that can execute
// before main. Normally we don't want to but creating a regex is fine.

// Also with regexp you can't return err. you can only panic. That's why its 'MustCompile'
// It  must give you what you want or panics since you are only doing once in the
// beginning. The same is true for other templating packages such as HTML.

// Finally note how we use ``s  which we can considered an raw string (escaped)
// You also use raw strings for multiline strings
// We can see the differences above in our print examples. The reason we need to know
// this now is because regexp have a lot of \s meaning we need to use raw strings
// in order to now worry about escaping them when needed.

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++ // in go this just works since if the key
			// doesn't exist, we still get a default zero.
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}

// Iterating files one line at a time. Files can get very big, gigs of data and you
// don't want to read the entire file at a time. there's a lot you can do processing
// 1 line at a time. We do it with bufio.NewScanner which gets an io.Reader and then
// it gives us a line at a time. its being discussed now about adding a generic
// iteration protocol to Go, but for now this is the closest we have and the
// suggestion is pretty close to what we are seeing here.

// So the above is the mechanism we use to get one line at a time. Once we have this
// we need the words. To do this correctly as we do in NLP, its tough. We are going
// to use regex here.

// So now that we have the mechanism of going line by line and getting the words
// we need to now count the appearance of each word and we'll be using maps - as of
// right now the implementation of Go maps is like a hashmap in that the order isn't
// something you can count on. This might change in the future but for now its an
// unordered hashmap.

func mapDemo() {
	var stocks map[string]float64 // sym -> price
	//Go maps are like dicts or hashtables. 2 types, key/value. Its good to always
	// document your maps cuz you know the  types but may not know what those types
	// are going to be used for in the map.
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)
	// this works even though we didn't initialize the key, if the key doesn't exist
	// we don't get an error but we get the zero default value. So again how do you
	// determine if its missing vs  an actual value of zero? - comma,ok  ;). So
	// in this case getting just 1 value, regardless if it exists in the map we will
	// get a value. But if we ask for 2 vals, then we get val and a bool. true if it
	// exists and false if it doesn't.
	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	//stocks = make(map[string]float64) // one way for initializing the map
	stocks = map[string]float64{ // or we can use a map literal
		sym:    137.74,
		"AAPL": 172.35,
	}
	stocks[sym] = 136.73              // We can work with a nil map until we try to assign vals
	if price, ok := stocks[sym]; ok { // once we want to assign vals we need to initialize the map
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}
	fmt.Println(stocks)

	// for loops always work with the maps
	for k, v := range stocks { // key & value
		fmt.Println(k, "->", v)
	}
	for _, v := range stocks { // values
		fmt.Println(v)
	}

	// and finally we can remove from a map
	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "APPL") // no panic. Same as working with nil map

}

// In this exercise I need to change the mostCommon  function to return n words in
// a slice. the new func sig should be:
// func mostCommon(r io.Reader,n int)([]string,error)
