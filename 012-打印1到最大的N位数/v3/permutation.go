package main

import (
	"fmt"
	"strings"
)

// Permutation 采用递归实现
// 递归循环：某一个元素加上剩余元素的全排列。如果让第一个元素与其他
// 剩余元素一次交换位置，那么就变成第一个元素加上剩余元素的全排列
// 递归终止条件：只有一个元素的全排列就是它本身，这时的 stringSlice 就是
// 一种排列结果
func Permutation(stringSlice []string) {
	permutation(stringSlice, 0, len(stringSlice)-1)
}

func permutation(stringSlice []string, startIndex, endIndex int) {
	if startIndex == endIndex {
		fmt.Println(strings.Join(stringSlice, ", "))
	} else {
		for i := startIndex; i <= endIndex; i ++ {
			// 依次将所有 [startIndex, endIndex] 与 startIndex 交换
			stringSlice[startIndex], stringSlice[i] = stringSlice[i], stringSlice[startIndex]
			// 剩余元素的全排列
			permutation(stringSlice, startIndex+1, endIndex)
			// 恢复初始位置，为下一次交换做准备
			stringSlice[startIndex], stringSlice[i] = stringSlice[i], stringSlice[startIndex]
		}
	}
}
