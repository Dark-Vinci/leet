package dynamicProgramming

import "math"

func findRotateSteps(ring string, key string) int {
	var (
		dp   func(i, j int, memo [][]int) int
		l    = len(ring)
		ll   = len(key)
		memo = make([][]int, ll)
	)

	for i := 0; i < ll; i++ {
		memo[i] = make([]int, l)

		for j := 0; j < l; j++ {
			memo[i][j] = math.MaxInt
		}
	}

	dp = func(i, j int, memo [][]int) int {
		if i == ll {
			return 0
		}

		// MEMOIZATION
		if v := memo[i][j]; v != math.MaxInt {
			return v
		}

		clock, anti, press := math.MaxInt, math.MaxInt, math.MaxInt

		// PRESS
		if key[i] == ring[j] {
			press = dp(i+1, j, memo)
		}

		// BACKWARDS
		for m := j - 1; m >= j-l+1; m-- {
			mm := (m + l) % l

			if ring[mm] == key[i] {
				dist := (j - mm + l) % l
				anti = min(anti, dist+dp(i+1, mm, memo))
				break
			}
		}

		// FORWARD
		for k := j + 1; k < l+j-1; k++ {
			kk := k % l

			if ring[kk] == key[i] {
				dist := (kk - j + l) % l
				clock = min(clock, dist+dp(i+1, kk, memo))
				break
			}
		}

		result := min(press, clock, anti)
		if result != math.MaxInt {
			result += 1
		}

		memo[i][j] = result

		return memo[i][j]
	}

	return dp(0, 0, memo)
}
