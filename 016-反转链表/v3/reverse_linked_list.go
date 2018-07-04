package main

// ReverseLinkedList 反转链表
// 使用三指针 prev, current, next，current.next = prev
// 然后再变换三指针
// 之后处理特殊情况
func ReverseLinkedList(head *Node) (result *Node) {
	if head == nil {
		return
	}

	var prev *Node = nil
	current := head

	for current.next != nil {
		next := current.next

		// 当前节点的下一个节点为前一个节点
		current.next = prev
		prev = current
		current = next
	}

	// current 是最后一个节点
	current.next = prev

	result = current
	return
}
