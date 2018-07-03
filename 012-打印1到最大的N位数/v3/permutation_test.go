package main

import "testing"

func TestPermutation(t *testing.T) {
	stringSlice := []string{"a", "b", "c"}
	Permutation(stringSlice)
	//want := [][]string{
	//	[]string{"a", "b", "c"},
	//	[]string{"a", "c", "b"},
	//	[]string{"b", "a", "c"},
	//	[]string{"b", "c", "a"},
	//	[]string{"c", "a", "b"},
	//	[]string{"c", "b", "a"},
	//}
	//
	//if len(got) != len(want) {
	//	t.Errorf("got %v want %v", got, want)
	//}
}