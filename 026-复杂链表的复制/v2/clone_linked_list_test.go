package main

import (
	"testing"
)

func TestCopyNextLinkedList(t *testing.T) {
	t.Run("three nodes linked list", func(t *testing.T) {
		thirdNode := &Node{
			3,
			nil,
			nil,
		}
		secondNode := &Node{
			2,
			thirdNode,
			nil,
		}
		firstNode := &Node{
			1,
			secondNode,
			nil,
		}

		newNode := copyNextLinkedList(firstNode)
		if !linkedListEqual(newNode, firstNode) {
			t.Errorf("got %v want %v", newNode, firstNode)
		}
	})

	t.Run("empty linked list", func(t *testing.T) {
		want := &Node{}
		got := copyNextLinkedList(want)
		if !linkedListEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestCopyLinkedList(t *testing.T) {
	t.Run("5 nodes", func(t *testing.T) {
		fifthNode := &Node{
			5,
			nil,
			nil,
		}
		fourthNode := &Node{
			4,
			fifthNode,
			nil,
		}
		thirdNode := &Node{
			3,
			fourthNode,
			nil,
		}
		secondNode := &Node{
			2,
			thirdNode,
			nil,
		}
		firstNode := &Node{
			1,
			secondNode,
			nil,
		}

		// 不能在初始化时定义 arbit 在前的 Node，由于那个 Node 当时还没有出现
		firstNode.arbit = thirdNode
		secondNode.arbit = firstNode
		thirdNode.arbit = fifthNode
		fourthNode.arbit = thirdNode
		fifthNode.arbit = secondNode

		got := CopyLinkedList(firstNode)
		want := firstNode

		if !linkedListEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("empty linked list", func(t *testing.T) {
		want := &Node{}
		got := CopyLinkedList(want)

		if !linkedListEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func linkedListEqual(node1 *Node, node2 *Node) bool {
	for node1 != nil && node2 != nil {
		if node1.value != node2.value {
			return false
		}
		node1 = node1.next
		node2 = node2.next
	}

	if node1 == nil && node2 == nil {
		return true
	} else {
		return false
	}

}
