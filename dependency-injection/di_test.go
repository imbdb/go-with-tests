package dependencyinjection

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGreet(t *testing.T) {
	Convey("Should greet given name", t, func() {
		buffer := bytes.Buffer{}
		Greet(&buffer, "imbdb")

		got := buffer.String()
		want := "Hello, imbdb!"

		So(got, ShouldEqual, want)
	})

}
