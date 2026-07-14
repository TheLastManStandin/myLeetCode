package sliding_window

func characterReplacement(s string, k int) int {
	slen := len(s)
	l := 0
	r := 0
	freeChars := 0
	windowMap := make(map[byte]int)
	ans := 1

	if slen <= 1 {
		return slen
	}
	windowMap[s[0]] = 1

	for r < slen-1 {
		if freeChars <= k {
			r++
			_, ok := windowMap[s[r]]
			if ok {
				windowMap[s[r]]++
			} else {
				windowMap[s[r]] = 1
			}

			maxChar := findMaxChar(windowMap)

			windowLen := r - l + 1
			if windowLen-windowMap[maxChar] > k {
				windowMap[s[l]]--
				if windowMap[s[l]] == 0 {
					delete(windowMap, s[l])
				}
				l++
			} else {
				ans = windowLen
			}
		}
	}
	return ans
}

func findMaxChar(windowMap map[byte]int) byte {
	maxCount := 0
	char := byte(0)
	for k, v := range windowMap {
		if v > maxCount {
			maxCount = v
			char = k
		}
	}
	return char
}
