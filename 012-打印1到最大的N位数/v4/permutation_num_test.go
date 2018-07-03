package main

import (
	"testing"
	"reflect"
)

func TestPermutationNum(t *testing.T) {
	got := PermutationNum(1)
	want := []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
