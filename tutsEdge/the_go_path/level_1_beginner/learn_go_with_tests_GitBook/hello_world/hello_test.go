package main

import "testing"

// no with subtests
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		// We can refactor, instead of repeating the above action of checking if got and want match, we can
		// do as Radu always says, 'when we have repeated functionality, create a function
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})
}

// outside of TestHello
func assertCorrectMessage(t testing.TB, got, want string) { // testing.TB so we can run from test or benchmark
	t.Helper() // this tells our testing suite that this is a helper func and not a test func
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
