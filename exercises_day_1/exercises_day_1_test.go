package exercises_day_1

import "testing"

func assertIntegers(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but wanted %d", got, want)
	}
}

func TestPositiveSum(t *testing.T) {
	t.Run("only positives", func(t *testing.T) {
		got := PositiveSum([]int{1, 2, 3, 4, 5})
		want := 15
		assertIntegers(t, got, want)
	})

	t.Run("only negatives", func(t *testing.T) {
		got := PositiveSum([]int{-1, -2, -3, -4, -5})
		want := 0
		assertIntegers(t, got, want)
	})

	t.Run("any numbers", func(t *testing.T) {
		got := PositiveSum([]int{-1, 0, 2, -4, 5})
		want := 7
		assertIntegers(t, got, want)
	})
}
