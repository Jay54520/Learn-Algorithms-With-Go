package main

import "testing"

func TestFib(t *testing.T) {
	t.Run("base condition", func(t *testing.T) {
		got := Fib(0)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		got1 := Fib(1)
		want1 := 1

		if got1 != want1 {
			t.Errorf("got %d want %d", got1, want1)
		}
	})

	t.Run("normal condition", func(t *testing.T) {
		got := Fib(3)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkFib(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Fib(10)
	}
}

