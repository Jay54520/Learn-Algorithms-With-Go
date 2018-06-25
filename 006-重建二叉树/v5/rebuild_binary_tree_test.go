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
	want := &root
	if !TreeEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TreeEqual(node1 *Node, node2 *Node) bool {
	var result1 []int
	var result2 []int
	result1 = node1.PreOrderTraversal(result1)
	result2 = node2.PreOrderTraversal(result2)
	if reflect.DeepEqual(result1, result2) {
		return false
	}
	return true
}
