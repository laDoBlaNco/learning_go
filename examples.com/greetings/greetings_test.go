package greetings

import(
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return val.
func TestHelloName(t *testing.T){
	name := "L@D0Bl@Nc0"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg,err := Hello("L@D0Bl@Nc0")
	if !want.MatchString(msg) || err != nil{
		t.Fatalf(`Hello("L@D0Bl@Nc0") = %q,%v,want match for %#q,nil`,msg,err,want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T){
	msg,err := Hello("")
	if msg != "" || err==nil{
		t.Fatalf(`Hello("") = %q,%v,want "",error`,msg,err)
	}
}
