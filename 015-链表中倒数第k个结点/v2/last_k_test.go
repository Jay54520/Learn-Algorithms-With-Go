package main

import "testing"

func TestLastK(t *testing.T) {
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

	t.Run("get last 2 node", func(t *testing.T) {
		got := LastK(&firstNode, 2)
		want := &secondNode

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("k greater than length", func(t *testing.T) {
		got := LastK(&firstNode, 4)

		if got != nil {
			t.Errorf("got %v should be nil", got)
		}
	})

}
