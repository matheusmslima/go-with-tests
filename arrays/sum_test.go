package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{10, 10, 10, 10, 10}

		got := Sum(numbers)
		want := 50

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
