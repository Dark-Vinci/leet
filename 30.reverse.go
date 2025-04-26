package main

import (
	"math"
)

func reverse(x int) int {
	isPositive := true

	if x < 0 {
		isPositive = false
		x = -x
	}

	var recurse func(a int, b int) int

	recurse = func(a, base int) int {
		if a%10 == a {
			return a
		}

		pow := int(math.Pow(10, float64(base)))

		base--

		return ((a % 10) * pow) + recurse(a/10, base)
	}

	base := int(math.Log10(float64(x)))

	result := recurse(x, base)

	if result > math.MaxInt32 {
		return 0
	}

	if !isPositive {
		return -result
	}

	return result
}
