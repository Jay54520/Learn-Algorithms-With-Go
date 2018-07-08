package main

import (
	"testing"
	"reflect"
)

func TestConstructBST(t *testing.T) {
	leftLeft := &Node{
		nil,
		nil,
		1,
	}
	leftRight := &Node{
		nil,
		nil,
		7,
	}
	left := &Node{
		leftLeft,
		leftRight,
		5,
	}

	rightRight := &Node{
		nil,
		nil,
		50,
	}
	right := &Node{
		nil,
		rightRight,
		40,
	}

	root := &Node{
		left,
		right,
		10,
	}

	postOrderTraversal := []int{1, 7, 5, 50, 40, 10}
	got := ConstructBST(postOrderTraversal)
	if !TreeEqual(got, root) {
		t.Errorf("got %v want %v", got, root)
	}
}

func TestGetSubTreeTraversal(t *testing.T)  {
	t.Run("6 nodes", func(t *testing.T) {
		leftGot, rightGot := GetSubTreeTraversal([]int{1, 7, 5, 50, 40, 10})
		leftWant := []int{1, 7, 5}
		rightWant := []int{50, 40}

		if !reflect.DeepEqual(leftGot, leftWant) {
			t.Errorf("got %v want %v", leftGot, leftWant)
		}
		if !reflect.DeepEqual(rightGot, rightWant) {
			t.Errorf("got %v want %v", rightGot, rightWant)
		}
	})

	t.Run("empty", func(t *testing.T) {
		leftGot, rightGot := GetSubTreeTraversal([]int{})
		var leftWant []int
		var rightWant []int

		if !reflect.DeepEqual(leftGot, leftWant) {
			t.Errorf("got %v want %v", leftGot, leftWant)
		}
		if !reflect.DeepEqual(rightGot, rightWant) {
			t.Errorf("got %v want %v", rightGot, rightWant)
		}
	})

	t.Run("left left bst", func(t *testing.T) {
		leftGot, rightGot := GetSubTreeTraversal([]int{2, 1})
		leftWant := []int{}
		rightWant := []int{2}

		if !reflect.DeepEqual(leftGot, leftWant) {
			t.Errorf("got %v want %v", leftGot, leftWant)
		}
		if !reflect.DeepEqual(rightGot, rightWant) {
			t.Errorf("got %v want %v", rightGot, rightWant)
		}
	})
}

func TreeEqual(node1 *Node, node2 *Node) (equality bool) {
	var result1 []int
	var result2 []int
	result1 = node1.PreOrderTraversal(result1)
	result2 = node2.PreOrderTraversal(result2)
	if reflect.DeepEqual(result1, result2) {
		equality = true
	}
	return
}
