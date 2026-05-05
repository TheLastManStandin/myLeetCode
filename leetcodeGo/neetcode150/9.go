package main

func longestConsecutive(nums []int) int {
	hash := map[int]int{}
	maxx := 1
	if len(nums) < 1 {
		return 0
	}

	for _, n := range nums {
		if _, ok := hash[n]; ok {
			continue
		}
		hash[n] = 1

		toAdd := 1

		if _, ok := hash[n-1]; ok {
			hash[n] = hash[n-1] + toAdd
			if maxx < hash[n] {
				maxx = hash[n]
			}
			toAdd = hash[n]
		}

		if _, ok := hash[n+1]; ok {
			for j := 1; j < len(nums); j++ {
				if _, okk := hash[n+j]; !okk {
					break
				} else {
					hash[n+j] = hash[n+j] + toAdd
					if maxx < hash[n+j] {
						maxx = hash[n+j]
					}
				}
			}
		}
	}

	return maxx
}
