package main

import (
	"testing"
	"reflect"
)

func TestPermutation(t *testing.T) {
	t.Run("unique characters", func(t *testing.T) {
		got := Permutation("abc")
		want := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("has same characters", func(t *testing.T) {
		got := Permutation("aba")
		want := []string{"aab", "aba", "baa"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("nil string", func(t *testing.T) {
		var s string
		got := Permutation(s)
		var want []string

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
