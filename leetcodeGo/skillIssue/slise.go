package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[2:4]        // b = {3, 4} 5; len 2 cap 3; based on a
	c := append(b, 10) // c = {3, 4, 10}; len 3 cap 3; based on b; b = {3, 4} 10; a = {1,2,3,4,10}
	c[1] = 55          // a = {1, 2, 3, 55, 10}; b = {3, 55}; c = {3 55 10}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
