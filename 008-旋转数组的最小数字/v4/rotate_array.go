package main

func RotateArray(array []int, k int) []int {
	arrayLength := len(array)

	if arrayLength == 0 {
		return array
	}

	k = k % arrayLength

	array = append(array[arrayLength - k:], array[:k + 1]...)
	return array
}
