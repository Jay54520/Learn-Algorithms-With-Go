package main

func ListLinkedListValue(node *Node) []int {
	var ascendingValueList []int

	for node.next != nil {
		ascendingValueList = append(ascendingValueList, node.value)
		node = node.next
	}
	ascendingValueList = append(ascendingValueList, node.value)

	descendingValueList := reversed(ascendingValueList)

	return descendingValueList
}

func reversed(list []int) (reversedList []int) {
	for i := len(list) - 1; i >= 0; i -- {
		reversedList = append(reversedList, list[i])
	}
	return
}
