package arraynslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum of collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected '%d' but got '%d', given %v", expected, got, numbers)
		}
	})

}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTail(t *testing.T) {
	t.Run("Tail Sum of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely tail sum empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
