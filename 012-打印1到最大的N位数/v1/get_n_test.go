package main

import (
	"testing"
	"reflect"
)

func TestGetN(t *testing.T)  {
	got, _ := GetN(1)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
