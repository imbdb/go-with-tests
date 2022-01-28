package main

import "testing"

func TestHello(t *testing.T) {
	// Reducing duplication
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Sub tests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("imbdb", "")
		want := "Hello, imbdb!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("imbdb", "spanish")
		want := "Hola, imbdb!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("imbdb", "french")
		want := "Bonjour, imbdb!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

}
