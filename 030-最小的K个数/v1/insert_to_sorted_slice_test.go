package main

import (
	"testing"
	"reflect"
)

func TestInsertToSortedSlice(t *testing.T) {
	sortedSlice := []int{1, 2, 4}
	number := 3
	InsertToSortedSlice(&sortedSlice, number)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(sortedSlice, want) {
		t.Errorf("got %v want %v", sortedSlice, want)
	}
}