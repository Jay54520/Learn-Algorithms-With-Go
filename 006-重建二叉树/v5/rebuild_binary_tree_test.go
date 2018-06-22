package main

import (
	"testing"
	"reflect"
)

func TestRebuildBinaryTree(t *testing.T) {
	preOrderTraversalResult := []int{0, 1, 2}
	inOrderTraversalResult := []int{1, 0, 2}
	got := RebuildBinaryTree(preOrderTraversalResult, inOrderTraversalResult)

	left := Node{
		nil,
		nil,
		1,
	}
	right := Node{
		nil,
		nil,
		2,
	}
	root := Node{
		&left,
		&right,
		0,
	}
	want := root
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
