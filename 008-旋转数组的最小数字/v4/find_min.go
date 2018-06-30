package main

// FindMin 把一个数组最开始的若干个元素搬到数组的末尾，
// 我们称之为数组的旋转。输入一个非递减序列（元素都不同）的一个旋转，输出旋转数组的最小元素
func FindMin(array []int) int {
	// 如果旋转后的数组与原数组不同，那么旋转后的顺序为 递增 递减 递增，即
	// 结果在含有 "递减" 的区间内
	// 如果旋转后的数组与原数组相同，那么旋转后的数组的顺序为 “递增”
	// 所以采用类似二分搜索的递归搜索
	// 终止条件是：1. 发现了递减 2. 没有发现递减
	// 重复执行的代码是：在可能含有“递减”的区间内搜索
	return findMin(array, 0, len(array) - 1)
}

func findMin(array []int, low, high int) int {
	// 没有发现递减
	if low > high {
		return array[0]
	}

	middle := (high + low) / 2

	if array[middle] > array[middle+1] {
		return array[middle+1]
	}

	if array[middle] > array[high] {
		return findMin(array, middle + 1, high)
	} else {
		return findMin(array, low, middle - 1)
	}
}
