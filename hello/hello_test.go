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

	testHelloHelper := func(t testing.T, testName string, name string, language string, want string) {
		t.Helper()
		t.Run(testName, func(t *testing.T) {
			got := Hello(name, language)
			assertCorrectMessage(t, got, want)
		})
	}

	// Sub tests
	testHelloHelper(*t, "saying hello to people", "imbdb", "", "Hello, imbdb!")

	testHelloHelper(*t, "saying hello to people in spanish", "imbdb", "spanish", "Hola, imbdb!")

	testHelloHelper(*t, "saying hello to people in french", "imbdb", "french", "Bonjour, imbdb!")

	testHelloHelper(*t, "saying 'Hello, World!' when an empty string is supplied", "", "", "Hello, World!")

}
