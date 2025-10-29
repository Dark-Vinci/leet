package arrays

func canReach(s string, minJump int, maxJump int) bool {
	dp, l := make([]bool, len(s)), len(s)
	ii, jj := 0, 0

	dp[0] = true

	for i := 0; i < l; i++ {
		if !dp[i] {
			continue
		}

		ii, jj = max(i+minJump, jj), min(i+maxJump, l-1)

		for k := ii; k <= jj; k++ {
			if s[k] != '0' {
				continue
			}

			dp[k] = true
		}

		if dp[l-1] {
			return true
		}
	}

	return dp[len(dp)-1]
}
