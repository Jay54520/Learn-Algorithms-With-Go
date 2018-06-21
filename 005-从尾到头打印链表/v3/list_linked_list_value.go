package main

func ListLinkedListValue(node *Node) []interface{} {
	var ascendingValueList []interface{}

	for node.next != nil {
		ascendingValueList = append(ascendingValueList, node.value)
		node = node.next
	}
	ascendingValueList = append(ascendingValueList, node.value)

	descendingValueList := reversed(ascendingValueList)

	return descendingValueList
}

func reversed(list []interface{}) (reversedList []interface{}) {
	for i := len(list) - 1; i >= 0; i -- {
		reversedList = append(reversedList, list[i])
	}
	return
}
