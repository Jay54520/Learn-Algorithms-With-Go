package main

import "sort"

// 时间复杂度为 O( m * log(n) )，最坏情况是没有搜索到结果，空间复杂度为 O(m + n)，m、n 分别为 array 的行数、列数
func Search(array [][]int, number int) (result bool) {
	for _, row := range array {
		i := sort.Search(len(row), func(i int) bool { return row[i] >= number })
		if i < len(row) && row[i] == number {
			result = true
			break
		}
	}

	return
}
