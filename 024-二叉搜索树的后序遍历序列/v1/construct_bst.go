package main

// ConstructBST 题意见 https://www.geeksforgeeks.org/construct-a-binary-search-tree-from-given-postorder/
// post order 的流程：后续遍历左子树，再根节点，再后续遍历右子树。由于 BST 的
// 左节点 < 根节点 < 右节点，所以 postOrderTraversal 中符合这个位置的就是
// 左右子树的分割点，所以以此能获取左右子树
//
// 循环：获取根节点以及左右子树，返回构建后的树
// 终止条件：当前的 postOrderTraversal 节点数小于等于 1
func ConstructBST(postOrderTraversal []int) (rootNode *Node) {
	length := len(postOrderTraversal)

	if length == 0 {
		return
	} else if length == 1 {
		rootNode = &Node{
			nil,
			nil,
			postOrderTraversal[0],
		}
		return
	}

	leftSubTreeTraversal, rightSubTreeTraversal := GetSubTreeTraversal(postOrderTraversal)
	leftSubTree := ConstructBST(leftSubTreeTraversal)
	rightSubTree := ConstructBST(rightSubTreeTraversal)

	rootNode = &Node{
		leftSubTree,
		rightSubTree,
		postOrderTraversal[length-1],
	}
	return
}

// GetSubTreeTraversal 分别返回左右子树的 postOrderTraversal
func GetSubTreeTraversal(postOrderTraversal []int) (leftSubTreeTraversal []int,
	rightSubTreeTraversal []int) {

	originalLength := len(postOrderTraversal)
	if originalLength == 0 {
		return
	}

	rootNodeValue := postOrderTraversal[originalLength-1]

	currentLength := originalLength - 1
	// 剔除根节点，因为左右子树的 postOrderTraversal 不包含根节点
	postOrderTraversal = postOrderTraversal[:currentLength]

	partitionedIndex := -1

	if currentLength == 1 {
		if postOrderTraversal[0] < rootNodeValue {
			partitionedIndex = 1
		} else {
			partitionedIndex = 0
		}
	} else {
		for index, value := range postOrderTraversal {
			nextIndex := index + 1
			if value < rootNodeValue && rootNodeValue < postOrderTraversal[nextIndex] {
				partitionedIndex = nextIndex
			}
		}
	}

	leftSubTreeTraversal = postOrderTraversal[:partitionedIndex]
	rightSubTreeTraversal = postOrderTraversal[partitionedIndex:]
	return
}
