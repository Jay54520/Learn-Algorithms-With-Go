package main

import (
	"testing"
	"reflect"
)

func TestReverseLinkedList(t *testing.T) {
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

	t.Run("reverse 3 nodes linked list", func(t *testing.T) {
		got := ReverseLinkedList(&firstNode)
		want := []int{3, 2, 1}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("empty linked list", func(t *testing.T) {
		got := ReverseLinkedList(nil)

		if got != nil {
			t.Errorf("%v should be nil", got)
		}
	})

}
