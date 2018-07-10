package main

import (
	"testing"
	"reflect"
)

func TestKSumPathsFromRoot(t *testing.T) {
	leftLeft := &Node{
		nil,
		nil,
		1,
	}
	left := &Node{
		leftLeft,
		nil,
		28,
	}
	rightLeft := &Node{
		nil,
		nil,
		14,
	}
	rightRight := &Node{
		nil,
		nil,
		15,
	}
	right := &Node{
		rightLeft,
		rightRight,
		13,
	}
	root := &Node{
		left,
		right,
		10,
	}

	got := KSumPathsFromRoot(root, 38)
	want := [][]int{
		{10, 28},
		{10, 13, 15},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
