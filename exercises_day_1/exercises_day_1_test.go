package exercises_day_1

import (
	"reflect"
	"testing"
)

func assertEqualsInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but wanted %d", got, want)
	}
}

func assertEqualsStringSlices(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPositiveSum(t *testing.T) {
	t.Run("only positives", func(t *testing.T) {
		got := PositiveSum([]int{1, 2, 3, 4, 5})
		want := 15
		assertEqualsInt(t, got, want)
	})

	t.Run("only negatives", func(t *testing.T) {
		got := PositiveSum([]int{-1, -2, -3, -4, -5})
		want := 0
		assertEqualsInt(t, got, want)
	})

	t.Run("any numbers", func(t *testing.T) {
		got := PositiveSum([]int{-1, 0, 2, -4, 5})
		want := 7
		assertEqualsInt(t, got, want)
	})
}

func TestMaxDiff(t *testing.T) {
	t.Run("max diff 4-1", func(t *testing.T) {
		got := MaxDiff([]int{1, 2, 3, 4})
		want := 3
		assertEqualsInt(t, got, want)
	})

	t.Run("max diff 3 - (-4)", func(t *testing.T) {
		got := MaxDiff([]int{1, 2, 3, -4})
		want := 7
		assertEqualsInt(t, got, want)
	})
}

func TestRemDups(t *testing.T) {
	t.Run("3 strings, 2 have dups", func(t *testing.T) {
		got := RemDups([]string{"abracadabra", "allottee", "assessee"})
		want := []string{"abracadabra", "alote", "asese"}
		assertEqualsStringSlices(t, got, want)
	})
}
