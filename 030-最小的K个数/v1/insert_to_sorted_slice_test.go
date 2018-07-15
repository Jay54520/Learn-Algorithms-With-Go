package main

import (
	"testing"
	"reflect"
)

func TestInsertToSortedSlice(t *testing.T) {
	t.Run("three elements slice", func(t *testing.T) {
		sortedSlice := []int{1, 2, 4}
		number := 3
		InsertToSortedSlice(&sortedSlice, number)
		want := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(sortedSlice, want) {
			t.Errorf("got %v want %v", sortedSlice, want)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		sortedSlice := make([]int, 0, 1)
		number := 1
		InsertToSortedSlice(&sortedSlice, number)
		want := []int{1}

		if !reflect.DeepEqual(sortedSlice, want) {
			t.Errorf("got %v want %v", sortedSlice, want)
		}

	})
}