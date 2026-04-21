package main

import (
	"fmt"
)

func recur(n int, m *map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		val, ok := (*m)[n]
		if ok {
			return val
		} else {
			ans := recur(n-1, m) + recur(n-2, m)
			(*m)[n] = ans
			return ans
		}
	}
}

func main() {
	var n int
	m := make(map[int]int)
	fmt.Scan(&n)

	fmt.Println(recur(n, &m))
}
