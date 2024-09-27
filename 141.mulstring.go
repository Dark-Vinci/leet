package main

import (
	"fmt"
	"math"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var (
		l1     = len(num1) - 1
		l2     = len(num2) - 1
		result = make([]byte, l1+l2+2)
	)

	for i := l1; i >= 0; i-- {
		for j := l2; j >= 0; j-- {
			result[i+j+1] += (num1[i] - '0') * (num2[j] - '0')

			if result[i+j+1] >= 10 {
				result[i+j] += result[i+j+1] / 10
				result[i+j+1] %= 10
			}
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	for i := range result {
		result[i] += '0'
	}

	return string(result)
}

func multiply_INT_OVERFLOW_AT_LARGE_NUMBER(num1 string, num2 string) string {
	var (
		l1     = len(num1) - 1
		l2     = len(num2) - 1
		index  = 0
		result = 0
	)

	for i := l1; i >= 0; i-- {
		carry, a, ind, current := 0, int(num1[i]-'0'), 0, 0

		for j := l2; j >= 0; j-- {
			mul := a * int(num2[j]-'0')

			if carry > 0 {
				mul += carry
				carry = 0
			}

			strMul := fmt.Sprintf("%v", mul)

			if len(strMul) > 1 {
				carry = int(strMul[0] - '0')
				strMul = string(strMul[1])
			}

			p := int(math.Pow(10, float64(ind)))

			current += p * int(strMul[0]-'0')

			ind++
		}

		if carry > 0 {
			current += int(math.Pow(10, float64(ind))) * carry
		}

		result += int(math.Pow(10, float64(index))) * current
		index++
	}

	return fmt.Sprintf("%v", result)
}
