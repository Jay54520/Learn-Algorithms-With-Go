package main

// 一个 ints 的最大堆
type MaxIntHeap []int

// 需要实现 heap.Interface 中的接口：
// type Interface interface {
//	sort.Interface
//	Push(x interface{}) // add x as element Len()
//	Pop() interface{}   // remove and return element Len() - 1.
//}
//
// sort.Interface
//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}
// sort 的接口嵌套在 heap 的接口中，所以一共要实现 Len()、Less(), Swap
// Push()、Pop() 5 个接口

func (h MaxIntHeap) Len() int  {
	return len(h)
}

func (h MaxIntHeap) Less(i, j int) bool  {
	// i, j 表示元素的索引。这里的返回用于指出 i 是否应该排在 j 前面。
	// 由于我们这里是最大堆，所以如果 h[i] > h[j] 那么 i 应该排在 j 前面，
	// 否则， i 应该排在 j 后面。
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h MaxIntHeap) Swap(i, j int)  {
	h[i], h[j] = h[j], h[i]
}

// 将 num 添加到最后。由于要改变 h，所以要用指针
func (h *MaxIntHeap) Push(number interface{})  {
	*h = append(*h, number.(int))
}

// 删除并返回最后一个元素
func (h *MaxIntHeap) Pop() (result interface{})  {
	hValue := *h
	length := len(hValue)
	result = hValue[length - 1]
	*h = hValue[:length - 1]
	return 
}