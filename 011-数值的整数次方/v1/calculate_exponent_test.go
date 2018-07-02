package main

import "testing"

func TestCalculateExponent(t *testing.T)  {
	t.Run("exponent > 0", func(t *testing.T) {
		var base float64 = 2.2
		exponent := 3
		got, _ := CalculateExponent(base, exponent)
		want := 10.648000000000003

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("exponent == 0", func(t *testing.T) {
		var base float64 = 2.2
		exponent := 0
		got, _ := CalculateExponent(base, exponent)
		want := 1.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("exponent < 0 && base != 0", func(t *testing.T) {
		var base float64 = 2.2
		exponent := -3
		got, _ := CalculateExponent(base, exponent)
		want := 1.0 / 10.648000000000003

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("exponent < 0 && base == 0", func(t *testing.T) {
		var base float64 = 0
		exponent := -3
		_, err := CalculateExponent(base, exponent)

		if err == nil {
			t.Errorf("err %v shouldn't be nil", err)
		}
	})

}
