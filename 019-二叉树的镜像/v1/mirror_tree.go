package main

// MirrorTree 返回 originalTree 的镜像树，图片见 https://cdncontribute.geeksforgeeks.org/wp-content/uploads/Untitled-drawing-18.jpg
// 循环：交换当前节点的左右节点，如果它们都可以被交换；否则进入交换不可交换的子节点的循环。
// 可以交换：一个节点被标记为可交换，当它的左右节点被交换后
// 终止：完成了当前节点的左右节点交换，将叶子节点标记为可交换
func MirrorTree(originalTree *Node) (mirroredTree *Node) {
	exchange(originalTree)
	return originalTree
}

func exchange(node *Node) {
	if node.left == nil && node.right == nil {
		node.canExchange = true
	} else if node.left.canExchange && node.right.canExchange {
		left := node.left
		node.left = node.right
		node.right = left
		node.canExchange = true
	} else {
		if !node.left.canExchange {
			exchange(node.left)
		}
		if !node.right.canExchange {
			exchange(node.right)
		}
		exchange(node)
	}

}
