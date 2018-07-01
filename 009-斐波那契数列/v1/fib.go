package main

func Fib(i int) int {
	if i == 0 || i == 1 {
		return i
	} else {
		return Fib(i - 1) + Fib(i - 2)
	}
}

