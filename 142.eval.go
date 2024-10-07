package main

import (
	"slices"
	"strconv"
)

func evalRPN(tokens []string) int {
	result := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		if s, e := strconv.Atoi(tokens[i]); e == nil {
			result = append(result, s)
			continue
		}

		a, b := result[len(result)-1], result[len(result)-2]
		result = slices.Delete(result, len(result)-2, len(result))

		switch tokens[i] {
		case "*":
			result = append(result, a*b)
		case "/":
			result = append(result, b/a)
		case "+":
			result = append(result, a+b)
		case "-":
			result = append(result, b-a)
		}
	}

	return result[0]
}
