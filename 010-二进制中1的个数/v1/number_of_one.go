package main

import (
	"strconv"
	"strings"
)

const BINARY = 2

// NumberOfOne 输出输入参数二进制表示的 1 的个数
func NumberOfOne(i int64) int {
	binaryString := strconv.FormatInt(i, BINARY)
	return strings.Count(binaryString, "1")
}

