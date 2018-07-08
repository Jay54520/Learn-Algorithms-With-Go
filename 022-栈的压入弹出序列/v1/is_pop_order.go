package main

// IsPopOrder 比如 {1, 2, 3} 通过 push 1, push 2, pop 2, pop 1, push 3, pop3
// 可以得到 {2, 1, 3}
// 循环：对于 popOrder 中的当前元素，如果 stack 的栈顶元素不是它，那么从 pushOrder 中寻找
// 否则继续查看下一个 popOrder 元素是否符合条件
// 终止条件：1. 找完了 popOrder，返回 true
// 2. push 完所有的 pushOrder 还没找到，返回 false
// 3. 找到了，继续找下一个
func IsPopOrder(pushOrder []int, popOrder []int) (result bool) {
	var stack Stack
	popOrderIndex := 0
	pushOrderIndex := 0
	length := len(pushOrder)
	return isPopOrder(&stack, length, pushOrder, popOrder, popOrderIndex, pushOrderIndex)
}

func isPopOrder(stack *Stack, length int, pushOrder []int, popOrder []int,
	popOrderIndex int, pushOrderIndex int) (result bool) {

	// 找完了 popOrder
	if popOrderIndex == length {
		return true
	}

	popOrderCurrentValue := popOrder[popOrderIndex]
	stackTopValue := stack.Pop()
	stack.Push(stackTopValue)

	if popOrderCurrentValue != stackTopValue {
		// 找遍所有的 pushOrder 都没找到
		if pushOrderIndex == length {
			return
		}

		// 查看当前 pushOrderIndex 是否符合要求
		pushOrderCurrentValue := pushOrder[pushOrderIndex]
		// 符合要求，找下一个 popOrder
		if pushOrderCurrentValue == popOrderIndex {
			return isPopOrder(stack, length, pushOrder, popOrder, popOrderIndex+1, pushOrderIndex+1)
		} else { // 不符合要求，将当前 pushOrderCurrentValue 存入，再找下一个 pushOrderIndex
			stack.Push(pushOrderCurrentValue)
			return isPopOrder(stack, length, pushOrder, popOrder, popOrderIndex, pushOrderIndex+1)
		}

	} else {
		stack.Pop()
		return isPopOrder(stack, length, pushOrder, popOrder, popOrderIndex+1, pushOrderIndex)
	}
}
