package main

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
