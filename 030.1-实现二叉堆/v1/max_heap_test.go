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

func TestMaxIntHeap_Pop(t *testing.T) {
	var num int
	var err error
	h := &MaxIntHeap{}
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)
	h.Push(6)

	var got []int
	num, err = h.Pop()
	for err == nil {
		got = append(got, num)
		num, err = h.Pop()
	}
	want := []int{6, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
