package main

type MaxCountNum struct {
	Num           int
	RelativeCount int
}

// MoreThanHalfNum 数组中有一个数字出现的次数可能超过数组长度的一半
// （指数组的长度除以 2 的整数部分），请找出这个数字。
// 例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
// 由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
// 如果不存在则输出0，如果 0 超过一半也输出 0。
//
// 算法：v1 版本需要维护一个键的个数最大为数组的元素个数的字典，这里尝试优化这部分。
// 只记录当前出现次数相对最多的数字（次数是相对的、不准确的，数字是准确的）。
// 如果相对次数为 0，那么设置当前的数字为 MaxCountNum
// 如果当前数字与 MaxCountNum.Num 相同，则相对次数加一，否则，相对次数减一
// 遍历完数组后，如果相对次数大于数组长度的一半，那么这个数字就是要找的数
// 如果相对次数大于 0，那么它可能是要找的数字，所以再遍历一遍数组来检验
// 如果等于 0，说明不存在这个数字
// 算法参考 『阵地攻守』 of https://github.com/gatieme/CodingInterviews/tree/master/029-%E6%95%B0%E7%BB%84%E4%B8%AD%E5%87%BA%E7%8E%B0%E6%AC%A1%E6%95%B0%E8%B6%85%E8%BF%87%E4%B8%80%E5%8D%8A%E7%9A%84%E6%95%B0%E5%AD%97
// 从这里可以学习到，有时候将低对准确性的要求可以大大提升算法的效率，比如从准确的出现次数
// 变为相对的出现次数。
//
// 复杂度分析：
// 时间复杂度：没有这个数字，需要遍历两次整个数组，所以为 O(n) + O(n) = O(n)
// 空间复杂度：只需要 MaxCountNum 这一个额外储存，所以为 O(1)
func MoreThanHalfNum(nums []int) (result int) {
	halfLength := len(nums) / 2
	var maxCountNum MaxCountNum

	for _, num := range nums {
		if maxCountNum.RelativeCount == 0 {
			maxCountNum.Num = num
			maxCountNum.RelativeCount = 1
		} else {
			if num == maxCountNum.Num {
				maxCountNum.RelativeCount += 1
			} else {
				maxCountNum.RelativeCount -= 1
			}
		}
	}

	relativeCount := maxCountNum.RelativeCount
	switch {
	case relativeCount > halfLength:
		return maxCountNum.Num
	case relativeCount > 0:
		count := 0
		for _, num := range nums {
			if num == maxCountNum.Num {
				count += 1
			}
		}
		if count > halfLength {
			return maxCountNum.Num
		}
	}

	return
}
