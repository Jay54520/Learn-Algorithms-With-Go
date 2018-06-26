package main

func RebuildBinaryTree(preOrderTraversalResult []int, inOrderTraversalResult []int) *Node {
	return rebuild(
		preOrderTraversalResult,
		0,
		len(preOrderTraversalResult)-1,
		inOrderTraversalResult,
		0,
		len(inOrderTraversalResult)-1,
	)
}

func SliceIndex(nums []int, predicate int) int {
	for i, num := range nums {
		if num == predicate {
			return i
		}
	}
	return -1
}

func rebuild(preOrderTraversalResult []int, preOrderStartIndex int, preOrderEndIndex int,
	inOrderTraversalResult []int, inOrderStartIndex int, inOrderEndIndex int) (rootNode *Node) {

	rootValue := preOrderTraversalResult[preOrderStartIndex]
	rootNode = &Node{
		left:  nil,
		right: nil,
		value: rootValue,
	}

	// 直接返回只有一个节点的树
	if preOrderStartIndex == preOrderEndIndex {
		return
	}

	rootIndexOfInOrderTraversalResult := SliceIndex(inOrderTraversalResult, rootValue)
	leftTreeNodeNum := rootIndexOfInOrderTraversalResult - inOrderStartIndex
	leftPreOrderEndIndex := preOrderStartIndex + leftTreeNodeNum

	if leftTreeNodeNum > 0 {
		rootNode.left = rebuild(
			preOrderTraversalResult,
			preOrderStartIndex+1, // 排除当前 root 节点
			leftPreOrderEndIndex,
			inOrderTraversalResult,
			inOrderStartIndex,
			rootIndexOfInOrderTraversalResult-1,
		)
	}

	nonRootNodeNum := preOrderEndIndex - preOrderStartIndex
	rightTreeNodeNum := nonRootNodeNum - leftTreeNodeNum
	if rightTreeNodeNum > 0 {
		rootNode.right = rebuild(
			preOrderTraversalResult,
			leftPreOrderEndIndex+1,
			preOrderEndIndex,
			inOrderTraversalResult,
			rootIndexOfInOrderTraversalResult+1,
			inOrderEndIndex,
		)
	}

	return
}
