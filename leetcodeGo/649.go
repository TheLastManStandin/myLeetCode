package main

import "fmt"

func predictPartyVictory(senate string) string {
	n := len(senate)
	i := 0
	var DSenators []int = make([]int, 0)
	var RSenators []int = make([]int, 0)

	for ; i < n; i++ {
		if senate[i] == 'D' {
			DSenators = append(DSenators, i)
		} else {
			RSenators = append(RSenators, i)
		}
	}

	for len(DSenators) > 0 && len(RSenators) > 0 {
		i++
		if DSenators[0] > RSenators[0] {
			DSenators = DSenators[1:]
			RSenators = append(RSenators, i)
			RSenators = RSenators[1:]

		} else {
			RSenators = RSenators[1:]
			DSenators = append(DSenators, i)
			DSenators = DSenators[1:]
		}
	}

	if len(DSenators) == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func main() {
	fmt.Println(predictPartyVictory("RRDDD"))
}
