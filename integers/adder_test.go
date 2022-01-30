package integers

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdder(t *testing.T) {
	Convey("Should return sum of the arguments", t, func() {
		sum := Add(2, 2)
		expected := 4

		So(sum, ShouldEqual, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
