package b_search

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	lastMid := nums[right]

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > lastMid {
			lastMid = mid
			left = mid + 1
		} else if nums[mid] < lastMid {
			right = mid
		}
	}
	return nums[right]
}
