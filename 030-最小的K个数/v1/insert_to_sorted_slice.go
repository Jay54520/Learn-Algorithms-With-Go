package main

import "sort"

// InsertToSortedSlice 将 number 插入 sortedSLice 并使插入后的
// sortedSLice 仍然有序
// 使用二分搜索在 sortedSLice 中查找大于等于 number 的第一个索引，
// 如果找到了，将 number 插入到这个索引对应的数字前面；
// 否则，说明 number 是大于 sortedSLice 中的所有数，将 number 插入到最后
func InsertToSortedSlice(sortedSLice *[]int, number int) {
	sortedSliceValue := *sortedSLice
	var index int

	if len(sortedSliceValue) == 0 {
		sortedSliceValue = append(sortedSliceValue, number)
		*sortedSLice = sortedSliceValue
		return
	} else {
		index = sort.Search(number, func(i int) bool {
			return sortedSliceValue[i] >= number
		})
	}

	// 没有找到大于等于 number 的数
	if index == -1 {
		*sortedSLice = append(sortedSliceValue, number)
	} else {
		// ------------将 number 插入到 index 这个位置------------
		// 参考 https://gist.github.com/zhum/57cb45d8bbea86d87490
		// 扩大 cap，将 index 和之后的所有元素复制到 index + 1 之后，然后替换 index 的值
		if len(sortedSliceValue) == cap(sortedSliceValue) {
			sortedSliceValue = append(sortedSliceValue, 0)
		}
		copy(sortedSliceValue[index+1:], sortedSliceValue[index:])
		sortedSliceValue[index] = number
		// ------------将 number 插入到 index 这个位置------------

		*sortedSLice = sortedSliceValue
	}
}
