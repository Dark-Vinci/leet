package main

import "math"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	result := 0

	for i := 2; i < n; i++ {
		m, rem := n/i, n%i

		if rem == 0 {
			result = max(result, int(math.Pow(float64(m), float64(i))))
			continue
		}

		arr := make([]int, i)

		for j := 0; j < len(arr); j++ {
			arr[j] = m
		}

		for j := 0; j < rem; j++ {
			arr[j] = arr[j] + 1
		}

		mx := 1

		for _, val := range arr {
			mx *= val
		}

		result = max(result, mx)
	}

	return result
}
