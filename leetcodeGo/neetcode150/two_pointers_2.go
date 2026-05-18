package main

func twoSum(numbers []int, target int) []int {
	indexOne := 0
	indexTwo := len(numbers) - 1

	for indexOne < indexTwo {
		summ := numbers[indexOne] + numbers[indexTwo]

		if summ < target {
			indexOne++
		} else if summ > target {
			indexTwo--
		} else {
			return []int{indexOne + 1, indexTwo + 1}
		}
	}

	return []int{0, 0}
}
