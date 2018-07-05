package main

import "testing"

func TestIsSubtree(t *testing.T) {
	t.Run("not empty trees", func(t *testing.T) {

		leftLeft := &Node{
			nil,
			nil,
			0,
		}

		left := &Node{
			leftLeft,
			nil,
			1,
		}
		right := &Node{
			nil,
			nil,
			2,
		}

		parentTree := &Node{
			left,
			right,
			0,
		}
		subTree := left

		got := IsSubtree(subTree, parentTree)
		if got != true {
			t.Errorf("%v should be true", got)
		}
	})

	t.Run("empty trees", func(t *testing.T) {
		got := IsSubtree(nil, nil)
		if got != true {
			t.Errorf("%v should be true", got)
		}

	})
}
