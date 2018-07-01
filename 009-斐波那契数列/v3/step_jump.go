package main


// StepJump 一次可以跳上 1 或 2 级台阶，返回跳上 n 级台阶的方法个数
func StepJump(i int) int {
	if i == 1 {
		// 1 级台阶只有一种跳法
		return 1
	} else if i == 2 {
		// 1 级台阶有两种跳法： 1 + 1 或 2
		return 2
	} else {
		// 跳了 2 级台阶后的方法个数 + 跳了 1 级台阶后的方法个数
		return StepJump(i - 2) + StepJump(i - 1)
	}
}
