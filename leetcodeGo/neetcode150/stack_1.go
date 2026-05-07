package main

func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, byte(v))
		} else {
			if len(stack) == 0 {
				return false
			}
			lastInStack := stack[len(stack)-1]
			if lastInStack == '(' && v == ')' {
				stack = stack[:len(stack)-1]
			} else if lastInStack == '{' && v == '}' {
				stack = stack[:len(stack)-1]
			} else if lastInStack == '[' && v == ']' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
