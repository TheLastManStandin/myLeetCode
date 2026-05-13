package main

import "fmt"

func main() {
	target := 10
	carPos := 6
	carSpeed := 3
	//timeToRichTarget := (target - carPos) / carSpeed
	//if (target-carPos)%carSpeed != 0 {
	//	timeToRichTarget++
	//}
	timeToRichTarget := (target - carPos + carSpeed - 1) / carSpeed

	fmt.Println(timeToRichTarget)
}
