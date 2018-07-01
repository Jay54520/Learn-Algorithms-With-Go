package main

import "testing"

func TestStepJump(t *testing.T)  {
	t.Run("one step", func(t *testing.T) {
		got := StepJump(1)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("two steps", func(t *testing.T) {
		got := StepJump(2)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("three steps", func(t *testing.T) {
		got := StepJump(3)
		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}
