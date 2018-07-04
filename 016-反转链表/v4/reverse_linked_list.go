package main

// ReverseLinkedList 反转链表
// 使用三指针 prev, current, next，current.next = prev
// 然后再变换三指针
// 之后处理特殊情况
// 使用递归的写法
func ReverseLinkedList(head *Node) (result *Node) {
	var prev *Node = nil
	current := head
	return reverseLinkedList(prev, current)
}

func reverseLinkedList(prev, current *Node) *Node {
	if current == nil {
		return prev
	} else {
		next := current.next
		current.next = prev
		prev = current
		current = next
		return reverseLinkedList(prev, current)
	}
}
