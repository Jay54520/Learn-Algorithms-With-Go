package main

// SortedMerge()  takes two lists, each of which is sorted in increasing order,
// and merges the two together into one list which is in increasing order.
// SortedMerge() should return the new list.
// The new list should be made by splicing
// together the nodes of the first two lists.
// For example if the first linked list a is
// 5->10->15 and the other linked list b is 2->3->20,
// then SortedMerge() should return a pointer to
// the head node of the merged list 2->3->5->10->15->20.
// 比较两个链表的节点，当前节点为最小的节点，然后变动当前节点为当前
// 节点的下一个节点
// 设置过后的 Index 向后移动一位
func SortedMerge(l1 *Node, l2 *Node) (result *Node) {
	l1Index := l1
	l2Index := l2
	var current *Node = nil

	for (l1Index != nil) || (l2Index != nil) {
		minNode := getMinNode(l1Index, l2Index)

		// 第一个 Node —— 链表头
		if current == nil {
			result = minNode
			current = minNode
		} else {
			current.next = minNode
			current = current.next
		}

		if minNode == l1Index {
			l1Index = l1Index.next
		} else {
			l2Index = l2Index.next
		}
	}
	return
}

func getMinNode(node1 *Node, node2 *Node) *Node {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	} else {
		if node1.value <= node2.value {
			return node1
		} else {
			return node2
		}
	}
}
