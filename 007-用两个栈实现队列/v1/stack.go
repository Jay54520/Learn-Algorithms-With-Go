package main

type Stack struct {
	Data []int
}

func (stack *Stack) Push(i int) {
	stack.Data = append(stack.Data, i)
}

