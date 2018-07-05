package main

import "reflect"

// IsSubtree 二叉树 subTree 是否是 parentTree 的子树
// 遍历 parentTree
// 终止条件：找到 subTree 或到了空节点
func IsSubtree(subTree *Node, parentTree *Node) (result bool) {
	// 空树是任何树的子树
	if subTree == nil {
		return true
	}

	inTraversalResults := subTree.InOrderTraversal([]int{})
	preTraversalResults := subTree.PreOrderTraversal([]int{})

	if parentTree != nil {
		if equalSubtree(parentTree, preTraversalResults, inTraversalResults) {
			return true
		}
	}

	if parentTree.left != nil {
		return IsSubtree(subTree, parentTree.left)
	}

	if parentTree.right != nil {
		return IsSubtree(subTree, parentTree.right)
	}

	return
}

func equalSubtree(tree1 *Node, preTraversalResults, inTraversalResults []int) (equality bool) {
	if !reflect.DeepEqual(
		tree1.PreOrderTraversal([]int{}),
		preTraversalResults,
	) {
		return
	}

	if !reflect.DeepEqual(
		tree1.InOrderTraversal([]int{}),
		inTraversalResults,
	) {
		return
	}

	equality = true
	return
}
