package main

import "errors"

const rootIndex = 0

// 使用数组实现最大堆，并含有 Push 和 Pop 方法
type MaxIntHeap []int

// 将 i 添加到数组末尾，然后使其上浮（与父节点交换）到合适位置。
// 2i + 1, 2i + 2 分别为 i 节点的左右子节点，比如索引分别为 0 1 2 3 4 5 6。
// i 的父节点为 (i - 1) /2。因为 (2i +1 - 1)/2 = i 和 (2i + 2 - 1) / 2 = i
// 最大堆的关系是父节点大于两个左右子节点，左右节点之间没有大小关系，以下简称关系。
//
// 添加一个节点后，新增节点是一个叶子节点。如果新增节点小于它的父节点，那么符合关系，返回；
//否则，新增节点大于它的父节点。由于添加节点前二叉堆处于有序状态，所以如果父节点存在一个左子节点，
//那么父节点大于左子节点，因此新增节点也大于左子节点。将新增节点与父节点互换后，以新增节点为根节点的
//树符合关系。同理，新增节点再次与它的父节点比较并决定返回或是交换。如果新增节点变成了根节点，那么以
//新增节点为根节点的树就是整棵树，即整棵树符合关系了。
//
// 将节点添加到末尾，然后
// 循环代码为：查看新增节点是否满足终止条件，如果满足，则返回；否则将新增节点与其父节点交换
// 终止条件：新增节点变成了根节点或者新增节点小于它的父节点
func (h *MaxIntHeap) Push(element int) {
	var parentIndex int
	*h = append(*h, element)
	elementIndex := len(*h) - 1
	shouldStopUp := h.shouldStopUp(elementIndex)

	for !shouldStopUp {
		parentIndex = (elementIndex - 1) / 2
		(*h)[parentIndex], (*h)[elementIndex] = (*h)[elementIndex], (*h)[parentIndex]
		elementIndex = parentIndex
		shouldStopUp = h.shouldStopUp(elementIndex)
	}
}

func (h *MaxIntHeap) shouldStopUp(index int) (shouldStopUp bool) {
	// 因为根节点没有父节点，所以要先判断这个
	if index == rootIndex {
		shouldStopUp = true
	}
	parentIndex, _ := h.getParentIndex(index)
	if (*h)[index] < (*h)[parentIndex] {
		shouldStopUp = true
	}
	return
}


func (h *MaxIntHeap) getParentIndex(childIndex int) (parentIndex int, err error)  {
	if childIndex == rootIndex {
		err = errors.New("根节点没有父节点")
		return
	}
	parentIndex = (childIndex - 1) / 2
	return
}