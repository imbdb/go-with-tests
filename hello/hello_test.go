package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	// Sub tests
	Convey("saying hello to people", t, func() {
		got := Hello("imbdb", "")
		expected := "Hello, imbdb!"
		So(got, ShouldEqual, expected)
	})

	Convey("saying hello to people in spanish", t, func() {
		got := Hello("imbdb", "spanish")
		expected := "Hola, imbdb!"
		So(got, ShouldEqual, expected)
	})

	Convey("saying hello to people in french", t, func() {
		got := Hello("imbdb", "french")
		expected := "Bonjour, imbdb!"
		So(got, ShouldEqual, expected)
	})

	Convey("saying 'Hello, World!' when an empty string is supplied", t, func() {
		got := Hello("", "")
		expected := "Hello, World!"
		So(got, ShouldEqual, expected)
	})

}
