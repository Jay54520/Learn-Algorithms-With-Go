package main

type Node struct {
	left *Node
	right *Node
	value int
}

func (node *Node) PreOrderTraversal() []int {
	return []int{node.value, node.left.value, node.right.value}
}