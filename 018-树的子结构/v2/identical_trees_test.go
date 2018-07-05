package main

import "testing"

func TestIdenticalTrees(t *testing.T) {
	left := &Node{
		nil,
		nil,
		1,
	}
	right := &Node{
		nil,
		nil,
		2,
	}
	root := &Node{
		left,
		right,
		0,
	}

	t.Run("IdenticalTrees", func(t *testing.T) {
		got := IdenticalTrees(root, root)
		if !got {
			t.Errorf("%v sholuld be true", got)
		}
	})

	t.Run("Not identicalTrees", func(t *testing.T) {
		got := IdenticalTrees(root, left)
		if got {
			t.Errorf("%v sholuld be false", got)
		}
	})
}
