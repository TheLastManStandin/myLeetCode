package main

func productExceptSelf(nums []int) []int {
	numLen := len(nums)
	res := make([]int, 0, numLen)
	product := 1

	// hashmap key - index : [prefix, suffix]
	hash := make(map[int][2]int)

	for i := 0; i < numLen; i++ {
		product *= nums[i]
		hash[i] = [2]int{product, 0}
	}

	product = 1

	for i := 0; i < numLen; i++ {
		product *= nums[numLen-i-1]
		val := hash[numLen-i-1]
		val[1] = product
		hash[numLen-i-1] = val
	}

	for i := 0; i < numLen; i++ {
		pes := 1
		if i > 0 {
			pes = hash[i-1][0]
		}
		if i < numLen-1 {
			pes *= hash[i+1][1]
		}
		res = append(res, pes)
	}

	return res
}
