package main

// CopyLinkedList 题意见 https://www.geeksforgeeks.org/a-linked-list-with-next-and-arbit-pointer/
// 之前的 FindNew 用来根据新旧节点的关系来寻找，这里通过将新节点插入到旧节点后面
// 来保存这种关系而避免再构建一个 map，这样就能将空间复杂度降为 O(1)
// 具体来说，
// 1. 将新节点插入到旧节点后面
// 2. 设置新节点的 arbit（为旧节点的 arbit 后面那个节点，因为新节点在旧节点后面）
// 3. 分离出新节点
// 时间复杂度：
// copyNextLinkedList: O(n)
// for oldCurrent != nil: n
// FindInNew(): 1。通过 map 查找
// 所以这里是 O(n)
func CopyLinkedList(oldLinkedList *Node) (newLinkedList *Node) {

	if oldLinkedList == nil {
		return
	}

	oldCurrent := oldLinkedList
	var newNode *Node
	var nextOld *Node

	// ----------- 将新节点插入到旧节点后面 ------------
	for oldCurrent != nil {
		nextOld = oldCurrent.next
		newNode = &Node{
			oldCurrent.value,
			nextOld,
			nil,
		}
		oldCurrent.next = newNode
		oldCurrent = nextOld
	}
	// ----------- 将新节点插入到旧节点后面 ------------

	// ----------- 设置新节点的 arbit ------------
	oldCurrent = oldLinkedList
	for oldCurrent != nil {
		newNode = oldCurrent.next
		if oldCurrent.arbit != nil {
			newNode.arbit = oldCurrent.arbit.next
		}
		oldCurrent = newNode.next
	}
	// ----------- 设置新节点的 arbit ------------

	// 分离出 oldLinkedList 和 newLinkedList
	oldLinkedList, newLinkedList = splitLinkedList(oldLinkedList)
	return
}
