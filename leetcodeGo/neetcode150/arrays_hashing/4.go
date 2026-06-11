package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedStr := string(runes)

		val, ok := hash[sortedStr]
		if !ok {
			hash[sortedStr] = []string{
				str,
			}
		} else {
			hash[sortedStr] = append(val, str)
		}
	}

	ans := make([][]string, 0)
	for _, v := range hash {
		ans = append(ans, v)
	}

	return ans
}
