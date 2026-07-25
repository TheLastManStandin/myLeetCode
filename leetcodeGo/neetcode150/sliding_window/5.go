package sliding_window

func minWindow(s string, t string) string {
	slen := len(s)
	tlen := len(t)
	windowHash := make(map[rune]int)
	windowLeft := 0       // Включительно
	windowRight := len(t) // Включительно

	if slen < tlen {
		return ""
	}

	for windowRight < slen && windowLeft < slen-tlen {
		for !moveRight(s, &windowRight, &windowHash) && hashHasZeros(windowHash) {

		}
	}
}

// returns true if right border now is pointing on symbol from t
func moveRight(s string, windowRight *int, windowHash *map[rune]int) bool {
	*windowRight++
	key := rune(s[*windowRight])
	if _, ok := (*windowHash)[key]; ok {
		(*windowHash)[key]++
		return true
	}
	return false
}

func moveLeft(s string, windowLeft *int, windowHash *map[rune]int) bool {
	*windowLeft++
	key := rune(s[*windowLeft])
	if _, ok := (*windowHash)[key]; ok {
		(*windowHash)[key]--
		return true
	}
	return false
}

func hashHasZeros(windowHash map[rune]int) bool {
	for _, val := range windowHash {
		if val == 0 {
			return true
		}
	}
	return false
}
