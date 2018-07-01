package main

// NumberOfOne 输出输入参数二进制表示的 1 的个数
// flag 只有一位是 1，其他位是 0，所以如果 i 的对应为不是 1，
// 那么 i & flag 为 0
// flag 越界后变为 0，说明遍历完了
func NumberOfOne(i int64) (count int) {
	var flag int64 = 1

	for flag != 0 {
		if i&flag != 0 {
			count += 1
		}
		flag <<= 1
	}

	return
}
