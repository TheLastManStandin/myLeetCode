package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := Constructor()

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1 := stack.Top()
			stack.Pop()
			num2 := stack.Top()
			stack.Pop()

			if token == "+" {
				stack.Push(num2 + num1)
			} else if token == "-" {
				stack.Push(num2 - num1)
			} else if token == "*" {
				stack.Push(num2 * num1)
			} else if token == "/" {
				stack.Push(num2 / num1)
			}
		} else {
			num, err := strconv.ParseInt(token, 10, 32)
			if err != nil {
				panic(err)
			}

			stack.Push(int(num))
		}
	}

	return stack.Top()
}
