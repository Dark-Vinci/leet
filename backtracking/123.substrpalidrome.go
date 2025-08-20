package backtracking

import "slices"

func partition(s string) [][]string {
	var (
		result = make([][]string, 0)
		dfs    func(int, []string)
	)

	dfs = func(i int, bucket []string) {
		if i == len(s) {
			result = append(result, slices.Clone(bucket))
			return
		}

		for j := i; j < len(s); j++ {
			if isValidPal(s[i : j+1]) {
				bucket = append(bucket, s[i:j+1])
				dfs(j+1, bucket)
				bucket = bucket[:len(bucket)-1]
			}
		}
	}

	dfs(0, []string{})

	return result
}

func isValidPal(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
