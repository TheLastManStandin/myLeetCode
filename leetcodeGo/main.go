package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{}
	a = append(a, []int{1, 2, 3, 4, 5}...) // len = 5, cap = 6

	fmt.Println(cap(a)) // 6

	b := append(a, 6) // b - based on a; b = {1,2,3,4,5,6} len = 6, cap = 6; a = {1,2,3,4,5}6
	c := append(b, 7) // c - new massive; c = {1,2,3,4,5,6,7}; len = 7, cap = 12

	c[1] = 0 // a = {1,2,3,4,5}6; b = {1,2,3,4,5,6}; c = {1,0,3,4,5,6,7}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	rand.Int63n()
}
