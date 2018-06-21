package main

import "testing"

func TestBinaryTree(t *testing.T)  {
	level := Node{
		nil,
		nil,
		"value",
	}
	left := Node{
		nil,
		nil,
		"left",
	}
	right := Node{
		nil,
		nil,
		"right",
	}
	root := Node{
		&left,
		&right,
		"root",
	}

	t.Run("level node", func(t *testing.T) {
		if level.left != nil {
			t.Errorf("%v doesn't have left node", level)
		}
		if level.right != nil {
			t.Errorf("%v doesn't have right node", level)
		}
		if level.value != "value" {
			t.Errorf("got %v want %v", level.value, "value")
		}
	})
	
	t.Run("three nodes binary tree", func(t *testing.T) {
		if root.left != &left {
			t.Errorf("got %v want %v", root.left, &left)
		}
		if root.right != &right {
			t.Errorf("got %v want %v", root.right, &right)
		}
		if root.value != "root" {
			t.Errorf("got %v want %v", root.value, "root")
		}
	})
}
