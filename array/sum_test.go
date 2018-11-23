package array

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, given)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("saying hello to people", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
