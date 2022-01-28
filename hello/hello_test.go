package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("imbdb")
	want := "Hello, imbdb!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
