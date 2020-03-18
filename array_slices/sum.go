package array_slices

func SumArray(numbers [5]int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func SumSlice(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	length := len(numbersToSum)
	sums = make([]int, length)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := []int{0}
		if len(numbers) > 0 {
			tail = numbers[1:]
		}
		sums = append(sums, SumSlice(tail))
	}
	return sums
}