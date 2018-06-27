package main

type Stack struct {
	Data []int
}

func (stack *Stack) Push(i int) {
	stack.Data = append(stack.Data, i)
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
