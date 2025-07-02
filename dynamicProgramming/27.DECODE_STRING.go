package dynamicProgramming

import (
	"strconv"
	"strings"
)

func numDecodings(s string) int {
	var dp func(i int, memo map[int]int) int

	dp = func(i int, memo map[int]int) int {
		if i == len(s) {
			return 1
		}

		if v, ok := memo[i]; ok {
			return v
		}

		result := 0

		if i+1 <= len(s) && isValidd(string(s[i:i+1])) {
			result += dp(i+1, memo)
		}

		if i+2 <= len(s) && isValidd(string(s[i:i+2])) {
			result += dp(i+2, memo)
		}

		memo[i] = result

		return memo[i]
	}

	return dp(0, map[int]int{})
}

func isValidd(ss string) bool {
	if strings.HasPrefix(ss, "0") {
		return false
	}

	v, _ := strconv.Atoi(ss)

	if v > 26 {
		return false
	}

	return true
}
