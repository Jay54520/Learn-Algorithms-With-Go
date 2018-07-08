package main

type Stack struct {
	Data []int
	minimumIndex int
	length int
}

func (stack *Stack) Push(i int) {
	stack.Data = append(stack.Data, i)
	stack.length += 1

	currentIndex := stack.length - 1
	// 这行代码要位于 append 下面，因为当 stack.Data 为空时会
	// 超出 stack.Data 的长度
	if i < stack.Data[stack.minimumIndex] {
		stack.minimumIndex = currentIndex
	}

}

func (stack *Stack) Pop() (result int) {
	if len(stack.Data) == 0 {
		return
	}

	result = stack.Data[len(stack.Data)-1]
	newEndIndex := len(stack.Data) - 1
	if newEndIndex < 0 {
		newEndIndex = 0
	}
	stack.Data = stack.Data[:newEndIndex]
	return
}

// GetMin 题意见 https://www.geeksforgeeks.org/design-a-stack-that-supports-getmin-in-o1-time-and-o1-extra-space/
func (stack Stack) GetMin() (minimum int) {
	if stack.length == 0 {
		return -1
	}
	return stack.Data[stack.minimumIndex]
}
