package main

import "fmt"

func main() {
	var n int
	var cur_prefix string
	var new_str string

	fmt.Scan(&n)
	fmt.Scan(&cur_prefix)

	for i := 1; i < n; i++ {
		fmt.Scan(&new_str)
		for j := 0; j < len(cur_prefix); j++ {
			if new_str[j] != cur_prefix[j] {
				cur_prefix = cur_prefix[:j]
			}
		}
	}

	fmt.Println(cur_prefix)
	return
}
