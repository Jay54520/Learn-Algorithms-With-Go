package main

func RotateArray(array []int, k int) []int {
	arrayLength := len(array)

	if arrayLength == 0 {
		return array
	}

	k = k % arrayLength

	for i := 0; i < k; i ++ {
		array = append([]int{array[arrayLength - 1]}, array[:arrayLength-1]...)
	}
	return array
}
