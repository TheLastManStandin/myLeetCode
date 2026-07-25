package sliding_window

func minWindow(s string, t string) string {
	ans := ""
	ansLen := 0

	slen := len(s)
	tlen := len(t)

	if slen < tlen {
		return ""
	}

	hashmap := make(map[byte]int)

	for i := 0; i < tlen; i++ {
		if _, ok := hashmap[t[i]]; !ok {
			hashmap[t[i]] = 0
		}
	}

	l := 0
	r := 0
	hashmapHasNoZeros := false

	for r < slen && !hashmapHasNoZeros {
		if _, ok := hashmap[s[r]]; ok {
			hashmap[s[r]]++

			hasZeros := false
			for _, val := range hashmap {
				if val == 0 {
					hasZeros = true
					break
				}
			}
			if !hasZeros {
				hashmapHasNoZeros = true
			}
		}
		r++
	}
	ansLen = r
	ans = s[:ansLen]
	var buf byte = byte(0)

	for l < r {
		if _, ok := hashmap[s[l]]; ok {
			break
		} else {
			l++
			if ansLen > r-l {
				ansLen = r - l
				ans = s[l:r]
			}
		}
	}

	for r < slen && l < slen-tlen {
		if buf == 0 {
			if ansLen > r-l {
				ansLen = r - l
				ans = s[l : r+1]
			}
			buf = s[l]
			l++
			hashmap[buf]--
			if hashmap[buf] != 0 {
				buf = 0
			}
			for l < r && l < slen-tlen {
				if _, ok := hashmap[s[l]]; ok {
					if s[l] == buf {
						hashmap[buf]--
					} else {
						break
					}
				}
				l++
			}
		} else {
			if s[r] == buf {
				if ansLen > r-l {
					ansLen = r - l
					ans = s[l : r+1]
				}
				buf = 0
			} else if _, ok := hashmap[s[r]]; ok {
				hashmap[s[r]]++
				r++
			} else {
				r++
			}
		}
	}

	return ans
}
