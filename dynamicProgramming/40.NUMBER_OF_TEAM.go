package dynamicProgramming

func numTeams(rating []int) int {
	var dp func(curr, count, prevIndex int, typ bool, memo map[[4]int]int) int

	dp = func(curr, count, prevIndex int, typ bool, memo map[[4]int]int) int {
		if curr == len(rating) {
			return 0
		}

		key := [4]int{prevIndex, curr, count, 0}
		if typ {
			key[3] = 1
		}

		if v, ok := memo[key]; ok {
			return v
		}

		if count == 3 {
			return 1
		}

		result := 0
		count += 1

		for j := curr; j < len(rating); j++ {
			if prevIndex == -1 {
				result += dp(j+1, count, j, true, memo)
				result += dp(j+1, count, j, false, memo)
			} else if !typ && rating[j] < rating[prevIndex] {
				result += dp(j, count, j, false, memo)
			} else if typ && rating[j] > rating[prevIndex] {
				result += dp(j, count, j, true, memo)
			}
		}

		memo[key] = result

		return result
	}

	return dp(0, 0, -1, false, map[[4]int]int{})
}

// TLE
func numTeams_(rating []int) int {
	var (
		dp func(count, curr, prev int, isInc bool) int
		l  = len(rating)
	)

	dp = func(count, curr, prev int, isInc bool) int {
		if count == 3 {
			return 1
		}

		if curr == l {
			return 0
		}

		result := 0

		for next := curr + 1; next < l; next++ {
			if prev == -1 {
				result += dp(1, next, curr, true)
				result += dp(1, next, curr, false)
			} else if isInc && rating[next] > rating[prev] {
				result += dp(count+1, next, next, true)
			} else if !isInc && rating[next] < rating[prev] {
				result += dp(count+1, next, next, false)
			}
		}

		return result
	}

	return dp(0, 0, -1, false)
}
