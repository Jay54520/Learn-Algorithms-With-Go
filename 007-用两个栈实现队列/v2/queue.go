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
