package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoWithTests(t *testing.T) {

	Convey("Should be Go with Tests", t, func() {
		str := GoWithTests()
		So(str, ShouldEqual, "Go with Tests")
	})

}
