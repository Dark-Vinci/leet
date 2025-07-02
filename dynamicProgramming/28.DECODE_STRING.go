package dynamicProgramming

import (
	"fmt"
	"strconv"
)

func numDecodings_(s string) int {
	var (
		dp  func(i int, memo map[int]int) int
		mod = 1000000007
	)

	dp = func(i int, memo map[int]int) int {
		if i == len(s) {
			return 1
		}

		if v, ok := memo[i]; ok {
			return v
		}

		result := 0

		if s[i] == '*' {
			result = (result + 9*dp(i+1, memo)) % mod
		} else if s[i] != '0' {
			result = (result + dp(i+1, memo)) % mod
		}

		if i+1 < len(s) {
			a, b := string(s[i]), string(s[i+1])

			if a == "*" && b == "*" {
				result = (result + 15*dp(i+2, memo)) % mod
			} else if a == "*" {
				bb, _ := strconv.Atoi(b)

				if bb >= 0 && bb <= 6 {
					result = (result + 2*dp(i+2, memo)) % mod
				} else {
					result = (result + dp(i+2, memo)) % mod
				}
			} else if b == "*" {
				aa, _ := strconv.Atoi(a)

				if aa == 1 {
					result = (result + 9*dp(i+2, memo)) % mod
				} else if aa == 2 {
					result = (result + 6*dp(i+2, memo)) % mod
				}
			} else {
				num, _ := strconv.Atoi(s[i : i+2])

				if num >= 10 && num <= 26 {
					result = (result + dp(i+2, memo)) % mod
				}
			}
		}

		memo[i] = result

		return memo[i]
	}

	return dp(0, map[int]int{})
}

func numDecodings__(s string) int {
	var dp func(i int, memo map[int]int) int

	dp = func(i int, memo map[int]int) int {
		if i == len(s) {
			return 1
		}

		if v, ok := memo[i]; ok {
			return v
		}

		result := 0

		if i+1 <= len(s) {
			a := string(s[i : i+1][0])

			if a == "*" {
				for jj := 1; jj <= 9; jj++ {
					result += dp(i+1, memo)
				}
			} else if a != "0" {
				result += dp(i+1, memo)
			}
		}

		if i+2 <= len(s) {
			a, b := string(s[i : i+2][0]), string(s[i : i+2][1])

			if a == "*" && b == "*" {
				for jj := 1; jj <= 9; jj++ {
					for kk := 1; kk <= 9; kk++ {
						sss := fmt.Sprintf("%v%v", jj, kk)

						if isValidd(sss) {
							result += dp(i+2, memo)
						}
					}
				}
			} else if a == "*" {
				for jj := 1; jj <= 9; jj++ {
					sss := fmt.Sprintf("%v", jj) + b

					if isValidd(sss) {
						result += dp(i+2, memo)
					}
				}
			} else if b == "*" {
				for jj := 1; jj <= 9; jj++ {
					sss := a + fmt.Sprintf("%v", jj)

					if isValidd(sss) {
						result += dp(i+2, memo)
					}
				}
			} else {
				result += dp(i+2, memo)
			}
		}

		memo[i] = result

		return memo[i]
	}

	return dp(0, map[int]int{})
}
