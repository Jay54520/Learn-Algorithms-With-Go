package main

import "sort"

// InsertToSortedSlice 将 number 插入 sortedSLice 并使插入后的
// sortedSLice 仍然有序
// 使用sort.SearchInts（二分搜索）查找插入的索引，将 number 插入到这个索引位置
//
// 时间复杂度：二分搜索 O(logn)、插入到开头最多要拷贝 n 次，为 O(n)，所以为 O(n)
// 空间复杂度：O(1)
func InsertToSortedSlice(sortedSLice *[]int, number int) {
	sortedSliceValue := *sortedSLice
	var index int

	if len(sortedSliceValue) == 0 {
		sortedSliceValue = append(sortedSliceValue, number)
		*sortedSLice = sortedSliceValue
		return
	} else {
		index = sort.SearchInts(sortedSliceValue, number)
	}

	// ------------将 number 插入到 index 这个位置------------
	// 参考 https://gist.github.com/zhum/57cb45d8bbea86d87490
	// 扩大 len，将 index 和之后的所有元素复制到 index + 1 之后，然后替换 index 的值
	sortedSliceValue = append(sortedSliceValue, 0)
	copy(sortedSliceValue[index+1:], sortedSliceValue[index:])
	sortedSliceValue[index] = number
	// ------------将 number 插入到 index 这个位置------------

	*sortedSLice = sortedSliceValue
}
