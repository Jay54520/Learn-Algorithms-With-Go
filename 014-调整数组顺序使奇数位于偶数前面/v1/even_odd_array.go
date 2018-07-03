package main

func EvenOddArray(numbers []int) (result []int) {
	var evenNumbers []int
	var oddNumbers []int

	for _, v := range numbers {
		if v%2 == 0 {
			oddNumbers = append(oddNumbers, v)
		} else {
			evenNumbers = append(evenNumbers, v)
		}
	}

	result = append(evenNumbers, oddNumbers...)
	return
}
