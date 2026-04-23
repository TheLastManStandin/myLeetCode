package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lenn := len(s)
	hash := make(map[byte]int)

	for i := 0; i < lenn; i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
	}

	for i := 0; i < lenn; i++ {
		if _, ok := hash[t[i]]; ok {
			if hash[t[i]] >= 1 {
				if hash[t[i]] == 1 {
					delete(hash, t[i])
				} else {
					hash[t[i]]--
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(hash) == 0 {
		return true
	}
	return false
}

func main() {
	val := isAnagram("jepau", "pajei")

	fmt.Println(val)
}
