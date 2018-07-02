package main

func GetN(n int) (numbers []int, err error) {
	if n <= 0 {
		return
	}

	maxN := 10
	for i := 0; i < n; i ++ {
		maxN *= 10
	}
	numbers = make([]int, maxN-1)
	for i := 1; i < maxN; i++ {
		numbers = append(numbers, i)
	}
	return
}
