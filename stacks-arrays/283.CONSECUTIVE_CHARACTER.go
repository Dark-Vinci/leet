package arrays

func maxPower(s string) int {
	result, count := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			result = max(result, count)
			count = 1
			continue
		}

		count++
	}

	return max(result, count)
}
