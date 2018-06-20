package main

// 时间复杂度为 O( m + n)，空间复杂度为 O(m + n)，m、n 分别为 array 的行数、列数
func Search(array [][]int, number int) (result bool) {
	maxRowIndex := len(array) - 1
	maxColIndex := len(array[0]) - 1

	rowIndex := 0
	colIndex := maxColIndex

	for rowIndex <= maxRowIndex && colIndex >= 0 {
		topRightNumber := array[rowIndex][colIndex]
		if number == topRightNumber {
			result = true
			break
		} else if number < topRightNumber {
			colIndex -= 1
		} else {
			rowIndex += 1
		}
	}
	return
}
