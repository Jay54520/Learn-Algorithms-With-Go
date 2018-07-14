package main

import "strings"

// Permutation 返回 s 的全排列
func Permutation(s string) (result []string) {
	permutation(strings.Split(s, ""), 0, len(s)-1, &result)
	return
}

// permutation
// 循环：每个元素加上剩余元素的全排列。等同于让每个元素（包含第一个元素）依次与第一个元素交换位置，然后求
// 第一个元素与剩余元素的全排列。通过开始、结束索引来求全排列
// 终止条件：只有一个字符串时，全排列就是它本身，这时全排列也完整了，所以将其加入结果列表。
// 中间状态的全排列储存在调用栈而不是变量中。
func permutation(stringSlice []string, startIndex int, endIndex int, result *[]string) {
	if startIndex == endIndex {
		*result = append(*result, strings.Join(stringSlice, ""))
	} else {
		for index := startIndex; index <= endIndex; index++ {
			// 交换位置
			stringSlice[startIndex], stringSlice[index] = stringSlice[index], stringSlice[startIndex]
			// 求全排列。通过 startIndex + 1 把 index 对应的元素放入全排列，然后求剩余元素的全排列
			permutation(stringSlice, startIndex+1, endIndex, result)
			// 恢复交换前的位置，为下一次交换作准备
			stringSlice[startIndex], stringSlice[index] = stringSlice[index], stringSlice[startIndex]
		}

	}

	return
}
