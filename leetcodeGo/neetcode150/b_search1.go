package main

func search(nums []int, target int) int {
	numsLen := len(nums)
	index := numsLen / 2
	prevIndex := -1

	for index != prevIndex {
		if nums[index] == target {
			return index + 1
		}
		
		prevIndex = index
		if nums[index] > target {
			numsLen = index - 1
			index = numsLen / 2
		} else {
			numsLen = numsLen - index
			index = numsLen / 2
		}

	}
	return -1
}
