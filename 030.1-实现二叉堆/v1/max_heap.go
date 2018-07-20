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
//
//时间复杂度：最多要比较 h - 1 次、交换 h - 1 次，所以近似为 2h。假设高度为 h，那么一共有 (2^0 + 2^1 + ... + 2^(h-1)) = 2^h - 1 个节点（根据等比数列#求和公式）
//所以，如果有 n 个节点，那么 h = log(n)，所以为时间复杂度 O(2h) = O(2log(n)) = O(log(n))
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

func (h *MaxIntHeap) getParentIndex(childIndex int) (parentIndex int, err error) {
	if childIndex == rootIndex {
		err = errors.New("根节点没有父节点")
		return
	}
	parentIndex = (childIndex - 1) / 2
	return
}

// 设置弹出的结果为根节点的值，然后将根节点与最后一个节点交换，然后删除最后一个节点。设置根节点为指定节点。
// 推导过程与上面的 Push 相同，只是由不能上移变成不能下移，由比较新增节点与父节点变为比较指定节点与最大子节点。
// 将根节点与最后一个节点交换，可以简化代码的写法，这种方法叫什么？
// 如果指定节点不满足关系，那么指定节点必须与最大子节点交换，这样交换后的子树（最大子节点 - 指定节点 - （可能的）另一个子节点）才能满足关系。
//
// 将弹出的结果设置为根节点的值，然后将根节点与最后一个节点交换，然后删除最后一个节点。将根节点指定为指定节点。
// 循环代码为：查看指定节点是否满足终止条件，如果满足，则返回结果；否则将指定节点与其最大子节点交换
// 终止条件：指定节点变成了叶子节点或者指定节点大于它的所有子节点
func (h *MaxIntHeap) Pop() (result int, err error) {
	if len(*h) == 0 {
		err = errors.New("不能对空 MaxIntHeap 执行 Pop()")
		return
	}

	result = (*h)[rootIndex]
	endIndex := len(*h) - 1
	(*h)[rootIndex], (*h)[endIndex] = (*h)[endIndex], (*h)[rootIndex]
	*h = (*h)[:endIndex]

	h.down(rootIndex)
	return
}

// 将 index 下沉到合适的位置使二叉堆保持有效。index 下方的二叉堆都是有效的。
func (h *MaxIntHeap) down(index int) {
	shouldStopDown := h.shouldStopDown(index)
	var maxChildIndex int

	for !shouldStopDown {
		maxChildIndex, _ = h.getMaxChildIndex(index)
		(*h)[maxChildIndex], (*h)[index] = (*h)[index], (*h)[maxChildIndex]
		index = maxChildIndex
		shouldStopDown = h.shouldStopDown(index)
	}

}


func (h *MaxIntHeap) shouldStopDown(index int) bool {
	maxChildIndex, err := h.getMaxChildIndex(index)

	// 有错误，说明是叶子节点。能不能这样使用错误，如果有多种错误怎么处理？
	if err != nil {
		return true
	}

	if (*h)[index] > (*h)[maxChildIndex] {
		// 大于最大子节点就说明大于所有子节点
		return true
	} else {
		return false
	}
}
func (h *MaxIntHeap) getMaxChildIndex(parentIndex int) (maxChildIndex int, err error) {
	leftChildIndex := 2*parentIndex + 1
	if leftChildIndex >= len(*h) {
		err = errors.New("leaf node doesn't have children")
		return
	}
	rightChildIndex := 2*parentIndex + 2
	if rightChildIndex < len(*h) {
		// 同时拥有左右节点
		if (*h)[leftChildIndex] > (*h)[rightChildIndex] {
			maxChildIndex = leftChildIndex
		} else {
			maxChildIndex = rightChildIndex
		}
	} else {
		// 只拥有左节点
		maxChildIndex = leftChildIndex
	}
	return
}

// 从高度为 0 开始到根节点，每次构建一层的二叉堆。
// 高度为0的层只有一个节点，那么都是有效的二叉堆。然后构建高度为1的二叉堆，因为高度为0的二叉堆是有效的二叉堆，
// 那么只需要使高度为 1 的节点和高度为 0 的二叉堆符合关系。
// 终止条件：根节点为有效的二叉堆
//高度为 h 的节点的索引范围是 [n / 2^(h+1), n / 2^h) （h >= 0），当 n / 2^(h+1) == rootIndex 时，
//说明到了根节点，完成这个后整棵树就是有效的二叉堆了，因为与 rootIndex 同高度的只有它本身，所以不要考虑区间。
// 当前的索引范围等于上一个索引范围除以 2。
//
// 时间复杂度：每个高度的所有节点都需要交换：(2^h-1 + 2^h-2 + ... + 2^0) = 2^h = n
// 所以为 O(n)
// 而对于 n 个节点，每次插入需要 n * O(logn) = O(nlogn)。因为上面是整棵树移动，而这里是单个节点移动。
// 利用了树递归相似
func (h *MaxIntHeap) Build() {
	length := len(*h)

	if length == 0 {
		return
	}

	startIndex := length / 2
	endIndex := length

	for {
		// 对于 [startIndex, endIndex)，要使以它们为根节点的二叉堆有效。由于它已经它的下方
		//已经是有效的二叉堆，所以只需要将它沉到合适的位置即可
		for index := startIndex; index < endIndex; index++ {
			h.down(index)
		}

		if startIndex == rootIndex {
			return
		}

		startIndex /= 2
		endIndex /= 2
	}
}