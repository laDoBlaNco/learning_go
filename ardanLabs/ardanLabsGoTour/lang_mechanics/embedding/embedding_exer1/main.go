package main

/*
	This program defines a type Feed with two methods: Count and Fetch. We need to create
	a new type CachingFeed that embeds *Feed (pointer semantics) but overrides the
	Fetch method

	The CachingFeed type should ahve a map of Documents to limit the number of calls to
	Feed.Fetch
*/

import (
	"fmt"
	"log"
	"time"
)

// Document is the core data model we are working with
type Document struct {
	Key   string
	Title string
}

// ======================================================================
// Feed is a type that knows how to fetch Documents
type Feed struct{} // no fields as its just the behavior we are after

// Count tells how many doucments are in the feed
func (f *Feed) Count() int {
	return 42 // just for testing purposes
}

// Fetch simulates looking up a document specified by the key. It  is slow
func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

// ===================================================================
// Now FetchCounter is the behavior we depend on for our process function, so its our interface
// or method set
type FetchCounter interface {
	Fetch(key string) (Document, error) // NOTE the more specific method sig as we have args and
	// return values on this one.
	Count() int
}

func process(fc FetchCounter) { // here is our polymorphic function looking for the FetchCounter behavior
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Count not fetch %s : %v", key, err)
			return
		}
		fmt.Printf("%s : %v\n", key, doc)
	}

}

// =================================================================
// CachingFeed keeps a local copy of Documents that have already been retrieved. It embeds
// Feed to get the Fetch and Count behavior (for the FetchCounter interface) but overrides
// Fetch to ahve its own cache
type CachingFeed struct {
	// embed *Feed and add a field for a map[string]Document
	*Feed
	Docs map[string]Document
}

// NewCachingFeed will initialize a CachingFeed for use
func NewCachingFeed(f *Feed) *CachingFeed {
	// create  a CachingFeed with an initialized map and embedded feed.
	// Return its address (pointer)

	// res := CachingFeed{
	// 	f,
	// 	make(map[string]Document),
	// }
	// return &res
	
	return &CachingFeed{
		f,
		make(map[string]Document),
	}
	 

	// return nil // this needs to be removed after we create the real return. We do this now
	// just so gopls doesn't complain while we are prototyping
}

// Fetch calls the embedded  type's fetch method only if the key is not cached
func (cf *CachingFeed) Fetch(key string) (Document, error) {
	// TODO: implement this method that checks the map field for the specified key and
	// returns it if found. If not, it calls the embedded type's Fetch method. Store the
	// result in the map before returning it.
	doc,ok := cf.Docs[key]
	if !ok{
		doc,err:=cf.Feed.Fetch(key) 
		if err!=nil{
			return Document{},err
		}
		cf.Docs[key]=doc 
		return doc,nil
	}
	return doc, nil 
}

//================================================================
func main(){
	fmt.Println("Using feed directly")
	process(&Feed{}) 
	
	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{}) 
	process(c) 
	
}

// so in the end mine worked perfectly and I also noticed that the caching mechanism is of course
// faster than fetching it on every single call, which is the point of the cache. 
