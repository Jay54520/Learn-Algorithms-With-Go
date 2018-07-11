package main

// CopyLinkedList 题意见 https://www.geeksforgeeks.org/a-linked-list-with-next-and-arbit-pointer/
// 时间复杂度：
	// copyNextLinkedList: O(n)
	// for oldCurrent != nil: n
		// FindInNew(): n。所有 arbit 都是最后一个几点，所以要遍历 n 次找到 arbit
	// 所以这里是 O(n^2)
func CopyLinkedList(oldLinkedList *Node) (newLinkedList *Node) {

	newLinkedList = copyNextLinkedList(oldLinkedList)

	oldCurrent := oldLinkedList
	newCurrent := newLinkedList
	// 复制 arbit 链
	for oldCurrent != nil {
		newCurrent.arbit = FindInNew(newCurrent.arbit, oldLinkedList, newLinkedList)
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
	oldCurrent := oldLinkedList.next
	newPrevious := newLinkedList
	var newCurrent *Node

	for oldCurrent != nil {
		newCurrent = &Node{
			oldCurrent.value,
			nil,
			nil,
		}
		newPrevious.next = newCurrent

		oldCurrent = oldCurrent.next
		newPrevious = newCurrent
	}
	return
}

// FindInNew 找到旧链表的节点在新链表的对应节点
func FindInNew(oldNode *Node, oldLinkedList, newLinkedList *Node) (newNode *Node) {
	oldCurrent := oldLinkedList
	newCurrent := newLinkedList

	for oldCurrent != nil {
		if oldCurrent == oldNode {
			return newCurrent
		}
		newCurrent = newCurrent.next
		oldCurrent = oldCurrent.next
	}

	return
}
