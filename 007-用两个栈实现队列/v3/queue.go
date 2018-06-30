package main

type Queue struct {
	Data []int
}

func (queue *Queue) enqueue(i int) {
	queue.Data = append(queue.Data, i)
}

func (queue *Queue) dequeue() (result int) {
	if len(queue.Data) == 0 {
		return
	}

	result = queue.Data[0]
	newStartIndex := 1
	if newStartIndex == len(queue.Data) {
		queue.Data = []int{}
	} else {
		queue.Data = queue.Data[newStartIndex:]
	}

	return

}

type StackQueue struct {
	DataStack *Stack
	stack1    *Stack
}

func (stackQueue *StackQueue) enqueue(i int) {
	if stackQueue.DataStack == nil {
		stackQueue.DataStack = &Stack{}
	}
	stackQueue.DataStack.Push(i)
}
func (stackQueue *StackQueue) dequeue() (result int) {
	i := stackQueue.DataStack.Pop()
	for i != 0 {
		stackQueue.stack1.Push(i)
		i = stackQueue.DataStack.Pop()
	}
	result = stackQueue.stack1.Pop()

	stackQueue.DataStack, stackQueue.stack1 = stackQueue.stack1, stackQueue.DataStack

	return
}
