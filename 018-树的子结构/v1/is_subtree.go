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

	if parentTree != nil {
		if TreeIsEqual(parentTree, subTree) {
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

func TreeIsEqual(tree1 *Node, tree2 *Node) (result bool) {
	if !reflect.DeepEqual(
		tree1.PreOrderTraversal([]int{}),
		tree2.PreOrderTraversal([]int{}),
	) {
		return
	}

	if !reflect.DeepEqual(
		tree1.InOrderTraversal([]int{}),
		tree2.InOrderTraversal([]int{}),
	) {
		return
	}

	result = true
	return
}
