package main

import "fmt"

func main() {
	nameArray := [6]string{"D", "a", "n", "i", "i", "l"}
	nameSlice := nameArray[:3]
	nameSlice[len(nameSlice)-1] = "m"
	nameSlice = append(nameSlice, "o")
	fmt.Println(nameSlice) // [D a m]
	fmt.Println(nameArray) // [D a m i i l]
	fmt.Println(cap(nameSlice))
}
