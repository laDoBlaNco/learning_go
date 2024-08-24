package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

// this is where we set up the testing environment
var testApp Config

func TestMain(m *testing.M) { // in the func we'll set up things we need in place before testing
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	os.Exit(m.Run()) // this will actually run the tests
}

// here we have the json that we'll return when setting up our custom client. We do this so that
// we can run our test without actually having to have a network connection. If we don't do it
// this way, then we can only test when we are online and when the site is up. that puts limits
// on our testing.
var jsonToReturn = `
{"ts":1681993104192,"tsj":1681993096262,"date":"Apr 20th 2023, 08:18:16 am NY","items":[{"curr":"USD","xauPrice":2000.01,"xagPrice":25.2937,"chgXau":6.335,"chgXag":0.0367,"pcXau":0.3178,"pcXag":0.1453,"xauClose":1993.675,"xagClose":25.257}]}
`

// now we set up our custom client.
type RoundTripFunc func(req *http.Request) *http.Response // this is our func type

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
