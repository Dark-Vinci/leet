package stacks_arrays

func minSwaps(s string) int {
	closeCount, maxCloseCount := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			closeCount--
		} else {
			closeCount++
		}

		maxCloseCount = max(maxCloseCount, closeCount)
	}

	return (maxCloseCount + 1) / 2
}
