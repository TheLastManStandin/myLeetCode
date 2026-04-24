package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if index, ok := hash[diff]; ok {
			return []int{index, i}
		}
		_, ok := hash[num]
		if !ok {
			hash[num] = i
		}
	}
	return nil
}
