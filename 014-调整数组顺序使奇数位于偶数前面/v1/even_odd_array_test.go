package main

import (
	"testing"
	"reflect"
)

func TestEvenOddArray(t *testing.T) {
	got := EvenOddArray([]int{5, 1, 2, 3, 4})
	want := []int{1, 3, 5, 2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
