package exercises_day_1

func PositiveSum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}

	return sum
}
