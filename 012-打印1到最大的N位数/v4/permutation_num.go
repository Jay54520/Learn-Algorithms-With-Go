package main

import (
	"fmt"
	"strings"
	"strconv"
)

// PermutationNum 当前位数的全排列 + 剩余位数的全排列
// 终止条件：只有一位时，依次返回 0-9
// 代码循环：当前位数的全排列 + 剩余位数的全排列
// 由于是打印数字，所以可以是字符串类型，这样就避免了整数的范围问题
func PermutationNum(n int) (resultSlice []string) {
	stringSlice := make([]string, n)
	permutationNum(stringSlice, 0, n-1, &resultSlice)
	return
}
func permutationNum(stringSlice []string, startIndex int, endIndex int, resultSlice *[]string) {
	if startIndex == endIndex {
		for num := 0; num <= 9; num ++ {
			stringSlice[endIndex] = strconv.Itoa(num)
			result := fmt.Sprint(strings.Join(stringSlice, ""))
			*resultSlice = append(*resultSlice, result)
		}
	} else {
		for i := startIndex; i <= endIndex; i ++ {
			for num := 0; num <= 9; num ++ {
				stringSlice[startIndex] = strconv.Itoa(num)
				permutationNum(stringSlice, startIndex+1, endIndex, resultSlice)
			}
		}
	}
}
