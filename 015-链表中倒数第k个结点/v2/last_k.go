package main

// LastK 获取倒数第 K 个节点
// 相当于第 length - K + 1 个节点
func LastK(head *Node, k int) (result *Node) {
	length := 1
	indexNode := head
	for indexNode.next != nil {
		indexNode = indexNode.next
		length += 1
	}

	for i := 0; i < (length - k); i ++ {
		result = head.next
	}

	return
}

