package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	result, base := 10, 9

	for i := 2; i <= n; i++ {
		base *= 9 - i + 2
		result += base
	}

	return result
}
