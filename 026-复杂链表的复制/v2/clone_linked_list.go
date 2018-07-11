package main

var oldToNew = make(map[*Node]*Node)

// CopyLinkedList 题意见 https://www.geeksforgeeks.org/a-linked-list-with-next-and-arbit-pointer/
// 时间复杂度：
	// copyNextLinkedList: O(n)
	// for oldCurrent != nil: n
		// FindInNew(): 1。通过 map 查找
	// 所以这里是 O(n)
func CopyLinkedList(oldLinkedList *Node) (newLinkedList *Node) {

	newLinkedList = copyNextLinkedList(oldLinkedList)

	oldCurrent := oldLinkedList
	newCurrent := newLinkedList
	// 复制 arbit 链
	for oldCurrent != nil {
		newCurrent.arbit = FindInNew(newCurrent.arbit)
		newCurrent = newCurrent.next
		oldCurrent = oldCurrent.next
	}

	return
}

// copyNextLinkedList 拷贝 oldLinkedList 的 next 链到 newLinkedList
// 复杂度：O(n) n 为 oldLinkedList 的节点个数
func copyNextLinkedList(oldLinkedList *Node) (newLinkedList *Node) {

	newLinkedList = &Node{
		oldLinkedList.value,
		nil,
		nil,
	}
	oldToNew[oldLinkedList] = newLinkedList

	oldCurrent := oldLinkedList.next
	newPrevious := newLinkedList
	var newCurrent *Node

	for oldCurrent != nil {
		newCurrent = &Node{
			oldCurrent.value,
			nil,
			nil,
		}
		oldToNew[oldCurrent] = newCurrent
		newPrevious.next = newCurrent

		oldCurrent = oldCurrent.next
		newPrevious = newCurrent
	}
	return
}

// FindInNew 找到旧链表的节点在新链表的对应节点
func FindInNew(oldNode *Node) (newNode *Node) {
	return oldToNew[oldNode]
}
