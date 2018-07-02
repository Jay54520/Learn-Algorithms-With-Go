package main

import (
	"testing"
	"reflect"
)

func TestGetN(t *testing.T) {
	t.Run("n > 0", func(t *testing.T) {

		got, _ := GetN(1)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("maxN is too big", func(t *testing.T) {

		_, err := GetN(20)

		if err == nil {
			t.Errorf("%v shouldn't be nil", err)
		}
	})

	t.Run("n <= 0", func(t *testing.T) {
		got, _ := GetN(-1)

		if len(got) != 0 {
			t.Errorf("%v must be empty", got)
		}

	})
}
