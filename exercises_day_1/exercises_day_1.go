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

func MaxDiff(numbers []int) int {
	var min int
	var max int

	for i, number := range numbers {
		if i == 0 {
			min = number
			max = number
		}

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	return max - min
}

func RemDups(strings []string) (result []string) {
	for _, str := range strings {
		var strNoDups string
		var prevChar string

		for i, c := range str {
			if i == 0 {
				prevChar = string(c)
			}
			thisChar := string(c)

			if thisChar != prevChar || i == 0 {
				strNoDups += thisChar
			}

			prevChar = thisChar
		}

		result = append(result, strNoDups)
	}

	return result
}
