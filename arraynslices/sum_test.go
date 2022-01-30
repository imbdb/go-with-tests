package arraynslices

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {

	Convey("sum of collection of any size", t, func() {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		So(got, ShouldEqual, expected)
	})

}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	Convey("Sum of some slices", t, func() {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		So(got, ShouldResemble, want)
	})

	Convey("safely sum empty slices", t, func() {
		got := SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}

		So(got, ShouldResemble, want)
	})
}

func TestSumAllTail(t *testing.T) {
	Convey("Tail Sum of some slices", t, func() {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		So(got, ShouldResemble, want)
	})

	Convey("safely tail sum empty slices", t, func() {
		got := SumAllTail([]int{}, []int{0, 9})
		want := []int{0, 9}

		So(got, ShouldResemble, want)
	})
}

func TestSumAllHead(t *testing.T) {
	Convey("Head Sum of some slices", t, func() {
		got := SumAllHead([]int{1, 2, 3}, []int{0, 9, 8, 1})
		want := []int{3, 17}

		So(got, ShouldResemble, want)
	})

	Convey("safely head sum empty slices", t, func() {
		got := SumAllHead([]int{}, []int{0, 9, 8, 1})
		want := []int{0, 17}

		So(got, ShouldResemble, want)
	})
}
