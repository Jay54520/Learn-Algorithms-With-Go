package main

func RotateArray(array []int, k int) []int {
	arrayLength := len(array)
	for i := 0; i < k; i ++ {
		array = append([]int{array[arrayLength - 1]}, array[:arrayLength-1]...)
	}
	return array
}
