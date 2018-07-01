package main

func Fib(i int) int {
	if i == 0 || i == 1 {
		return i
	} else {
		return Fib(i-1) + Fib(i-2)
	}
}

func Fib2(i int) int {
	cache := make(map[int]int, i+1)
	return fib(i, cache)
}

func fib(i int, cache map[int]int) (result int) {
	// 由于很多是重复计算，所以使用 cache 缓存计算结果
	if result, ok := cache[i]; ok {
		return result
	}

	if i == 0 || i == 1 {
		result = i
	} else {
		result = fib(i-1, cache) + fib(i-2, cache)
	}
	cache[i] = result
	return
}
