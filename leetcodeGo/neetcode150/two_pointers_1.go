package main

import "unicode"

func validate(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s) - 1
	for leftPointer <= rightPointer {
		if validate(s[leftPointer]) == true {
			if validate(s[rightPointer]) == true {
				if unicode.ToLower(rune(s[leftPointer])) != unicode.ToLower(rune(s[rightPointer])) {
					return false
				}
				leftPointer++
				rightPointer--
			} else {
				rightPointer--
			}
		} else {
			leftPointer++
		}
	}
	return true
}
