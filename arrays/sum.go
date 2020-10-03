package arrays

func Sum(numbers []int) int {
	var result int

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbers ...[]int) []int {
	var sum []int

	for _, number := range numbers {
		sum = append(sum, Sum(number))
	}

	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sum []int

	for _, number := range numbers {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			tail := number[1:]
			sum = append(sum, Sum(tail))
		}
	}

	return sum
}
