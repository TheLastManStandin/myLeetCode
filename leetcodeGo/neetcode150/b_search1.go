package main

func search(nums []int, target int) int {
	numsLen := len(nums) - 1
	index := numsLen / 2
	//prevIndex := -1

	for numsLen > 0 {
		if nums[index] == target {
			return index
		}

		//prevIndex = index
		if nums[index] > target {
			numsLen /= 2      // 1 2 3 4 -> nl = 0		1 2 3  -> nl = 1
			index = index / 2 //   i-1   -> i = 0           i    -> i = 0
		} else {
			numsLen /= 2 // 1 2 3 4 -> nl =
			index = index + numsLen/2 + 1
		}

	}
	return -1
}
