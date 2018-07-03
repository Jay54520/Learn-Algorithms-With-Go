package main

// LastK 获取倒数第 K 个节点
// 相当于第 length - K + 1 个节点
// 采用双指针法，减少第二次遍历
// 第一个指针right先向前走K步，然后left和right一起走，
// 此时两个指针差别K步，那么当right走到链表尾部的时候，
// left指向的就是倒数第K个节点
func LastK(head *Node, k int) (result *Node) {
	leftNode := head
	rightNode := head

	for i := 0; i < k; i ++ {
		// 链表的长度小于 k
		if rightNode == nil {
			return
		}
		rightNode = rightNode.next
	}

	for {
		// rightNode 到达了第一个空节点
		// 不能写在 rightNode.next 后面
		// 因为当链表长度为 k 时，rightNode 为 nil
		// 造成 nil pointer dereference
		if rightNode == nil {
			return leftNode
		}

		leftNode = leftNode.next
		rightNode = rightNode.next
	}
	return
}
