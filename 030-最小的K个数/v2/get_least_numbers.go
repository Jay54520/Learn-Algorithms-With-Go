package main

import "container/heap"

// 输入n（n>=1）个整数，找出其中最小的k（1<=k<=n）个数。
// 例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。
//
// 算法3：相对算法 2 维护一个有序的数组，这里使用最大堆。插入元素到最大堆，并保持堆有序，如果此时
// 最大堆的长度为 k + 1，则删除最大元素。
// 返回时，pop 出最大堆中的所有元素（此时是由大到小），然后逆序返回它，因为结果要求由小到大。
//
// 复杂度分析：
// 时间复杂度：遍历数组为 O(n)、 创建堆为 O(k)，维护容量为 k 的最大堆为 O((n - k)* logk)
// 所以为 O(nlogk)。k 相比于 n 可以被忽略掉。
// 空间复杂度： O(k)，用于储存 k 个元素的最大堆。
// 相比算法 2 时间复杂度有优化。
func GetLeastNumbers(numbers []int, k int) (result []int) {
	hValue := make(MaxIntHeap, 0, k + 1)
	h := &hValue
	heap.Init(h)

	for _, number := range numbers {
		heap.Push(h, number)
		if h.Len() == k + 1{
			heap.Pop(h)
		}
	}
	var reversedResult []int
	for h.Len() > 0 {
		reversedResult = append(reversedResult, heap.Pop(h).(int))
	}

	for i := len(reversedResult) - 1; i >= 0; i -- {
		result = append(result, reversedResult[i])
	}
	return
}
