package main

import "testing"

func TestMirrorTree(t *testing.T)  {
	originalLeftLeft := &Node{
		nil,
		nil,
		5,
		false,
	}
	originalRightRight := &Node{
		nil,
		nil,
		4,
		false,
	}
	originalRight := &Node{
		originalLeftLeft,
		originalRightRight,
		2,
		false,
	}
	originalLeft := &Node{
		nil,
		nil,
		3,
		false,
	}
	originalTree := &Node{
		originalLeft,
		originalRight,
		1,
		false,
	}
	
	mirroredLeftRight := &Node{
		nil,
		nil,
		5,
		false,
	}
	mirroredLeftLeft := &Node{
		nil,
		nil,
		4,
		false,
	}
	mirroredLeft := &Node{
		mirroredLeftLeft,
		mirroredLeftRight,
		2,
		false,
	}
	mirroredRight := &Node{
		nil,
		nil,
		3,
		false,
	}
	mirroredTree := &Node{
		mirroredLeft,
		mirroredRight,
		1,
		false,
	}

	MirrorTree(originalTree)
	want := mirroredTree

	if !IdenticalTrees(originalTree, want) {
		t.Errorf("got %v want %v", originalTree, want)
	}
}
