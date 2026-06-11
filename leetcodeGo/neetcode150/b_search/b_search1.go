package b_search

func search(nums []int, target int, startsFromIndex ...int) int {
	if startsFromIndex == nil {
		startsFromIndex = make([]int, 1)
		startsFromIndex[0] = 0
	}

	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	} else if numsLen == 1 {
		if nums[0] == target {
			return startsFromIndex[0]
		}
		return -1
	}

	midIndex := numsLen / 2
	if nums[midIndex] == target {
		return startsFromIndex[0] + midIndex
	} else if nums[midIndex] < target {
		return search(nums[midIndex:], target, startsFromIndex[0]+midIndex)
	}
	return search(nums[:midIndex], target, startsFromIndex[0])
}
