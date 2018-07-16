package main

import (
	"testing"
	"reflect"
)

func TestGetLeastNumbers(t *testing.T) {
	numbers := []int{4, 5, 1, 6, 2, 7, 3, 8}
	k := 4
	got := GetLeastNumbers(numbers, k)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
