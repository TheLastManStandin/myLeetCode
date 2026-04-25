package main

import "strings"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	outStringBuilder := strings.Builder{}

	for _, str := range strs {
		for _, ch := range str {
			if ch == '^' {
				outStringBuilder.WriteRune('^')
			}
			outStringBuilder.WriteRune(ch)
		}
		outStringBuilder.WriteString("^f")
	}

	return outStringBuilder.String()
}

func (s *Solution) Decode(encoded string) []string {
	oneStr := strings.Builder{}
	strs := make([]string, 0)

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == '^' {
			i++
			if encoded[i] == 'f' {
				strs = append(strs, oneStr.String())
				oneStr.Reset()
			} else {
				oneStr.WriteByte(encoded[i])
			}
		} else {
			oneStr.WriteByte(encoded[i])
		}
	}

	return strs
}
