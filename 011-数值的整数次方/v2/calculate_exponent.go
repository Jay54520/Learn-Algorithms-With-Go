package main

import (
	"math"
	"errors"
)

func CalculateExponent(base float64, exponent int) (result float64, err error) {
	result = 1

	for i := 0; i < int(math.Abs(float64(exponent))); i++ {
		result *= base
	}

	if exponent < 0 {

		if base == 0 {
			err = errors.New("base 不能为 0 当 exponent 为负数")
		}
		result = 1 / result
	}

	return
}


// CalculateExponent2 使用快速幂解法：base^22 = base^16 * base^4 * base^2
// 22 的二进制为 10110，为 1 的部分分别为 2^4 = 16, 2^2=4, 2^1=2
// 所以 base^exponent 可以通过上面的方法计算
// base *= base，如果二进制为1，则将当前 base 乘入结果
// 时间复杂度是 O(log(exponent))，十进制的 exponent 转换为二进制就是 logn
// 解开这题的关键是要知道快速幂。我认为推导快速幂是不可能推导的出来的，只有靠背或者做题碰到过
// 参考 https://henrybear327.gitbooks.io/gitbook_tutorial/content/Algorithm/fast_pow/index.html
func CalculateExponent2(base float64, exponent int) (result float64, err error) {
	result = 1
	newExponent := int64(math.Abs(float64(exponent)))

	var i uint64 = 0

	for ; (1 << i) <= newExponent; i ++ {
		if (newExponent & (1 << i)) != 0 {
			result *= base
		}
		base = base * base
	}

	if exponent < 0 {

		if base == 0 {
			err = errors.New("base 不能为 0 当 exponent 为负数")
		}
		result = 1 / result
	}

	return
}
