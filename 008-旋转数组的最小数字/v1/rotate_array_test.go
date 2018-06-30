package main

import (
	"testing"
	"reflect"
)

func TestRotateArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	got := RotateArray(array, k)
	want := []int{5, 6, 7, 1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
