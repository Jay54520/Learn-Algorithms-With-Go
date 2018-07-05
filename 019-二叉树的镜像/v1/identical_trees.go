package main

// IdenticalTrees 判断两棵树是否相等
// 循环：判断对应位置的节点
// 终止条件：存在节点为 nil、节点不相等
func IdenticalTrees(tree1 *Node, tree2 *Node) (identity bool) {

	if tree1 == nil || tree2 == nil {
		if tree1 == nil && tree2 == nil {
			return true
		} else {
			return false
		}
	}

	if tree1.value != tree2.value {
		return
	}

	return IdenticalTrees(tree1.left, tree2.left) && IdenticalTrees(tree1.right, tree2.right)
}
