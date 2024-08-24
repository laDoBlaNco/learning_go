package nlp

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// one thing we like to do in testing is 'regression testing' or in Go 'table testing'
// one test for many different cases. But since testing only takes a t we need to pass it
// the cases in another way using an anony struct which we can tell by its {}{} structure

// var tokenizeCases = []struct { // anony struct '{}{}' we  used with json (json and testing are the most  common uses for anony structs)
// 	text   string
// 	tokens []string
// }{
// 	{"Who's on first?", []string{"who", "s", "on", "first"}},
// 	{"", nil},
// }

type tokenizeCase struct {
	Text   string
	Tokens []string // when dealing with serialization, the fields must be  exported
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	data, err := ioutil.ReadFile("./testdata/tokenize_cases.toml")
	require.NoError(t, err, "Read file")
	var testCases struct {
		Cases []tokenizeCase // tehse cases are tokenizeCase type from above as the toml has
		// Text and Tokens in each [[case]]
	}
	err = toml.Unmarshal(data, &testCases)
	require.NoError(t, err, "unmarshal TOML")
	return testCases.Cases // so this .Cases is linked to our testCases struct
}

// Excercise: Read test cases from tokenize_cases.toml

func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	// Before testify
	// if tokens!= expected{} // Can't compare slices with == in Go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got #%v", expected, tokens)
	}

}

// NOTE: If you have more test files in sub-folders use, go test ./... to test recursively

// there's a new type of testing as well call fuzz testing in which the test suite
// comes up with random parameters for our tests to pass, increasing accuracy that our tests
// are covering all the cases
func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _,tok:=range tokens{
			if !strings.Contains(lText,tok){
				t.Fatal(tok) 
			}
		}
	})
}
