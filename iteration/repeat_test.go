package iteration

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepeat(t *testing.T) {
	Convey("Should repeat character given number of times", t, func() {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		So(repeated, ShouldEqual, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}
