package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sumOfAll := []int{}
	sum := 0
	for _, numberSlice := range numbersToSum {
		for _, num := range numberSlice {
			sum += num
		}
		sumOfAll = append(sumOfAll, sum)
		sum = 0
	}
	return sumOfAll
}
