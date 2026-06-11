package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	sliseMap := [][]int{}
	ans := []int{}
	for _, num := range nums {
		_, ok := hash[num]
		if ok {
			hash[num] = hash[num] + 1
		} else {
			hash[num] = 1
		}
	}

	for key, element := range hash {
		sliseMap = append(sliseMap, []int{key, element})
	}

	sort.Slice(sliseMap, func(i, j int) bool {
		return sliseMap[i][1] > sliseMap[j][1]
	})

	for i := 0; i < k; i++ {
		ans = append(ans, sliseMap[i][0])
	}
	return ans
}
