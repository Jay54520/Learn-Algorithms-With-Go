package main

import "testing"

func TestIsPopOrder(t *testing.T) {
	t.Run("is pop order", func(t *testing.T) {
		pushOrder := []int{1, 2, 3}
		// push 1, push 2, pop 2, pop 1, push 3, pop3
		popOrder := []int{2, 1, 3}
		got := IsPopOrder(pushOrder, popOrder)
		if !got {
			t.Errorf("%v should be true", got)
		}
	})

	t.Run("is not pop order", func(t *testing.T) {
		pushOrder := []int{1, 2, 3}
		popOrder := []int{3, 1, 2}
		got := IsPopOrder(pushOrder, popOrder)
		if got {
			t.Errorf("%v should be false", got)
		}
	})

	t.Run("empty", func(t *testing.T) {
		var pushOrder []int
		var popOrder []int
		got := IsPopOrder(pushOrder, popOrder)
		if !got {
			t.Errorf("%v should be true", got)
		}
	})

}
