package main

type Queue struct {
	Data []int
}

func (queue *Queue) enqueue(i int) {
	queue.Data = append(queue.Data, i)
}
