package main

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	stringSlice := []string{"a", "b", "c"}
	got := Permutation(stringSlice)
	want := [][]string{
		{"abc"},
		{"acb"},
		{"bac"},
		{"bca"},
		{"cab"},
		{"cba"},
	}
	
	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}
}