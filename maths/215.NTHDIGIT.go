package maths

import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	m, l, r := 1, 1, 9

	for r < n {
		m++
		width := 9 * int(math.Pow(10.0, float64(m-1)))
		count := m * width
		l = r
		r += count
	}

	leftNum := int(math.Pow(10.0, float64(m-1))) - 1
	change := n - l
	num := (change - 1) / m
	digit := (change - 1) % m

	targetNum := leftNum + num + 1

	targetStr := fmt.Sprintf("%v", targetNum)

	return int(targetStr[digit] - '0')
}
