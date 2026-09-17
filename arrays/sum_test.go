package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{10, 10, 10, 10, 10}

	got := Sum(numbers)
	want := 50

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
