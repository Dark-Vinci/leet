package arrays

import "strings"

func numSub(s string) int {
	sp := strings.Split(s, "0")
	result := 0

	for _, str := range sp {
		result += (len(str) * (len(str) + 1)) / 2

		result %= 1_000_000_007
	}

	return result
}
