package main

import (
	"testing"
	"container/heap"
	"reflect"
)

func TestMaxIntHeap(t *testing.T) {
	h := &MaxIntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 10)
	var got []int
	for h.Len() > 0 {
		got = append(got, heap.Pop(h).(int))
	}

	want := []int{10, 5, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
