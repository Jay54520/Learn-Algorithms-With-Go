package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestRebuildBinaryTree(t *testing.T) {
	t.Run("simple three nodes tree", func(t *testing.T) {
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
	})

	t.Run("four nodes tree", func(t *testing.T) {
		preOrderTraversalResult := []int{0, 1, 2, 3}
		inOrderTraversalResult := []int{2, 1, 0, 3}
		got := RebuildBinaryTree(preOrderTraversalResult, inOrderTraversalResult)

		leftLeft := Node{
			nil,
			nil,
			2,
		}
		left := Node{
			&leftLeft,
			nil,
			1,
		}
		right := Node{
			nil,
			nil,
			3,
		}
		root := Node{
			&left,
			&right,
			0,
		}
		want := &root
		if !TreeEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("no node", func(t *testing.T) {
		preOrderTraversalResult := []int{}
		inOrderTraversalResult := []int{}
		got := RebuildBinaryTree(preOrderTraversalResult, inOrderTraversalResult)

		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}

	})
}

func TreeEqual(got *Node, want *Node) (equality bool) {
	var got_result []int
	var want_result []int
	got_result = got.PreOrderTraversal(got_result)
	want_result = want.PreOrderTraversal(want_result)
	if reflect.DeepEqual(got_result, want_result) {
		equality = true
	} else {
		fmt.Printf("got %v want %v", got_result, want_result)
	}

	return
}
