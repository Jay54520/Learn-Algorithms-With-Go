package main

import "testing"

func TestNode(t *testing.T)  {
	lastNode := Node{
		3,
		nil,
	}

	t.Run("get node value", func(t *testing.T) {
		if lastNode.value != 3 {
			t.Errorf("got %v want %v", lastNode.value, 3)
		}
	})

}
