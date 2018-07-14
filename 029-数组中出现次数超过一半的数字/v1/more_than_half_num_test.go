package main

import "testing"

func TestMoreThanHalfNum(t *testing.T) {
	t.Run("has more than half num", func(t *testing.T) {
		nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
		got := MoreThanHalfNum(nums)
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
