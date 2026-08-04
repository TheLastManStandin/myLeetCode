package main

import "fmt"

func modX(x *int) {
	*x = 5
}

func getCoordinates() (x int, y int) {
	//var x, y int
	x = 10
	y = 20
	defer func(x *int) {
		*x = 5
	}(&x)
	return
}

func main() {
	x, y := getCoordinates()
	fmt.Println(x, y)
}
