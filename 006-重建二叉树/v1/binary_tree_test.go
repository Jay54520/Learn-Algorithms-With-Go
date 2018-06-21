package main

import "testing"

func TestBinaryTree(t *testing.T)  {
	level := Node{
		nil,
		nil,
		"value",
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
}
