package stacks_arrays

func countSubstrings(s string, c byte) int64 {
	result, count := int64(0), 0

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++

			result += int64(count)
		}
	}

	return result
}
