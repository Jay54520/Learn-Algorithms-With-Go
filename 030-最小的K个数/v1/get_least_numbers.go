package main

// 输入n（n>=1）个整数，找出其中最小的k（1<=k<=n）个数。
// 例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。
//
// 算法1：首先最简单的是调用语言自带的排序函数对 numbers 排序，然后
// 返回排序后数组中的前 k 个数字。以归并排序为例，时间复杂度为 O(nlogn)，
// 空间复杂度为 O(n)。
//
// 算法2：上面的排序是针对所有的 n 个数字，但是这里只要求 k 个数字，
// 假如 k 为 1，那么相当于找出最小的数字，时间复杂度为 O(n)，空间复杂度为 O(1)，
// 所以空间复杂度一定是可以优化的，不可能你找一个最小的数字还要储存 n 个元素吧？
// 由于 k 不一定等于 1，所以我们需要从储存一个最小元素的变量
// 扩展成维护一个容量为 k+1 的有序数组：将数字插入到使有序数组依旧有序的位置，
// 如果加入后有序数组的长度等于 k + 1，那么删除当前最大元素即最后一个元素。
// 使用容量为 k + 1 而不是 k 的数组可以简化代码的写法，因为不用先删除再加入，
// 而是先加入再删除，和有序数组没有满时相比只是多了一个删除的操作。
// 遍历完成后，返回这个有序数组。
//
// 复杂度分析：
// 时间复杂度：遍历数组为 O(n)、 维护容量为 k 的有序数组为 O(logk)（二分法查找位置然后插入）
// 所以为 O(nlogk)。
// 空间复杂度： O(k)，用于储存 k 个有序元素。
// 相比算法 1 有优化。
func GetLeastNumbers(numbers []int, k int) []int {
	sortedSlice := make([]int, 0, k+1)
	var length int

	for _, number := range numbers {
		InsertToSortedSlice(&sortedSlice, number)
		length = len(sortedSlice)
		if length == k+1 {
			sortedSlice = sortedSlice[:length-1]
		}
	}
	return sortedSlice
}
