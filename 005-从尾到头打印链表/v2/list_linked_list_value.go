package main

func ListLinkedListValue(node *Node) []int  {
	ascendingValueList := []int{}
	descendingValueList := []int{}
	for node.next != nil {
		ascendingValueList = append(ascendingValueList, node.value)
		node = node.next
	}
	ascendingValueList = append(ascendingValueList, node.value)

	for i:= len(ascendingValueList) - 1; i >= 0; i -- {
		descendingValueList = append(descendingValueList, ascendingValueList[i])
	}

	return descendingValueList
}
