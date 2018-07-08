package main

import "testing"

func TestGetMin(t *testing.T)  {
	t.Run("two elements stack", func(t *testing.T) {
		var stack Stack
		stack.Push(1)
		stack.Push(2)

		got := stack.GetMin()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		var stack Stack
		got := stack.GetMin()
		want := -1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
