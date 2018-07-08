package main

// IsPopOrder 比如 {1, 2, 3} 通过 push 1, push 2, pop 2, pop 1, push 3, pop3
// 可以得到 {2, 1, 3}
// popOrder 中的相邻元素必须在处理后的 pushOrder 中相邻，因为栈只能通过 push、pop 操作
// 循环：判断 popOrder 中的当前元素与下一个元素的位置关系
// 终止条件：1. 判断完所有的 popOrder，返回 true
// 2. 下一个元素与当前元素在处理后的 pushOrder 中不相邻，返回 false
// 3. 继续找下一个
func IsPopOrder(pushOrder []int, popOrder []int) (result bool) {
	popOrderIndex := 0
	length := len(pushOrder)
	return isPopOrder(length, pushOrder, popOrder, popOrderIndex)
}

func isPopOrder(length int, pushOrder []int, popOrder []int,
	popOrderIndex int) (result bool) {

	// 找完了 popOrder
	if popOrderIndex+1 >= length {
		return true
	}

	// 不相邻，返回 false
	if !isAdjacent(pushOrder, popOrder[popOrderIndex], popOrder[popOrderIndex+1]) {
		return
	} else {
		pushOrder = pop(pushOrder, index(pushOrder, popOrder[popOrderIndex]))
		return isPopOrder(length, pushOrder, popOrder, popOrderIndex+1)
	}
}

// index
func index(intSlice []int, findValue int) int {
	for index, value := range intSlice {
		if value == findValue {
			return index
		}
	}
	return -1
}

// pop pop index at intSlice
func pop(intSlice []int, index int) []int {
	return append(intSlice[:index], intSlice[index+1:]...)
}

// isAdjacent num1 和 num2 是否在 ints 中相邻，ints 中的数都不相同
func isAdjacent(numSlice []int, num1 int, num2 int) (result bool) {

	num1Index := -1
	num2Index := -1
	for index, value := range numSlice {
		if value == num1 {
			num1Index = index
		} else if value == num2 {
			num2Index = index
		}
	}

	// 存在数不存在于 numSlice
	if num1Index == -1 || num2Index == -1 {
		return
	}

	diff := num1Index - num2Index
	if diff == -1 || diff == 1 {
		return true
	} else {
		return
	}
}
