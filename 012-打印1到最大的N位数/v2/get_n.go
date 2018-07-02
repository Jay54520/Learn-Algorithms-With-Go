package main

import (
	"errors"
	"fmt"
)

func GetN(n int) (numbers []int, err error) {
	if n <= 0 {
		return
	} else if n >= 19 {
		err = errors.New(fmt.Sprintf("%d is too big", n))
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
