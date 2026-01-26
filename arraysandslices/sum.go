package main

func Sum(numbers []int) int {

	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func Sum(number []int) int {
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int

	for _, slice := range tailsToSum {
		if len(slice) != 0 {
			sums = append(sums, Sum(slice[1:]))
		}
	}

	return sums
}
