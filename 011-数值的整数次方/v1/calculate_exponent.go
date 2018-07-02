package main

import (
	"math"
	"errors"
)

func CalculateExponent(base float64, exponent int) (result float64,  err error) {
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
