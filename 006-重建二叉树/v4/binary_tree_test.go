package main

import (
	"testing"
	"reflect"
)

func TestBinaryTree(t *testing.T)  {
	level := Node{
		nil,
		nil,
		0,
	}
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

	t.Run("level node", func(t *testing.T) {
		if level.left != nil {
			t.Errorf("%v doesn't have left node", level)
		}
		if level.right != nil {
			t.Errorf("%v doesn't have right node", level)
		}
		if level.value != 0 {
			t.Errorf("got %v want %v", level.value, 0)
		}
	})
	
	t.Run("three nodes binary tree", func(t *testing.T) {
		if root.left != &left {
			t.Errorf("got %v want %v", root.left, &left)
		}
		if root.right != &right {
			t.Errorf("got %v want %v", root.right, &right)
		}
		if root.value != 0 {
			t.Errorf("got %v want %v", root.value, 0)
		}
	})

	t.Run("pre order traversal", func(t *testing.T) {
		var got []int
		got = root.PreOrderTraversal(got)
		want := []int{0, 1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("pre order traversal with one node", func(t *testing.T) {
		var got []int
		got = level.PreOrderTraversal(got)
		want := []int{0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("in order traversal", func(t *testing.T) {
		var got []int
		got = root.InOrderTraversal(got)
		want := []int{1, 0, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("breadth first traversal", func(t *testing.T) {
		var got []int
		got = root.BreadthFirstTraversal(got)
		want := []int{0, 1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
