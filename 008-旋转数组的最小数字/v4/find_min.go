package main

// FindMin 把一个数组最开始的若干个元素搬到数组的末尾，
// 我们称之为数组的旋转。输入一个非递减序列（元素都不同）的一个旋转，输出旋转数组的最小元素
func FindMin(array []int) int {
	// 由于数组是非递减序列，所以旋转前的数组的第一个就是最小元素
	// 旋转后的数组是一个有序的数组，除了最大元素到最小元素这个位置是递减的
	// 所以采用类似二分搜索的递归搜索
	// 终止条件是：1. 发现了递减 2. 发现最大值是最后一个
	// 重复执行的代码是：在 array 的 [low, high) 中搜索
	return findMin(array, 0, len(array))
}

func findMin(array []int, low, high int) int {
	middle := (high + low) / 2

	if middle + 1 == high {
		// 最大值是最后一个，那么最小值是第一个
		return array[0]
	}

	if array[middle] > array[middle+1] {
		return array[middle+1]
	} else {
		return findMin(array, middle, high)
	}
}
