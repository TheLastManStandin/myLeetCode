package sliding_window

//func lengthOfLongestSubstring(s string) int {
//	hash := make(map[int]int)
//	lastCharIndex := 0
//	res := 1
//
//	if len(s) == 0 {
//		return 0
//	}
//
//	for i := 0; i < len(s); i++ {
//		if _, ok := hash[int(s[i])]; ok {
//			if i-lastCharIndex > res {
//				res = i - lastCharIndex
//			}
//
//			for s[lastCharIndex] != s[i] && lastCharIndex != i {
//				lastCharIndex += 1
//				delete(hash, int(s[lastCharIndex]))
//			}
//			for s[lastCharIndex] == s[i] && lastCharIndex != i {
//				lastCharIndex += 1
//			}
//		}
//		hash[int(s[i])] = true
//	}
//
//	if len(s)-lastCharIndex > res {
//		res = len(s) - lastCharIndex
//	}
//	return res
//}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res := 0, 0

	for r := 0; r < len(s); r++ {
		if idx, found := mp[s[r]]; found {
			l = max(idx+1, l)
		}
		mp[s[r]] = r
		if (r-l)+1 > res {
			res = r - l + 1
		}
	}
	return res
}
