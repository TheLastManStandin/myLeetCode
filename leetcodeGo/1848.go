package main

import "fmt"

func getMinDistance(nums []int, target int, start int) int {
	for i := 0; ; i++ {
		if start+i < len(nums) && nums[start+i] == target {
			return i
		} else if start-i >= 0 && nums[start-i] == target {
			return i
		}
	}
}

func main() {
	fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
}
