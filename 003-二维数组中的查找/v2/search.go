package main

import "sort"

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
