package main

// PrintMatrix 题意见 https://www.geeksforgeeks.org/print-a-given-matrix-in-spiral-form/
// 循环：依次向某个方向移动，并添加元素到结果中，遇到边界就转向，每次行动后，缩小下次行动的范围。方向的循环顺序为向右、向下、向左、向上。
// 终止：传入方向前面是边界
func PrintMatrix(matrix [][]int) (result []int) {
	startRow := 0
	endRow := len(matrix) - 1
	startColumn := 0
	endColumn := len(matrix[0]) - 1
	direction := "right"

	move(matrix, &result, startRow, endRow, startColumn, endColumn, direction)
	return
}

func move(matrix [][]int, result *[]int, startRow int, endRow int, startColumn int, endColumn int, direction string) {
	switch direction {
	case "right":
		if startColumn > endColumn {
			return
		}

		// 将 startRow 这一行的 [startColumn, endColumn] 元素从左向右添加到 result
		for _, value := range matrix[startRow][startColumn : endColumn+1] {
			*result = append(*result, value)
		}
		startRow += 1
		move(matrix, result, startRow, endRow, startColumn, endColumn, "down")

	case "down":
		if startRow > endRow {
			return
		}

		// 将 endColumn 这一列的所有元素添加到 result，endColumn -= 1，然后向左
		for _, row := range matrix[startRow : endRow+1] {
			*result = append(*result, row[endRow])
		}
		endColumn -= 1
		move(matrix, result, startRow, endRow, startColumn, endColumn, "left")

	case "left":
		// 将 endRow 这一行的所有元素从右到左添加到 result，endRow -= 1，然后向上
		endRowValues := matrix[endRow]
		if startColumn > endColumn {
			return
		}
		for i := endColumn; i >= startColumn; i-- {
			*result = append(*result, endRowValues[i])
		}
		endRow -= 1
		move(matrix, result, startRow, endRow, startColumn, endColumn, "up")

	case "up":
		if startRow > endRow {
			return
		}
		// 将 startColumn 这一列的所有元素从下到上添加到 result，startColumn += 1，然后向右
		for i := endRow; i >= startRow; i-- {
			row := matrix[i]
			*result = append(*result, row[startColumn])
		}
		startColumn += 1
		move(matrix, result, startRow, endRow, startColumn, endColumn, "right")

	}
}
