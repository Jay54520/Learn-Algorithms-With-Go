package main

// 时间复杂度为 O( m * n)，空间复杂度为 O(m + n)，m、n 分别为 array 的行数、列数
func Search(array [][]int, number int) (result bool) {
	for _, row := range array {
		for _, n := range row {
			if number == n {
				result = true
			}
		}
	}

	return
}
