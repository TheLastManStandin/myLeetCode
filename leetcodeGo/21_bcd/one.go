package main

import (
	"fmt"
)

func main() {
	var n, ans int
	fmt.Scan(&n)

	for i := 3; i < n; i = i + 2 {
		for j := 0; j < i; j++ {
			ans += j
		}
	}
	fmt.Println(ans)
}
