package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to your friends", func(t *testing.T) {
		got := Hello("Chris", "eng")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to strangers", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, stranger!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in polish", func(t *testing.T) {
		got := Hello("Krystian", "pol")
		want := "Dzien dobry, Krystian!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Krystian", "esp")
		want := "Hola, Krystian!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want %q", got, want)
	}
}
