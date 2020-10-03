package arrays

// Sum will take an array of numbers and return the total
func Sum(numbers []int) int {
	var result int

	for _, number := range numbers {
		result += number
	}

	return result
}

// SumAll will take a varying number of slices, returning a new slice containing the totals for each slice passed in
func SumAll(numbers ...[]int) []int {
	var sum []int

	for _, number := range numbers {
		sum = append(sum, Sum(number))
	}

	return sum
}

// SumAllTails calculates the totals of the "tails" of each slice.
// The tail of a collection is all the items apart from the first one (the "head")
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
