package main

func RebuildBinaryTree(preOrderTraversalResult []int, inOrderTraversalResult []int) *Node {
	return &Node{
		value: 0,
		left: nil,
		right: nil,
	}
}
