package main

// 奇数节点属于第一个链表，偶数节点属于第二个链表。原 linkedList 一定有偶数个节点
func splitLinkedList(linkedList *Node) (firstLinkedList *Node, secondLinkedList *Node) {
	if linkedList == nil {
		return
	}

	firstLinkedList = linkedList
	secondLinkedList = linkedList.next

	firstIndex := firstLinkedList
	secondIndex := secondLinkedList
	nextFirstIndex := secondIndex.next

	for nextFirstIndex != nil {
		firstIndex.next = nextFirstIndex
		secondIndex.next = nextFirstIndex.next

		firstIndex = firstIndex.next
		secondIndex = secondIndex.next
		nextFirstIndex = secondIndex.next
	}

	// 修复第一个链表最后一个节点
	firstIndex.next = nil

	return
}
