package main

import "testing"

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
}
