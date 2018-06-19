package main

func Search(array [][]int, number int) (result bool) {
	max_row_index := len(array) - 1
	max_col_index := len(array[0]) - 1

	row_index := 0
	col_index := max_col_index

	for row_index <= max_row_index && col_index >= 0 {
		topRightNumber := array[row_index][col_index]
		if number == topRightNumber {
			result = true
			break
		} else if number < topRightNumber {
			col_index -= 1
		} else {
			row_index += 1
		}
	}
	return
}
