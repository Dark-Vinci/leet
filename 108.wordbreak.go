package main

import (
	"slices"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	var (
		n  = len(s)
		dp = make([]bool, len(s)+1)
	)

	dp[n] = true

	for i := n; i >= 0; i-- {
		for _, v := range wordDict {
			if i+len(v) <= n && s[i:i+len(v)] == v {
				dp[i] = dp[i+len(v)]
			}

			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

// This won't work for one case for some reason unknown
func wordBreakWrong(s string, wordDict []string) bool {
	indexes := make([]int, 0)

	for i := 0; i < len(wordDict); i++ {
		index := strings.Index(s, wordDict[i])

		if index < 0 {
			continue
		}

		for j := index; j < index+len(wordDict[i]); j++ {
			if slices.Contains(indexes, j) {
				return false
			} else {
				indexes = append(indexes, j)
			}
		}
	}

	if len(indexes) == 0 {
		return false
	}

	return true
}
