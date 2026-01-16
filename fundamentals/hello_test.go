package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to your friends", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to strangers", func(t *testing.T) {
		got := Hello("")
		want := "Hello, stranger!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want %q", got, want)
	}
}
