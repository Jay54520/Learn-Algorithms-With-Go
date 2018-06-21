package main

import "testing"

func TestNode(t *testing.T)  {
	lastNode := Node{
		3,
		nil,
	}
	secondNode := Node{
		2,
		&lastNode,
	}
	firstNode := Node{
		1,
		&secondNode,
	}

	t.Run("get node value", func(t *testing.T) {
		if lastNode.value != 3 {
			t.Errorf("got %v want %v", lastNode.value, 3)
		}
	})

	t.Run("nodes are linked", func(t *testing.T) {
		if firstNode.next.next != &lastNode {
			t.Errorf("got %v want %v", firstNode.next.next.value, &lastNode)
		}
	})

}
