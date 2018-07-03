package main

import (
	"testing"
	"reflect"
)

func TestEvenOddArray(t *testing.T) {
	got := EvenOddArray([]int{12, 34, 45, 9, 8, 90, 3})
	want := []int{45, 9, 3, 12, 34, 8, 90}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
