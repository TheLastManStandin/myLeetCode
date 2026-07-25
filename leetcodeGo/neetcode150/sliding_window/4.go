package sliding_window

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	hashmap := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		hashmap[s1[i]]++
	}

	l := 0
	r := len(s1) - 1

	for i := 0; i <= r; i++ {
		if _, ok := hashmap[s2[i]]; ok {
			hashmap[s2[i]]--
		}
	}
	allZero := true
	for _, v := range hashmap {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return true
	}
	r++
	l++

	for r < len(s2) {
		lchar := s2[l-1]
		rchar := s2[r]
		if _, ok := hashmap[rchar]; ok {
			hashmap[rchar]--
		}
		if _, ok := hashmap[lchar]; ok {
			hashmap[lchar]++
		}

		allZero := true
		for _, v := range hashmap {
			if v != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return true
		}
		r++
		l++
	}
	return false
}
