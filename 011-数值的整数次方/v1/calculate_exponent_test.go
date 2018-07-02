package main

import "testing"

func TestCalculateExponent(t *testing.T)  {
	t.Run("exponent > 0", func(t *testing.T) {
		var base float64 = 2.2
		exponent := 3
		got := CalculateExponent(base, exponent)
		want := 10.648000000000003

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("exponent == 0", func(t *testing.T) {
		var base float64 = 2.2
		exponent := 0
		got := CalculateExponent(base, exponent)
		want := 1.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
