package main

// MoreThanHalfNum 数组中有一个数字出现的次数可能超过数组长度的一半
// （指数组的长度除以 2 的整数部分），请找出这个数字。
// 例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
// 由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
// 如果不存在则输出0，如果 0 超过一半也输出 0。
//
// 算法：维护一个数字出现次数的字典，如果数字不存在于字典，则设置它的出现次数为 1，
// 否则给它的出现次数加 1。
// 每次设置后，检验次数，如果大于等于数组长度的一半则返回这个数字。
// 遍历完成数组后返回 0，因为如果代码执行到了遍历完成后，说明没有
// 这样的数字
func MoreThanHalfNum(nums []int) (result int) {
	countOfNum := make(map[int]int)
	var count int
	var existence bool
	halfLength := len(nums) / 2

	for _, num := range nums {
		if count, existence = countOfNum[num]; existence {
			count += 1
		} else {
			count = 1
		}

		if count >= halfLength {
			return num
		}

		countOfNum[num] = count
	}

	return
}

