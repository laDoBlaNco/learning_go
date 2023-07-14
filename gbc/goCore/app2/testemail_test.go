package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil { // o sea, we know this should be wrong, so if we didn't get an err
		// we know our function isn't working
		t.Error("hello is not an email, check your func")
	}

	_, err = IsEmail("whitesidekevin@gmail.com")
	if err != nil { // this should pass so our err should == nil
		t.Error("whitsidekevin@gmail is an email, check your func")
	}
}
