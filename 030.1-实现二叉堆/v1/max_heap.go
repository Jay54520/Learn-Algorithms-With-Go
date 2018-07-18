package main

const rootIndex = 1
const placeHolder = 0

// 使用数组实现最大堆，并含有 Push 和 Pop 方法
type MaxIntHeap []int

// 将 i 添加到数组末尾，然后使其上浮（与父节点交换）到合适位置。以索引 1 为
// 根节点（第一个节点），i 表示第 i 个节点，由于二叉堆是完全二叉树，所以左右子节点（如果有的话）
// 分别为 2i、2i + 1。最大堆的关系是父节点大于两个左右子节点，以下简称关系。
// 添加一个节点后，如果添加节点不满足关系，那么上浮它
// 终止条件：满足关系或者没有父节点了（它变成了根节点）
func (h *MaxIntHeap) Push(newNum int) {
	if len(*h) == 0 {
		*h = append(*h, placeHolder)
	}

	*h = append(*h, newNum)
	newNumIndex := len(*h) - 1
	relation := h.getRelation(newNumIndex)
	for !relation {
		parentIndex := newNumIndex / 2
		(*h)[parentIndex], (*h)[newNumIndex] = (*h)[newNumIndex], (*h)[parentIndex]
		relation = h.getRelation(parentIndex)
	}
}

// 是否满足关系或没有父节点
func (h *MaxIntHeap) getRelation(index int) bool {
	if index == rootIndex {
		return true
	}

	parentIndex := index / 2
	// 父节点大于子节点
	if (*h)[parentIndex] > (*h)[index] {
		return true
	} else {
		return false
	}
}
