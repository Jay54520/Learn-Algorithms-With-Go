package main

import "testing"

func TestFindMin(t *testing.T) {
	t.Run("minimum at middle + 1", func(t *testing.T) {
		array := []int{3, 4, 5, 1, 2}
		got := FindMin(array)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("minimum at first", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		got := FindMin(array)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
