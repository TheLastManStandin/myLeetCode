package main

import "fmt"

const MAX = 5

func main() {
	s := generate()
	mutation(s)
	fmt.Println(s[0:MAX])
}

func generate() []int {
	out := make([]int, 0, MAX)
	for i := 1; i < MAX; i++ {
		out = append(out, i) // {1,2,3,4}0 len = 4 cap = 5
	}
	return out
}

func mutation(s []int) {
	s = append(s, -1)
}
