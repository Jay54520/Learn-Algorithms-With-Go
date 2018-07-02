package main

func CalculateExponent(base float64, exponent int) float64 {
	var result float64 = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
