package main

// MirrorTree 返回 originalTree 的镜像树，图片见 https://cdncontribute.geeksforgeeks.org/wp-content/uploads/Untitled-drawing-18.jpg
// 把树看做由左右子树和根节点递归构成的集合
// 循环：使当前节点的左右子树完成交换，然后交换当前节点的左右子树
// 终止：当前节点为空则返回
func MirrorTree(originalTree *Node) {
	if originalTree == nil {
		return
	}

	MirrorTree(originalTree.left)
	MirrorTree(originalTree.right)

	originalTree.left, originalTree.right = originalTree.right, originalTree.left
}