package main

import (
	"fmt"
	"strings"
)

// Permutation 返回 stringSlice 的全排列
func Permutation(stringSlice []string) (resultSlice []string) {
	permutation(stringSlice, 0, len(stringSlice)-1, &resultSlice)
	return
}

// permutation 采用递归实现
// 递归循环：当前序列的全排列 = 某一个元素加上剩余元素的全排列。如果让第一个元素与其他
// 剩余元素依次交换位置，那么全排列就 =（交换后的）第一个元素加上剩余元素的全排列
// 递归终止条件：为当前序列中只有一个元素，因为一个元素的全排列就是它本身。这时的 stringSlice 就是
// 一组完整的排列结果
func permutation(stringSlice []string, startIndex, endIndex int, resultSlice *[]string)  {
	if startIndex == endIndex {
		result := fmt.Sprint(strings.Join(stringSlice, ""))
		*resultSlice = append(*resultSlice, result)
	} else {
		for i := startIndex; i <= endIndex; i ++ {
			// 依次将所有 [startIndex, endIndex] 与 startIndex 交换
			stringSlice[startIndex], stringSlice[i] = stringSlice[i], stringSlice[startIndex]
			// 剩余元素的全排列
			permutation(stringSlice, startIndex+1, endIndex, resultSlice)
			// 恢复初始位置，为下一次交换做准备
			stringSlice[startIndex], stringSlice[i] = stringSlice[i], stringSlice[startIndex]
		}
	}
}
