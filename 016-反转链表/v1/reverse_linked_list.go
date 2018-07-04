package main

func ReverseLinkedList(head *Node) (result []int) {
	var tmpResult []int
	indexNode := head.next

	for indexNode != nil {
		tmpResult = append(tmpResult, indexNode.value)
		indexNode = indexNode.next
	}

	for i := len(tmpResult) - 1; i >= 0; i -- {
		result = append(result, tmpResult[i])
	}

	return
}
