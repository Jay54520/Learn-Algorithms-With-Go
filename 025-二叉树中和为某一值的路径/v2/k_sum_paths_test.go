package main

import (
	"testing"
	"reflect"
)

func TestKSumPaths(t *testing.T) {
	leftLeft := &Node{
		nil,
		nil,
		2,
	}
	leftRightLeft := &Node{
		nil,
		nil,
		1,
	}
	leftRight := &Node{
		leftRightLeft,
		nil,
		1,
	}
	left := &Node{
		leftLeft,
		leftRight,
		3,
	}

	rightRightRight := &Node{
		nil,
		nil,
		6,
	}
	rightRight := &Node{
		nil,
		rightRightRight,
		5,
	}
	rightLeftLeft := &Node{
		nil,
		nil,
		1,
	}
	rightLeftRight := &Node{
		nil,
		nil,
		2,
	}
	rightLeft := &Node{
		rightLeftLeft,
		rightLeftRight,
		4,
	}
	right := &Node{
		rightLeft,
		rightRight,
		-1,
	}

	root := &Node{
		left,
		right,
		1,
	}

	got := KSumPaths(root, 5)
	want := [][]int{
		{3, 2},
		{3, 1, 1},
		{1, 3, 1},
		{4, 1},
		{1, -1, 4, 1},
		{-1, 4, 2},
		{5},
		{1, -1, 5},
	}

	if !sliceElementsEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// sliceElementsEqual 两个 slices 的元素相同，与顺序无关
func sliceElementsEqual(slice1 [][]int, slice2 [][]int) bool {
	if !(len(slice1) == len(slice2)) {
		return false
	}

	for _, value := range slice1 {
		if !inSlice(slice2, value) {
			return false
		}
	}

	return true
}

// inSlice is value in slice or not
func inSlice(slice [][]int, value []int) bool {
	for _, sliceValue := range slice {
		if reflect.DeepEqual(value, sliceValue) {
			return true
		}
	}
	return false
}
