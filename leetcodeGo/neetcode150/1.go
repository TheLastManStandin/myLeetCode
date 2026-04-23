package main

import "fmt"

func hasDuplicate(nums []int) bool {
	hash := make(map[int]interface{})

	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = nil
	}
	return false
}

func main() {
	val := hasDuplicate([]int{1, 2, 3, 4})
	fmt.Println(val)
}
