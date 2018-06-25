package main

func RebuildBinaryTree(preOrderTraversalResult []int, inOrderTraversalResult []int) *Node {
	rootValue := preOrderTraversalResult[0]
	leftValue := inOrderTraversalResult[0]
	rightValue := inOrderTraversalResult[len(inOrderTraversalResult) - 1]

	leftNode := Node{
		left: nil,
		right: nil,
		value: leftValue,
	}
	rightNode := Node{
		left: nil,
		right: nil,
		value: rightValue,
	}
	rootNode := Node{
		left: &leftNode,
		right: &rightNode,
		value: rootValue,
	}
	return &rootNode
}
