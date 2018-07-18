package main

import (
	"testing"
	"reflect"
)

func TestMaxIntHeap_Push(t *testing.T) {
	h := &MaxIntHeap{}
	h.Push(1)
	h.Push(2)
	h.Push(3)
	want := MaxIntHeap{3, 1, 2}
	if !reflect.DeepEqual(*h, want) {
		t.Errorf("got %v want %v", *h, want)
	}
}